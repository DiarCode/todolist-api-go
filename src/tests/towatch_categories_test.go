package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/DiarCode/todo-go-api/src/controllers"
	"github.com/DiarCode/todo-go-api/src/database"
	"github.com/DiarCode/todo-go-api/src/dto"
	"github.com/DiarCode/todo-go-api/src/models"
	"github.com/DiarCode/todo-go-api/src/utils"
)

const (
	towatch_category_url = "/api/v1/towatch-category"
)

func TestGetAllTowatchCategories(t *testing.T) {
	user := &models.User{
		Name:     "TestName",
		Email:    "exists@mail.ru",
		Password: utils.HashPassword([]byte("Password123")),
	}
	createNewUser(user)

	tc := &models.TowatchCategory{
		Value:  "TestTodoCategory",
		Color:  "red",
		UserId: user.ID,
	}
	createTowatchCategory(tc)

	tests := []struct {
		name           string
		body           interface{}
		path           string
		method         string
		wantErr        bool
		wantStatusCode int
	}{
		{
			name:           "Invalid query",
			path:           towatch_category_url + "?user=''",
			wantStatusCode: 400,
			method:         "GET",
		},
		{
			name:           "Invalid user id in query",
			path:           towatch_category_url + "/invalid",
			wantStatusCode: 400,
			method:         "GET",
		},
		{
			name:           "Succesfully get todo categories",
			path:           towatch_category_url + fmt.Sprintf("?user=%v", user.ID),
			wantStatusCode: 200,
			method:         "GET",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body bytes.Buffer
			json.NewEncoder(&body).Encode(tt.body)
			r := httptest.NewRequest(tt.method, tt.path, &body)
			resp, _ := MockApp.app.Test(r, -1)

			if resp.StatusCode != tt.wantStatusCode {
				t.Errorf("GetAllTowatchCategories() want code = %v, got %v", tt.wantStatusCode, resp.StatusCode)
			}
		})
	}

	deleteTowatchCategory(tc)
	deleteUser(user)
}

func TestGetTowatchCategoryById(t *testing.T) {
	user := &models.User{
		Name:     "TestName",
		Email:    "exists@mail.ru",
		Password: utils.HashPassword([]byte("Password123")),
	}
	createNewUser(user)

	tc := &models.TowatchCategory{
		Value:  "TestTodoCategory",
		Color:  "red",
		UserId: user.ID,
	}
	createTowatchCategory(tc)

	tests := []struct {
		name           string
		body           interface{}
		path           string
		method         string
		wantErr        bool
		wantStatusCode int
	}{
		{
			name:           "Invalid params",
			path:           towatch_category_url + "/invalid",
			wantStatusCode: 400,
			method:         "GET",
		},
		{
			name:           "Todo category does not exist",
			path:           towatch_category_url + "/100",
			wantStatusCode: 404,
			method:         "GET",
		},
		{
			name:           "Succesfully get todo category",
			path:           towatch_category_url + fmt.Sprintf("/%v", tc.ID),
			wantStatusCode: 200,
			method:         "GET",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body bytes.Buffer
			json.NewEncoder(&body).Encode(tt.body)

			r := httptest.NewRequest(tt.method, tt.path, &body)
			resp, _ := MockApp.app.Test(r, -1)

			if resp.StatusCode != tt.wantStatusCode {
				t.Errorf("GetTowatchCategoryById() want code = %v, got %v", tt.wantStatusCode, resp.StatusCode)
			}
		})
	}

	deleteTowatchCategory(tc)
	deleteUser(user)
}

func TestCreateTowatchCategory(t *testing.T) {
	user := &models.User{
		Name:     "TestName",
		Email:    "exists@mail.ru",
		Password: utils.HashPassword([]byte("Password123")),
	}
	createNewUser(user)

	tests := []struct {
		name    string
		body    interface{}
		wantErr bool
	}{
		{
			name: "Invalid body",
			body: struct {
				Gender string `json:"gender"`
			}{
				Gender: "Male"},
			wantErr: true,
		},
		{
			name: "Succesfully create",
			body: &dto.CreateTowatchCategoryDto{
				Value:  "TestValue",
				Color:  "blue",
				UserId: user.ID,
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

			err := controllers.CreateTowatchCategory(MockApp.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTowatchCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	deleteUser(user)
}

func TestDeleteTowatchCategoryById(t *testing.T) {
	user := &models.User{
		Name:     "TestName",
		Email:    "exists@mail.ru",
		Password: utils.HashPassword([]byte("Password123")),
	}
	createNewUser(user)

	tc := &models.TowatchCategory{
		Value:  "TestTodoCategory",
		Color:  "red",
		UserId: user.ID,
	}
	createTowatchCategory(tc)

	tests := []struct {
		name           string
		body           interface{}
		path           string
		method         string
		wantErr        bool
		wantStatusCode int
	}{
		{
			name:           "Invalid params",
			path:           towatch_category_url + "/invalid",
			wantStatusCode: 400,
			method:         "DELETE",
		},
		{
			name:           "Towatch category does not exist",
			path:           towatch_category_url + "/100",
			wantStatusCode: 404,
			method:         "DELETE",
		},
		{
			name:           "Succesfully delete towatch category",
			path:           towatch_category_url + fmt.Sprintf("/%v", tc.ID),
			wantStatusCode: 200,
			method:         "DELETE",
		},
	}

	t.Log(tc)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body bytes.Buffer
			json.NewEncoder(&body).Encode(tt.body)

			r := httptest.NewRequest(tt.method, tt.path, &body)
			resp, _ := MockApp.app.Test(r, -1)

			if resp.StatusCode != tt.wantStatusCode {
				t.Errorf("DeleteTowatchCategoryById() want = %v, got %v", tt.wantStatusCode, resp.StatusCode)
			}
		})
	}

	deleteTowatchCategory(tc)
	deleteUser(user)
}

func createTowatchCategory(tc *models.TowatchCategory) {
	database.DB.Create(&tc)
}
func deleteTowatchCategory(tc *models.TowatchCategory) {
	database.DB.Delete(&tc)
}
