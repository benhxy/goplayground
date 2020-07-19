package books

type Book struct{
	Author string
	PageNumber int
	Description string
}

func (book *Book) IsZero() bool {
	return book.Author == "" && book.PageNumber == 0 && book.Description == ""
}
