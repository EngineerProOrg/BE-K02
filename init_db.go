package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type Professor struct {
	ProfID    uint `gorm:"primaryKey"`
	ProfLName string
	ProfFName string
}

type Student struct {
	StuID       uint `gorm:"primaryKey"`
	StudFName   string
	StudLName   string
	StudStreet  string
	StudCity    string
	StudZIP     string
	Classes     []*Class `gorm:"many2many:enroll"`	
}

type Class struct {
	ClassID     uint `gorm:"primaryKey"`
	ClassName   string
	Room        Room
	CourseID    uint
	Course      Course `gorm:"foreignKey:CourseID"`
	ProfID      uint
	Professor   Professor `gorm:"foreignKey:ProfID"`
	Students    []*Student `gorm:"many2many:enroll"`
}

type Course struct {
	CourseID    uint `gorm:"primaryKey"`
	CourseName  string
}

type Room struct {
	RoomID    uint `gorm:"primaryKey"`
	RoomLoc   string
	RoomCap   string
	ClassID   uint
}

func main() {
	dsn := "root:root@tcp(localhost:3307)/engineer_pro?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error to connect to database")
	}

	db.AutoMigrate(&Professor{})
	db.AutoMigrate(&Student{})
	db.AutoMigrate(&Class{})
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&Room{})

}
