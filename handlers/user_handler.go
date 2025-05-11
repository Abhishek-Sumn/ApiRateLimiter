package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ratelimiter/models"
	"ratelimiter/redis"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)
	key := fmt.Sprintf("user:%d", id)
	var user models.User

	cached, err := redis.Client.Get(redis.Ctx, key).Result()
	json.Unmarshal([]byte(cached), &user)

	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		resp := models.Response{
			Success: true,
			Data:    user,
			Message: "User Successfully created",
			Status:  200,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := models.Response{
		Success: false,
		Data:    nil,
		Message: "User not found",
		Status:  400,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

var validate = validator.New()

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(user); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	key := fmt.Sprintf("user:%d", user.ID)
	data, _ := json.Marshal(user)
	redis.Client.Set(redis.Ctx, key, data, 0)

	resp := models.Response{
		Success: true,
		Data:    nil,
		Message: "User Successfully created",
		Status:  200,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
