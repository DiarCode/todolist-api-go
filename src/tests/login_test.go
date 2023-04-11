package tests

import (
	"github.com/DiarCode/todo-go-api/src/database"
	"github.com/DiarCode/todo-go-api/src/models"
)

// func TestLogin(t *testing.T) {
// 	existingUser := &models.User{
// 		Name:     "TestName",
// 		Email:    "exists@mail.ru",
// 		Password: utils.HashPassword([]byte("Password123")),
// 	}
// 	createNewUser(existingUser)

// 	tests := []struct {
// 		name    string
// 		body    interface{}
// 		wantErr bool
// 	}{
// 		{
// 			name:    "Login not-existing user",
// 			body:    &dto.LoginDto{Email: "doesnot.exist@mail.ru", Password: "Password123"},
// 			wantErr: true,
// 		},

// 		{
// 			name: "Login invalid password",
// 			body: &dto.LoginDto{
// 				Email:    existingUser.Email,
// 				Password: "InvalidPassword123",
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "Login to existing user",
// 			body: &dto.LoginDto{
// 				Email:    existingUser.Email,
// 				Password: "Password123",
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Login invalid body",
// 			body: struct {
// 				Name string `json:"name"`
// 			}{
// 				Name: "Andrew"},
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			jsonBytes, _ := json.Marshal(tt.body)

// 			MockApp.ctx.Request().Reset()
// 			MockApp.ctx.Request().Header.Set("Content-Type", "application/json")
// 			MockApp.ctx.Request().SetBodyRaw(jsonBytes)

// 			err := controllers.Login(MockApp.ctx)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}

// 	deleteUser(existingUser)
// }

func createNewUser(user *models.User) {
	database.DB.Create(&user)
}

func deleteUser(user *models.User) {
	database.DB.Delete(&user)
}
