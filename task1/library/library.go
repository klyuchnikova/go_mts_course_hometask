package library

import (
	"hash/fnv"
)

type (
	BookHasher func(string) int
)

type BookLibrary interface {
	Append(book *Book) error
	Search(title string) (*Book, bool)
	SetHasher(hasher BookHasher)
	SetBookStore(bookstore BookStore)
}

type booklibrary struct {
	bookStore BookStore
	hasher    BookHasher
}

func NewLibrary() BookLibrary {
	defaultHasher := func(title string) int {
		h := fnv.New32a()
		h.Write([]byte(title))
		return int(h.Sum32())
	}
	return &booklibrary{
		bookStore: NewBookStore(),
		hasher:    defaultHasher,
	}
}

func (l *booklibrary) Append(book *Book) error {
	bookId := l.hasher(book.Title)
	return l.bookStore.Append(bookId, book)
}

func (l *booklibrary) Search(title string) (*Book, bool) {
	bookId := l.hasher(title)
	return l.bookStore.Search(bookId)
}

func (l *booklibrary) SetHasher(hasher BookHasher) {
	l.hasher = hasher
}

func (l *booklibrary) SetBookStore(bookstore BookStore) {
	l.bookStore = bookstore
}
