package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ReleaseDate string `json:"release_date"`
	Pages int `json:"pages"`
}

var books = []Book{
	Book{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", ReleaseDate: "1979-10-12", Pages: 224},
	Book{Title: "1984", Author: "George Orwell", ReleaseDate: "1949-06-08", Pages: 328},
	Book{Title: "The War of the Worlds", Author: "H. G. Wells", ReleaseDate: "1898-01-01", Pages: 192},
	Book{Title: "Brave New World", Author: "Aldous Huxley", ReleaseDate: "1932-01-01", Pages: 288},
	Book{Title: "Foundation", Author: "Isaac Asimov", ReleaseDate: "1951-01-01", Pages: 255},
	Book{Title: "The Martian Chronicles", Author: "Ray Bradbury", ReleaseDate: "1950-01-01", Pages: 222},
	Book{Title: "Pride and Prejudice", Author: "Jane Austen", ReleaseDate: "1813-01-28", Pages: 432},
}

// entry points
func main() {
	router := gin.Default()	// create a new Gin router
	router.GET("/books", getBooks)	// register the GET handler for /books route
    router.POST("/books", postBook)
	router.GET("/books/:title", getBookByTitle)	// register the GET handler for /books/:title route
	// router.GET("/books/:author", getBookByAuthor)	// register the GET handler for /books/:author route

	router.Run("localhost:8080")	// start the server
}

// GET /books/:title
func getBookByTitle(c *gin.Context) {
	title := c.Param("title")

	for _, a := range books {
		if a.Title == title {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// GET /books/:author
func getBookByAuthor(c *gin.Context) {
	author := c.Param("author")

	for _, a := range books {
		if a.Author == author {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// GET /books
// gin context is the object that provides access to the HTTP request and response objects for your handler function
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)	// serialize the books slice to JSON and send it back in the response
}


func postBook(c *gin.Context) {
	var newBook Book

	// Call BindJSON to bind the received JSON to newBook
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new book to the slice.
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}


