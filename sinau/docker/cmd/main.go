package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/IBM/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "ivandhitya/docker/cmd/docs"

	"ivandhitya/docker/model"

	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db            *gorm.DB
	redisClient   *redis.Client
	kafkaProducer sarama.SyncProducer
	ctx           = context.Background()
)

// Initialize PostgreSQL
func initDatabase() {
	dsn := "host=postgres user=sinau password=12345 dbname=app_db port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&model.Student{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database connected and migrated")
}

// Initialize Redis
func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Redis connected")
}

// Initialize Kafka
func initKafka() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	var err error
	kafkaProducer, err = sarama.NewSyncProducer([]string{"kafka:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	log.Println("Kafka producer initialized")
}

// Kafka Consumer
func startKafkaConsumer() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumerGroup([]string{"kafka:9092"}, "cache-invalidator", config)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer group: %v", err)
	}
	defer consumer.Close()

	handler := cacheInvalidator{}
	log.Println("Kafka consumer started...")
	for {
		if err := consumer.Consume(ctx, []string{"invalidate-cache"}, &handler); err != nil {
			log.Fatalf("Error consuming Kafka messages: %v", err)
		}
	}
}

type cacheInvalidator struct{}

func (h *cacheInvalidator) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (h *cacheInvalidator) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (h *cacheInvalidator) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		cacheKey := string(message.Value)
		log.Printf("Received cache invalidation message: %s", cacheKey)

		if err := redisClient.Del(ctx, cacheKey).Err(); err != nil {
			log.Printf("Failed to delete cache key %s: %v", cacheKey, err)
		} else {
			log.Printf("Cache key %s invalidated successfully", cacheKey)
		}
		sess.MarkMessage(message, "")
	}
	return nil
}

// @Summary Get student by ID
// @Description Retrieve student data by ID, either from cache or database
// @Tags students
// @Accept json
// @Produce json
// @Param id path string true "Student ID"
// @Success 200 {object} Student
// @Failure 400 {string} string "Missing 'id' parameter"
// @Failure 404 {string} string "Student not found"
// @Failure 500 {string} string "Error accessing cache"
// @Router /student/{id} [get]
func getStudentHandler(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "Missing 'id' parameter")
	}

	cacheKey := "student:" + id
	cached, err := redisClient.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		var student model.Student
		if err := db.First(&student, id).Error; err != nil {
			return c.String(http.StatusNotFound, "Student not found")
		}
		data, _ := json.Marshal(student)
		redisClient.Set(ctx, cacheKey, data, 10*time.Minute)
		return c.JSON(http.StatusOK, student)
	} else if err != nil {
		return c.String(http.StatusInternalServerError, "Error accessing cache")
	}
	return c.String(http.StatusOK, cached)
}

// updateStudentHandler handles updating student data.
func updateStudentHandler(c echo.Context) error {
	var student model.Student
	if err := c.Bind(&student); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	if err := db.Save(&student).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Failed to update student")
	}

	cacheKey := "student:" + strconv.Itoa(student.ID)
	message := &sarama.ProducerMessage{
		Topic: "invalidate-cache",
		Value: sarama.StringEncoder(cacheKey),
	}
	if _, _, err := kafkaProducer.SendMessage(message); err != nil {
		log.Printf("Failed to publish Kafka message: %v", err)
	}

	return c.String(http.StatusOK, "Student updated successfully")
}

// Main
func main() {
	initDatabase()
	initRedis()
	initKafka()
	go startKafkaConsumer()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/student/:id", getStudentHandler)
	e.PUT("/student/update", updateStudentHandler)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
