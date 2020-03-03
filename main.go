package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Libro struct {
	ID     string `json:"id"`
	Titulo string `json:"titulo"`
	Autor  *Autor `json:"autor"`
}

type Autor struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

var libros []Libro

// Retornar todos los libros
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(libros)
}

// Retornar un libro por id
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range libros {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Libro{})
}

// crea un nuevo libro
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var libro Libro
	_ = json.NewDecoder(r.Body).Decode(&libro)
	libros = append(libros, libro)
	json.NewEncoder(w).Encode(libro)
}

// Actualiza un libro
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range libros {
		if item.ID == params["id"] {
			libros = append(libros[:index], libros[index+1:]...)
			var libro Libro
			_ = json.NewDecoder(r.Body).Decode(&libro)
			libro.ID = params["id"]
			libros = append(libros, libro)
			json.NewEncoder(w).Encode(libro)
			return
		}
	}
}

// Elimina un libro
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range libros {
		if item.ID == params["id"] {
			libros = append(libros[:index], libros[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(libros)
}

func main() {
	// Router inicial
	r := mux.NewRouter()

	libros = append(libros, Libro{ID: "1", Titulo: "Harry Potter", Autor: &Autor{Nombre: "Joanne", Apellido: "Rowling​"}})
	libros = append(libros, Libro{ID: "2", Titulo: "El Arte de la guerra", Autor: &Autor{Nombre: "Sun", Apellido: "Tzu"}})
	libros = append(libros, Libro{ID: "3", Titulo: "El Señor de los anillos", Autor: &Autor{Nombre: "J.R.R", Apellido: "Tolkien"}})
	libros = append(libros, Libro{ID: "4", Titulo: "El resplandor", Autor: &Autor{Nombre: "Stephen", Apellido: "King"}})

	// Route handles y endpoints
	r.HandleFunc("/libros", getBooks).Methods("GET")
	r.HandleFunc("/libros/{id}", getBook).Methods("GET")
	r.HandleFunc("/libros", createBook).Methods("POST")
	r.HandleFunc("/libros/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/libros/{id}", deleteBook).Methods("DELETE")

	// este método lo que hace es recibir la peticion y mostrar el contenido, lo hace en el puerto 8080
	log.Fatal(http.ListenAndServe(":8080", r))
}
