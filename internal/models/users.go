package models

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	gorm.Model
	Username string `gorm:"primaryKey;unique"`
	Email    string `gorm:"unique"`
	Password string
}

type UserModel struct {
	DB *gorm.DB
}

// CreateUser Function to create a new user
func (m *UserModel) CreateUser(username string, email string, password string) {
	newUser := User{Username: username, Email: email, Password: password}
	if err := m.DB.Create(&newUser).Error; err != nil {
		msg := err.Error()
		// Checks database for existence of a matching username or email
		if err := m.DB.First(&User{Username: username}).Error; err == nil {
			msg = "username is already taken"
		} else if err := m.DB.First(&User{Email: email}).Error; err == nil {
			msg = "user with this email address already exists"
		}
		log.Printf("Unexpected error occurred: %s.\n", msg)
		return
	}
	log.Printf("✅ User '%s' created successfully.\n", username)
}

// DeleteUser Function to delete a user by username
func (m *UserModel) DeleteUser(username string) {
	user, err := m.getUser(username)
	if err != nil {
		log.Fatalf("Unepexted error occured: %s.\n", err.Error())
	}
	if err = m.DB.Delete(&user).Error; err != nil {
		log.Printf("Unexpected error occurred: %s.\n", err.Error())
	}
	log.Printf("✅ User '%s' deleted successfully.\n", username)
}

// getUser Helper function to find user by username and return user instance
func (m *UserModel) getUser(username string) (*User, error) {
	target := User{Username: username}
	if err := m.DB.First(&target).Error; err != nil {
		return nil, errors.New("user '" + username + "' not found")
	}
	return &target, nil
}

func (m *UserModel) ShowTable(name string) {
	rows, err := m.DB.Table(name).Rows()
	if err != nil {
		log.Printf("Error encountered while processing rows in table '%s': %s.\n", name, err.Error())
		return
	}
	columns, err := rows.Columns()
	if err != nil {
		log.Printf("Error encountered while processing rows in table '%s': %s.\n", name, err.Error())
		return
	}
	values := make([]sql.RawBytes, len(columns))
	args := make([]interface{}, len(values))
	for i := range values {
		args[i] = &values[i]
	}
	output := make([]string, 10)
	for rows.Next() {
		err = rows.Scan(args...)
		if err != nil {
			log.Printf("Error encountered while scanning table '%s': %s.\n", name, err.Error())
			return
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			output = append(output, columns[i]+": "+value+"\n")
		}
		output = append(output, "\n")
	}
	if err = rows.Err(); err != nil {

	}
	log.Println(output)
}

func (u *User) String() string {
	return fmt.Sprintf("ID: %d, Username: %s, Email: %s, Password: %s, CreatedAt: %s, UpdatedAt: %s, DeletedAt: %s",
		u.ID, u.Username, u.Email, u.Password, u.CreatedAt.Format(time.DateTime), u.UpdatedAt.Format(time.DateTime), u.DeletedAt.Time.Format(time.DateTime))
}
