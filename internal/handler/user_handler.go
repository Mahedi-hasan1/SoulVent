package handler

// import (
// 	"SoulVent/internal/model"
// 	"SoulVent/internal/repository/user"
// 	"encoding/json"
// 	"net/http"
// )

// var UserRepo user.UserRepository

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	var u model.User
// 	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
// 		http.Error(w, "Invalid request", http.StatusBadRequest)
// 		return
// 	}
// 	if err := UserRepo.Create(&u); err != nil {
// 		http.Error(w, "Failed to create user", http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(u)
// }

// func GetUserProfile(w http.ResponseWriter, r *http.Request) {
// 	// Handle user profile retrieval
// }
