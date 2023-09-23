package main

import (
	"fmt"
	"os"
	"task1/library"
)

func Assert(condition bool, message string) {
	if !condition {
		fmt.Println(message)
		os.Exit(1)
	}
}

func main() {
	secondTitle := "Title Two"
	books := []*library.Book{
		library.NewBook("Name 1", "Title 1", "Contents 1"),
		library.NewBook("Name Two", secondTitle, "Contents Two"),
		library.NewBook("Name Three", "Title Three", "Contents Three"),
	}
	var bookLibrary library.BookLibrary = library.NewLibrary()
	var err error

	for _, book := range books {
		err = bookLibrary.Append(book)
		Assert(err == nil, "library failed to append a book titled: "+book.Title)
	}

	foundBook, hasFound := bookLibrary.Search(secondTitle)
	Assert(hasFound == true, "The book with the name is not found: "+secondTitle)
	Assert(foundBook.Title == secondTitle, "The found book has wrong title")
	Assert(foundBook.Author == "Name Two", "The found book has wrong author")

	newHasher := func(title string) int {
		// a real bad example
		return len(title)
	}
	bookLibrary.SetHasher(newHasher)
	_, hasFound = bookLibrary.Search(secondTitle)
	Assert(hasFound == false, "The book after changed hasher is found: "+secondTitle)

	bookLibrary.SetBookStore(library.NewBookStore())
	for _, book := range books {
		err = bookLibrary.Append(book)
		Assert(err == nil, "library failed to append a book titled: "+book.Title)
	}

	foundBook, hasFound = bookLibrary.Search(secondTitle)
	Assert(hasFound == true, "After changed store the book with the name is not found: "+secondTitle)
	Assert(foundBook.Title == secondTitle, "After changed store the found book has wrong title")
	Assert(foundBook.Author == "Name Two", "After changed store the found book has wrong author")

	fmt.Println("Program finished! Test run passed")
}
