package entity

type Student struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Grade   int    `json:"grade"`
	Address string `json:"address"`
}
