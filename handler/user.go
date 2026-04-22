package handler

import (
	"encoding/json"
	"net/http"
)

var users = map[string]string{
	"alice":   "Alice Smith",
	"bob":     "Bob Jones",
	"charlie": "Charlie Brown",
	"dave":    "Dave Wilson",
	"eve":     "Eve Adams",
	"frank":   "Frank Miller",
	"grace":   "Grace Hopper",
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	if name == "" {
		http.Error(w, "Missing name", http.StatusBadRequest)
		return
	}

	fullName, ok := users[name]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"username": name, "full_name": fullName})
}
