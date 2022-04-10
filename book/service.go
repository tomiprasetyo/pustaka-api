package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, bookUpdateRequest BookUpdateRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	if err != nil {
		panic(err)
	}

	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {

	book := Book{
		Title:       bookRequest.Title,
		Price:       bookRequest.Price,
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,
		Discount:    bookRequest.Discount,
	}

	newBook, err := s.repository.Create(book)
	if err != nil {
		panic(err)
	}

	return newBook, err
}

func (s *service) Update(ID int, bookUpdateRequest BookUpdateRequest) (Book, error) {

	book, _ := s.repository.FindByID(ID)

	book.Title = bookUpdateRequest.Title
	book.Price = bookUpdateRequest.Price
	book.Description = bookUpdateRequest.Description
	book.Rating = bookUpdateRequest.Rating
	book.Discount = bookUpdateRequest.Discount

	newBook, err := s.repository.Update(book)
	if err != nil {
		panic(err)
	}

	return newBook, err
}
