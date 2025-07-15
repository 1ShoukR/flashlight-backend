package models
type Student struct {
    StudentID uint   `gorm:"primaryKey;autoIncrement"` 
    Name      string `gorm:"size:50;not null"`
    Grade     uint   `gorm:"not null"`  
}