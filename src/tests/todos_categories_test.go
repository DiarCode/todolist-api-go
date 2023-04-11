package tests

import (
	"github.com/DiarCode/todo-go-api/src/database"
	"github.com/DiarCode/todo-go-api/src/models"
)

const (
	todos_category_url = "/api/v1/todo-category"
)

// func TestGetAllTodoCategories(t *testing.T) {
// 	user := &models.User{
// 		Name:     "TestName",
// 		Email:    "exists@mail.ru",
// 		Password: utils.HashPassword([]byte("Password123")),
// 	}
// 	createNewUser(user)

// 	tc := &models.TodoCategory{
// 		Value:  "TestTodoCategory",
// 		Color:  "red",
// 		UserId: user.ID,
// 	}
// 	createTodoCategory(tc)

// 	tests := []struct {
// 		name           string
// 		body           interface{}
// 		path           string
// 		method         string
// 		wantErr        bool
// 		wantStatusCode int
// 	}{
// 		{
// 			name:           "Invalid query",
// 			path:           todos_category_url + "?user=''",
// 			wantStatusCode: 400,
// 			method:         "GET",
// 		},
// 		{
// 			name:           "Invalid user id in query",
// 			path:           todos_category_url + "/invalid",
// 			wantStatusCode: 400,
// 			method:         "GET",
// 		},
// 		{
// 			name:           "Succesfully get todo categories",
// 			path:           todos_category_url + fmt.Sprintf("?user=%v", user.ID),
// 			wantStatusCode: 200,
// 			method:         "GET",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var body bytes.Buffer
// 			err := json.NewEncoder(&body).Encode(tt.body)
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			r := httptest.NewRequest(tt.method, tt.path, &body)
// 			resp, _ := MockApp.app.Test(r, -1)

// 			if resp.StatusCode != tt.wantStatusCode {
// 				t.Errorf("GetAllTodoCategories() want = %v, got %v", tt.wantStatusCode, resp.StatusCode)
// 			}
// 		})
// 	}

// 	deleteTodoCategory(tc)
// 	deleteUser(user)
// }

// func TestGetTodoCategoryById(t *testing.T) {
// 	user := &models.User{
// 		Name:     "TestName",
// 		Email:    "exists@mail.ru",
// 		Password: utils.HashPassword([]byte("Password123")),
// 	}
// 	createNewUser(user)

// 	tc := &models.TodoCategory{
// 		Value:  "TestTodoCategory",
// 		Color:  "red",
// 		UserId: user.ID,
// 	}
// 	createTodoCategory(tc)

// 	tests := []struct {
// 		name           string
// 		body           interface{}
// 		path           string
// 		method         string
// 		wantErr        bool
// 		wantStatusCode int
// 	}{
// 		{
// 			name:           "Invalid params",
// 			path:           todos_category_url + "/invalid",
// 			wantStatusCode: 400,
// 			method:         "GET",
// 		},
// 		{
// 			name:           "Todo category does not exist",
// 			path:           todos_category_url + "/100",
// 			wantStatusCode: 404,
// 			method:         "GET",
// 		},
// 		{
// 			name:           "Succesfully get todo category",
// 			path:           todos_category_url + fmt.Sprintf("/%v", tc.ID),
// 			wantStatusCode: 200,
// 			method:         "GET",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var body bytes.Buffer
// 			err := json.NewEncoder(&body).Encode(tt.body)
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			r := httptest.NewRequest(tt.method, tt.path, &body)
// 			resp, _ := MockApp.app.Test(r, -1)

// 			if resp.StatusCode != tt.wantStatusCode {
// 				t.Errorf("GetTodoCategoryById() want = %v, got %v", tt.wantStatusCode, resp.StatusCode)
// 			}
// 		})
// 	}

// 	deleteTodoCategory(tc)
// 	deleteUser(user)
// }

// func TestCreateTodoCategory(t *testing.T) {
// 	user := &models.User{
// 		Name:     "TestName",
// 		Email:    "exists@mail.ru",
// 		Password: utils.HashPassword([]byte("Password123")),
// 	}
// 	createNewUser(user)

// 	tests := []struct {
// 		name    string
// 		body    interface{}
// 		wantErr bool
// 	}{
// 		{
// 			name: "Invalid body",
// 			body: struct {
// 				Gender string `json:"gender"`
// 			}{
// 				Gender: "Male"},
// 			wantErr: true,
// 		},
// 		{
// 			name: "User id that does not exists",
// 			body: &dto.CreateTodoCategoryDto{
// 				Value:  "TestValue",
// 				Color:  "blue",
// 				UserId: 100,
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "Succesfully create",
// 			body: &dto.CreateTodoCategoryDto{
// 				Value:  "TestValue",
// 				Color:  "blue",
// 				UserId: user.ID,
// 			},
// 			wantErr: false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			jsonBytes, _ := json.Marshal(tt.body)

// 			MockApp.ctx.Request().Reset()
// 			MockApp.ctx.Request().Header.Set("Content-Type", "application/json")
// 			MockApp.ctx.Request().SetBodyRaw(jsonBytes)

// 			err := controllers.CreateTodoCategory(MockApp.ctx)

// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("CreateTodoCategory() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}

// 	deleteUser(user)
// }

// func TestDeleteTodoCategoryById(t *testing.T) {
// 	user := &models.User{
// 		Name:     "TestName",
// 		Email:    "exists@mail.ru",
// 		Password: utils.HashPassword([]byte("Password123")),
// 	}
// 	createNewUser(user)

// 	tc := &models.TodoCategory{
// 		Value:  "TestTodoCategory",
// 		Color:  "red",
// 		UserId: user.ID,
// 	}
// 	createTodoCategory(tc)

// 	tests := []struct {
// 		name           string
// 		body           interface{}
// 		path           string
// 		method         string
// 		wantErr        bool
// 		wantStatusCode int
// 	}{
// 		{
// 			name:           "Invalid params",
// 			path:           todos_category_url + "/invalid",
// 			wantStatusCode: 400,
// 			method:         "DELETE",
// 		},
// 		{
// 			name:           "Todo category does not exist",
// 			path:           todos_category_url + "/100",
// 			wantStatusCode: 404,
// 			method:         "DELETE",
// 		},
// 		{
// 			name:           "Succesfully delete todo category",
// 			path:           todos_category_url + fmt.Sprintf("/%v", tc.ID),
// 			wantStatusCode: 200,
// 			method:         "DELETE",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var body bytes.Buffer
// 			err := json.NewEncoder(&body).Encode(tt.body)
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			r := httptest.NewRequest(tt.method, tt.path, &body)
// 			resp, _ := MockApp.app.Test(r, -1)

// 			if resp.StatusCode != tt.wantStatusCode {
// 				t.Errorf("DeleteTodoCategoryById() want = %v, got %v", tt.wantStatusCode, resp.StatusCode)
// 			}
// 		})
// 	}

// 	deleteTodoCategory(tc)
// 	deleteUser(user)
// }

func createTodoCategory(tc *models.TodoCategory) {
	database.DB.Create(&tc)
}
func deleteTodoCategory(tc *models.TodoCategory) {
	database.DB.Delete(&tc)
}
