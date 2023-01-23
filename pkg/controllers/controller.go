package controllers

import (
	"github.com/DiarCode/todo-go-api/pkg/config/database"
	"github.com/DiarCode/todo-go-api/pkg/models"
)

var (
	db = database.DB
)

type Todo models.Todo
type User models.User
