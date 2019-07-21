package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path"
	"time"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/book/", bookHandler)
	handler.HandleFunc("/books", LoggerMiddelware(booksHandler))
	handler.HandleFunc("/addBook", LoggerMiddelware(addBookHandler))

	server := http.Server{
		Addr:           ":8081",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}

type Resp struct {
	Message interface{}
	Error   string
}

type Book struct {
	Id     string `json:"id"`
	Author string `json:"author"`
	Name   string `json:"name"`
}

type BookStore struct {
	books []Book
}

var bookStore = BookStore{
	books: make([]Book, 0),
}

//Get book by id function
func (bs *BookStore) getBooksById(id string) *Book {
	for _, bk := range bs.books {
		if bk.Id == id {
			return &bk
		}
	}

	return nil
}

//Get all books
func (bs *BookStore) getBooks() []Book {
	return bs.books
}

//Update book function
func (bs *BookStore) updateBook(book Book) error {
	for i, bk := range bs.books {
		if bk.Id == book.Id {
			bs.books[i] = book

			return nil
		}
	}

	return errors.New(fmt.Sprintf("Book with id: %s not found", book.Id))
}

//Add books function
func (bs *BookStore) addBook(book Book) error {
	for _, bk := range bs.books {
		if bk.Id == book.Id {
			return errors.New(fmt.Sprintf("Failed to add new book with %s", book.Id))
		}
	}

	bs.books = append(bs.books, book)

	return nil
}

//Delete function
func (bs *BookStore) delete(id string) error {
	for i, bk := range bs.books {
		if bk.Id == id {
			bs.books = append(bs.books[:i], bs.books[i+1:]...)
		}

		return nil
	}

	return errors.New(fmt.Sprintf("Book with id: %s not found", id))
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	urlParam := path.Base(r.URL.Path)
	resp := Resp{}

	if urlParam == "book" {
		w.WriteHeader(http.StatusBadRequest)
		resp.Error = "Wrong request"
	} else {
		w.WriteHeader(http.StatusOK)
		resp.Message = urlParam
	}

	respJson, _ := json.Marshal(resp)

	w.Write(respJson)
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	/*	resp := Resp {
			Message: "Sorry, there are no books yet",
		}
		respJson, _ := json.Marshal(resp)*/

	respJson, _ := json.Marshal(bookStore.getBooks())

	w.WriteHeader(http.StatusOK)

	w.Write(respJson)
}

func addBookHandler(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var book Book
	var resp Resp

	err := decoder.Decode(&book)

	if err != nil {
		resp.Error = err.Error()
		respJson, _ := json.Marshal(resp)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(respJson)

		return
	}

	err = bookStore.addBook(book)

	if err != nil {
		resp.Error = err.Error()
		respJson, _ := json.Marshal(resp)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(respJson)

		return
	}

	respJson, _ := json.Marshal(book)
	writer.Write(respJson)
}

//Middelware
func LoggerMiddelware(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		fmt.Printf("\nConnection [net/http] request method [%s] connection from [%v]", request.Method, request.RemoteAddr)

		next.ServeHTTP(writer, request)
	}
}
