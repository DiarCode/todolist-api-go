package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var user1 = User{
	Id:       1,
	Name:     "Diar",
	Username: "diarCode",
	Password: "Reraha22",
	Email:    "diar@mail.ru",
}

var user2 = User{
	Id:       2,
	Name:     "Dana",
	Username: "danaBeauty",
	Password: "Krutaya22",
	Email:    "dana@gmail.com",
}

var users = []User{user1, user2}

func GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	json.NewEncoder(w).Encode(users)
}

func GetUserById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))

	if err != nil {
		w.Write([]byte(err.Error()))
	}

	for _, candidate := range users {
		if candidate.Id == id {
			json.NewEncoder(w).Encode(candidate)
			return
		}
	}

	w.Write([]byte(`{"error:" + "No such user"}`))
}

func AddUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var u User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, candidate := range users {
		if candidate.Id == u.Id {
			http.Error(w, "Such user already exists", http.StatusBadRequest)
			return
		}
	}

	users = append(users, u)
	fmt.Fprintf(w, "Person: %+v", u)
}
