package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n=============================\n", sql)
}

var db *gorm.DB

func main() {

	var err error

	dsn := "root:password@tcp(127.0.0.1:3306)/database?parseTime=true"
	dial := mysql.Open(dsn)
	db, err = gorm.Open(dial, &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: true,
	})

	if err != nil {
		panic(err)
	}

	//db.AutoMigrate(User{}, Test{}, Gender{})

	//CreateGender("Male")
	//GetGenders()
	//GetGender(1)
	//GetGenderbyName
}

type Customer struct{
	ID uint
	Name string
	Gender Gender
	GenderID uint
}


func CreateGender(name string) {

	gender := Gender{Name: name}
	tx := db.Create(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

func GetGenders() {
	genders := []Gender{}
	tx := db.Order("id").Find(&genders)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}

	fmt.Println(genders)

}

func GetGender(id uint) {
	gender := Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

func GetGenderbyName(name string) {
	gender := Gender{}
	tx := db.Order("id").First(&gender, "name=?", name)
	// tx := db.Where("name=?",name).First(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

func UpdateGender(id uint, name string) {
	gender := Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	gender.Name = name
	tx = db.Save(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
}

func UpdateGender2(id uint, name string) {
	gender := Gender{Name: name}
	tx := db.Model(&Gender{}).Where("id=?", id).Updates(gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
}

// case that not use gorm.model
func DeleteGender(id uint) {
	tx := db.Delete(&Gender{}, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
}

func CreateTest(code uint, name string) {
	test := Test{Code: code, Name: name}
	db.Create(&test)
}

func GetTests() {
	tests := []Test{}
	db.Find(&tests)
	for _, t := range tests {
		fmt.Printf("%v|%v\n", t.ID, t.Name)
	}
}

// case that use gorm.model
// has to use Unscoped if not use Unscoped
// It will soft delete
// by update delete_at column
func DeleteTest(id uint) {
	db.Unscoped().Delete(&Test{}, id)
}

type User struct {
	ID       uint
	Name     string
	Gmail    string
	CreateAt time.Time
}

// gorm.model https://gorm.io/docs/models.html
//
//	type Model struct {
//		ID        uint `gorm:"primarykey"`
//		CreatedAt time.Time
//		UpdatedAt time.Time
//		DeletedAt DeletedAt `gorm:"index"`
//	}
type Test struct {
	gorm.Model
	Code uint   `gorm:"primaryKey;comment:This is Code"`
	Name string `gorm:"column:myname;size:20;uniquedefault:hello;not null"`
}

// change table name
func (t Test) TableName() string {
	return "MyTest"
}

type Gender struct {
	ID   uint
	Name string `gorm:"unique;"`
}
