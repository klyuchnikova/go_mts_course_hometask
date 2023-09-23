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
		library.NewBook("Name 4", "Title four 4", "Contents 4"),
		library.NewBook("Name 5", "Title five   5", "Contents 5"),
	}

	var bookLibraryOnMap library.BookLibrary = library.NewLibraryOnMap()
	fmt.Println("testing library with map storage")
	testLibrary(bookLibraryOnMap, books)

	var bookLibraryOnSlice library.BookLibrary = library.NewLibraryOnSlice()
	fmt.Println("testing library with slice storage")
	testLibrary(bookLibraryOnSlice, books)

	fmt.Println("Program finished! Test run passed")
}

func testLibrary(bookLibrary library.BookLibrary, books []*library.Book) {
	var err error
	for _, book := range books {
		err = bookLibrary.Append(book)
		Assert(err == nil, "library failed to append a book titled: "+book.Title)
	}

	for i := 1; i <= 3; i++ {
		title := books[i].Title
		foundBook, hasFound := bookLibrary.Search(title)
		Assert(hasFound == true, "The book with the name is not found: "+title)
		Assert(foundBook.Title == title, "The found book has wrong title: "+foundBook.Title+" != "+title)
		Assert(foundBook.Author == books[i].Author, "The found book has wrong author")

	}

	newHasher := func(title string) int {
		// a real bad example
		return len(title) + 100
	}
	bookLibrary.SetHasher(newHasher)

	secondTitle := books[1].Title
	_, hasFound := bookLibrary.Search(secondTitle)
	Assert(hasFound == false, "The book after changed hasher is found: "+secondTitle)
}
