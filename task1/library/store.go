package library

type BookStore interface {
	Append(id int, book *Book) error
	Search(id int) (*Book, bool)
}

type storeOnMap struct {
	books map[int]*Book
}

func NewBookStoreOnMap() BookStore {
	return &storeOnMap{
		books: make(map[int]*Book),
	}
}

func (s *storeOnMap) Append(id int, book *Book) error {
	s.books[id] = book
	return nil
}

func (s *storeOnMap) Search(id int) (*Book, bool) {
	book, isFound := s.books[id]
	return book, isFound
}

type storeOnSlice struct {
	books []*Book
}

func NewBookStoreOnSlice() BookStore {
	return &storeOnSlice{}
}

func (s *storeOnSlice) Append(id int, book *Book) error {
	if id < len(s.books) {
		s.books[id] = book
	} else {
		new_books := make([]*Book, max(len(s.books)*2, id+1))
		for i := 0; i < len(s.books); i++ {
			new_books[i] = s.books[i]
		}
		new_books[id] = book
		s.books = new_books
	}
	return nil
}

func (s *storeOnSlice) Search(id int) (*Book, bool) {
	if id < len(s.books) {
		return s.books[id], true
	}
	return nil, false
}
