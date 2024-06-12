package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Usuario representa la estructura de un usuario
type Usuario struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      string `json:"age"`
	Location string `json:"location"`
}

// usuarios es una lista de usuarios
var usuarios []Usuario

// createUser crea un nuevo usuario
func createUser(w http.ResponseWriter, r *http.Request) {
	var usuario Usuario
	_ = json.NewDecoder(r.Body).Decode(&usuario)
	usuario.ID = len(usuarios) + 1
	usuarios = append(usuarios, usuario)
	json.NewEncoder(w).Encode(usuario)
}

// updateUser modifica un usuario existente
func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, usuario := range usuarios {
		if usuario.ID == id {
			usuarios = append(usuarios[:i], usuarios[i+1:]...)
			var updatedUser Usuario
			_ = json.NewDecoder(r.Body).Decode(&updatedUser)
			updatedUser.ID = id
			usuarios = append(usuarios, updatedUser)
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
	json.NewEncoder(w).Encode(struct{ Error string }{"Usuario no encontrado"})
}

// deleteUser elimina un usuario existente
func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, usuario := range usuarios {
		if usuario.ID == id {
			usuarios = append(usuarios[:i], usuarios[i+1:]...)
			json.NewEncoder(w).Encode(struct{ Message string }{"Usuario eliminado"})
			return
		}
	}
	json.NewEncoder(w).Encode(struct{ Error string }{"Usuario no encontrado"})
}

// getUsers obtiene todos los usuarios
func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(usuarios)
}

// getUserByID obtiene un usuario por su ID
func getUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, usuario := range usuarios {
		if usuario.ID == id {
			json.NewEncoder(w).Encode(usuario)
			return
		}
	}
	json.NewEncoder(w).Encode(struct{ Error string }{"Usuario no encontrado"})
}

func main() {
	router := mux.NewRouter()

	// Endpoints
	/*	router.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello World")
	})*/

	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUserByID).Methods("GET")

	fmt.Println("Servidor iniciado en http://localhost:8000")
	http.ListenAndServe(":8000", router)
}
