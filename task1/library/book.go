package library

type Book struct {
	Author   string
	Title    string
	Contents string
}

func NewBook(author string, title string, contents string) *Book {
	return &Book{
		Author:   author,
		Title:    title,
		Contents: contents,
	}
}

func (b *Book) Read() *string {
	return &b.Contents
}
