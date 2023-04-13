package tests

import (
	"encoding/json"
	"testing"

	"github.com/DiarCode/todo-go-api/src/controllers"
	"github.com/DiarCode/todo-go-api/src/dto"
	"github.com/DiarCode/todo-go-api/src/models"
	"github.com/DiarCode/todo-go-api/src/utils"
)

func TestSignup(t *testing.T) {
	existingUser := &models.User{
		Name:     "TestName",
		Email:    "exists@mail.ru",
		Password: utils.HashPassword([]byte("Password123")),
	}
	createNewUser(existingUser)
	tests := []struct {
		name    string
		body    interface{}
		wantErr bool
	}{
		{
			name: "Signup invalid body",
			body: struct {
				Gender string `json:"gender"`
			}{
				Gender: "Male"},
			wantErr: true,
		},
		{
			name: "Signup invalid email",
			body: &dto.SignupDto{
				Name:     "TestName",
				Email:    "invalid.ru",
				Password: "Password123",
			},
			wantErr: true,
		},
		{
			name: "Signup to already exist user",
			body: &dto.SignupDto{
				Name:     existingUser.Name,
				Email:    existingUser.Email,
				Password: existingUser.Password,
			},
			wantErr: true,
		},
		{
			name: "Signup to not-existing user",
			body: &dto.SignupDto{
				Name:     "NewName",
				Email:    "new.email@mail.ru",
				Password: "Password123",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBytes, _ := json.Marshal(tt.body)

			MockApp.ctx.Request().Reset()
			MockApp.ctx.Request().Header.Set("Content-Type", "application/json")
			MockApp.ctx.Request().SetBodyRaw(jsonBytes)

			err := controllers.Signup(MockApp.ctx)

			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	deleteUser(existingUser)
}
