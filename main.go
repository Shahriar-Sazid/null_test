package main

import (
	"fmt"
	"log"

	"github.com/LukaGiorgadze/gonull"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // fixed import line
)

// User model
type User struct {
	ID    uint                 `gorm:"primary_key"`
	Name  string               `gorm:"type:varchar(100)"`
	Score gonull.Nullable[int] `gorm:"type:numeric"`
}

var db *gorm.DB
var err error

const (
	host     = "localhost"
	port     = 5435
	user     = "gsfood"
	password = "admin"
	dbname   = "gsfood"
)

// func main() {
// 	// Initialize the connection to the PostgreSQL database
// 	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// 	db, err = gorm.Open("postgres", dsn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Migrate the schema (create the table)
// 	db.AutoMigrate(&User{})

// 	// Create a user
// 	createUser("John Doe")

// 	// Read a user by ID
// 	readUser(1)
// }

func main() {
	TestWithDefault()
}

// Create function to insert a new user
func createUser(name string) {
	user := User{Name: name, Score: gonull.NewNullable(3)}
	if err := db.Create(&user).Error; err != nil {
		log.Fatal("Error creating user:", err)
	}
	fmt.Println("User created:", user)
}

// Read function to retrieve a user by ID
func readUser(id uint) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		log.Fatal("Error reading user:", err)
	}
	fmt.Println("User found:", user)
}
