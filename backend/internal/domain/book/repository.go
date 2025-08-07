package book

type Repository interface {
	AddBook(book *Book)
	GetAllBooks() []Book
	GetBookByID(id string) Book
	UpdateBook(book *Book)
	DeleteBook(id string)
}
