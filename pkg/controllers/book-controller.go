package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	m "github.com/zigamedved/bookstore-management-API/pkg/models"
	"github.com/zigamedved/bookstore-management-API/pkg/utils"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(m.GetAllBooks())
	if err != nil {
		fmt.Println("error while marshaling books")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	if bookId == "" {
		fmt.Println("missing url parameter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing bookId")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	book, _ := m.GetBookById(Id)
	res, err := json.Marshal(book)
	if err != nil {
		fmt.Println("error while marshaling book")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book m.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		fmt.Println("error while decoding request")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	m.CreateBook(&book)
	res, err := json.Marshal(book)
	if err != nil {
		fmt.Println("error while marshaling book")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	if bookId == "" {
		fmt.Println("missing url parameter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing bookId")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	m.DeleteBookById(Id)
	w.WriteHeader(http.StatusNoContent)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	if bookId == "" {
		fmt.Println("missing url parameter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing bookId")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var book = m.Book{}
	utils.ParseBody(r, &book)
	m.UpdateBookById(Id, &book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	w.Write(nil)
}
