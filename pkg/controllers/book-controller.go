package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sunilym2277/go_BookStore/pkg/models"
	"github.com/sunilym2277/go_BookStore/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	NewBooks := models.GetAllBook()
	res, _ := json.Marshal(NewBooks)
	w.Header().Set("Content-Type", "pkglicaton/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	BookId := vars["bookId"]
	Id, err := strconv.ParseInt(BookId, 0, 0)
	if err != nil {
		fmt.Println("error in Parsing")
	}
	Bookdetails, _ := models.GetBookById(Id)
	res, _ := json.Marshal(Bookdetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["BookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error in Parseing the ID")
	}
	book := models.DeleteBook(Id)

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	BookId := vars["bookId"]
	Id, err := strconv.ParseInt(BookId, 0, 0)
	if err != nil {
		fmt.Println("Error in parsing the data")
	}
	BookDetails, db := models.GetBookById(Id)
	if updateBook.Name != "" {
		BookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		BookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		BookDetails.Publication = updateBook.Publication
	}
	db.Save(&BookDetails)

	res, _ := json.Marshal(BookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
