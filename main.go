package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model Buku
type Buku struct {
	idBuku  string   `json:"idBuku"`
	Isdn    string   `json:"Isdn"`
	Judul   string   `json:"Judul"`
	Penulis *Penulis `json:"Penulis"`
}

// Penulis struct
type Penulis struct {
	namaDepan    string `json:"namaDepan"`
	namaBelakang string `json:"namaBelakang"`
}

// Inisialisasi books sebagai potongan dari struct Buku
var books []Buku

// Mendapat semua data buku
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Mendapat data satu buku
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the idBuku from the params
	for _, item := range books {
		if item.idBuku == params["idBuku"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Buku{})
}

// Menambah buku baru
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Buku
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.idBuku = strconv.Itoa(rand.Intn(100000000)) // Mock idBuku - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Memperbarui data buku
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.idBuku == params["idBuku"] {
			books = append(books[:index], books[index+1:]...)
			var book Buku
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.idBuku = params["idBuku"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Menghapus data buku
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.idBuku == params["idBuku"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {

	r := mux.NewRouter()

	// Data hardcode
	books = append(books, Buku{idBuku: "1", Isdn: "8000", Judul: "Mastering Golang", Penulis: &Penulis{namaDepan: "Komar", namaBelakang: "Udin"}})
	books = append(books, Buku{idBuku: "2", Isdn: "8001", Judul: "Learn DevOps", Penulis: &Penulis{namaDepan: "Nata", namaBelakang: "Decoco"}})

	// Route handles & endpoints
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{idBuku}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{idBuku}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{idBuku}", deleteBook).Methods("DELETE")

	// Jalankan server
	log.Fatal(http.ListenAndServe(":8000", r))
}
