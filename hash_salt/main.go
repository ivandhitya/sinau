package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

type User struct {
	Username       string
	HashedPassword string // Hash dari password + salt
}

// Fungsi untuk membuat salt
func generateSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// Fungsi untuk hash password dengan salt
func hashPassword(password string, salt []byte) string {
	hash := sha256.New()
	hash.Write(salt)
	hash.Write([]byte(password))
	hashedPassword := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(append(salt, hashedPassword...))
}

// Fungsi untuk mendaftarkan user baru
func registerUser(username, password string) User {
	salt, _ := generateSalt(16)                    // Buat salt 16 byte
	hashedPassword := hashPassword(password, salt) // Hash password dengan salt
	return User{
		Username:       username,
		HashedPassword: hashedPassword,
	}
}

// Fungsi untuk verifikasi login
func loginUser(user User, inputPassword string) bool {
	// Decode hash yang disimpan untuk mendapatkan salt
	decodedHash, _ := base64.StdEncoding.DecodeString(user.HashedPassword)
	salt := decodedHash[:16] // Ambil salt (ukuran salt 16 byte)

	// Hash password yang dimasukkan dengan salt yang sama
	inputHashedPassword := hashPassword(inputPassword, salt)

	// Bandingkan hash hasil input user dengan hash yang disimpan
	return inputHashedPassword == user.HashedPassword
}

func main() {
	// Mendaftarkan user baru
	newUser := registerUser("john_doe", "securePassword123")
	fmt.Println("User registered with hashed password:", newUser.HashedPassword)

	// Simulasi login dengan password yang benar
	isAuthenticated := loginUser(newUser, "securePassword123")
	fmt.Println("Login successful?", isAuthenticated)

	// Simulasi login dengan password yang salah
	isAuthenticated = loginUser(newUser, "wrongPassword")
	fmt.Println("Login successful?", isAuthenticated)

}
