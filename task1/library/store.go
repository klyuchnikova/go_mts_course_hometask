package library

type BookStore interface {
	Append(id int, book *Book) error
	Search(id int) (*Book, bool)
}

type bookStore struct {
	books map[int]*Book
}

func NewBookStore() BookStore {
	return &bookStore{
		books: make(map[int]*Book),
	}
}

func (s *bookStore) Append(id int, book *Book) error {
	s.books[id] = book
	return nil
}

func (s *bookStore) Search(id int) (*Book, bool) {
	book, isFound := s.books[id]
	return book, isFound
}
