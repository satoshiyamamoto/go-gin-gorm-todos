package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/satoshiyamamoto/go-gin-gorm-todos/config"
)

// fetch all todos at once
func GetAllTodos(todo *[]Todo) (err error) {
	if err = config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

// insert a todo
func CreateATodo(todo *Todo) (err error) {
	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

// fetch one todo
func GetATodo(todo *Todo, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateATodo(todo *Todo, id string) (err error) {
	fmt.Println(todo)
	config.DB.Save(todo)
	return nil
}

// delete a todo
func DeleteATodo(todo *Todo, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(todo)
	return nil
}
