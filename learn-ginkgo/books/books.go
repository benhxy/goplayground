package books

import "encoding/json"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (book *Book) CategoryByLength() string {
	if book.Pages <= 300 {
		return "Short story"
	} else {
		return "Novel"
	}
}

func NewBookFromJSON(str string) (Book, error) {
	var book Book
	err := json.Unmarshal([]byte(str), &book)
	return book, err
}
