package models

type Student struct {
	StudentID 		uint		`gorm:"primaryKey;autoIncrement:true"`
	Name			string		`gorm:"size:50;not null"`
	Grade			uint		`gorm:"size:2;not null"`
}