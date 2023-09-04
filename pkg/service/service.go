package service

import (
	"canonicalTestAPI/pkg/models"
	"os/exec"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

type Service struct {
	DB *dbx.DB
}

// retrieves a book from the database by its ID.
func (sv *Service) GetBook(id string) (models.Book, error) {
	query := sv.DB.Select("*").From("book").Where(dbx.HashExp{"id": id})

	var book models.Book
	err := query.One(&book)

	return book, err
}

// retrieves a collection from the database by its name.
func (sv *Service) GetCollection(name string) (models.Collection, error) {
	query := sv.DB.Select("*").From("collection").Where(dbx.HashExp{"id": name})

	var collection models.Collection
	err := query.One(&collection)

	return collection, err
}

// inserts a new book into the database.
func (sv *Service) InsertBook(book models.Book) (string, error) {
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		return "", err
	}
	id := string(uuid)
	book.ID = id[:len(id)-1]

	if err := sv.DB.Model(&book).Insert(); err != nil {
		return "", err
	}

	return book.ID, nil
}

// inserts a new collection into the database.
func (sv *Service) InsertCollection(collection models.Collection) error {
	if err := sv.DB.Model(&collection).Insert(); err != nil {
		return err
	}

	return nil
}

// adds a book to a collection.
func (sv *Service) Collect(collectionName, bookID string) error {
	_, err := sv.DB.Insert("collectionbook", dbx.Params{"collection_name": collectionName, "book_uuid": bookID}).Execute()
	return err
}

// removes a book from a collection.
func (sv *Service) Discard(collectionName, bookID string) error {
	_, err := sv.DB.Delete("collectionbook", dbx.HashExp{"collection_name": collectionName, "book_uuid": bookID}).Execute()
	return err
}

func (sv *Service) BelongsTo(collectionName, bookID string) bool {
	var out dbx.NullStringMap
	err := sv.DB.Select("*").From("collectionbook").Where(dbx.HashExp{"collection_name": collectionName}).AndWhere(dbx.HashExp{"book_uuid": bookID}).One(&out)
	return err == nil
}

// deletes a book from the database.
func (sv *Service) DeleteBook(bookID string) error {
	book := models.Book{ID: bookID}
	if _, err := sv.DB.Delete("collectionbook", dbx.HashExp{"book_uuid": bookID}).Execute(); err != nil {
		return err
	}

	if err := sv.DB.Model(&book).Delete(); err != nil {
		return err
	}

	return nil
}

// deletes a collection from the database.
func (sv *Service) DeleteCollection(collectionName string) error {
	collection := models.Collection{Id: collectionName}
	if err := sv.DB.Model(&collection).Delete(); err != nil {
		return err
	}

	return nil
}

// retrieves all collections from the database.
func (sv *Service) GetAllCollections() ([]models.Collection, error) {
	var collections []models.Collection
	err := sv.DB.Select("*").From("collection").All(&collections)

	return collections, err
}

// performs a filtered query for books in the database.
func (sv *Service) Query(filter models.QueryFilter) ([]models.Book, error) {
	var query *dbx.SelectQuery

	if filter.Collection != "" {
		query = sv.DB.Select("b.*").
			From("book b").
			InnerJoin("collectionbook cb", dbx.NewExp("b.id = cb.book_uuid")).
			Where(dbx.HashExp{"cb.collection_name": filter.Collection})
	} else {
		query = sv.DB.Select("*").From("book")
	}

	query = addFilters(query, filter)

	var books []models.Book
	err := query.All(&books)

	return books, err
}

func addFilters(query *dbx.SelectQuery, filter models.QueryFilter) *dbx.SelectQuery {
	if filter.Max > 0 {
		query = query.Limit(filter.Max)
	}
	if filter.Title != "" {
		query = query.AndWhere(dbx.Like("title", filter.Title))
	}
	if filter.Author != "" {
		query = query.AndWhere(dbx.Like("author", filter.Author))
	}
	if filter.Genre != "" {
		query = query.AndWhere(dbx.Like("genre", filter.Genre))
	}
	if filter.Publisher != "" {
		query = query.AndWhere(dbx.Like("publisher", filter.Publisher))
	}
	if filter.Edition != "" {
		query = query.AndWhere(dbx.Like("edition", filter.Edition))
	}
	if filter.From != "" {
		query = query.
			AndWhere(dbx.NewExp("publish_date IS NOT NULL AND publish_date != ''")).
			AndWhere(dbx.NewExp("TO_DATE(publish_date, 'MM-DD-YYYY') >= TO_DATE({:from}, 'MM-DD-YYYY')", dbx.Params{"from": filter.From}))
	}
	if filter.To != "" {
		query = query.
			AndWhere(dbx.NewExp("publish_date IS NOT NULL AND publish_date != ''")).
			AndWhere(dbx.NewExp("TO_DATE(publish_date, 'MM-DD-YYYY') <= TO_DATE({:to}, 'MM-DD-YYYY')", dbx.Params{"to": filter.To}))
	}

	return query
}

// updates book information in the database.
func (sv *Service) UpdateBook(bookID string, updates models.Book) error {
	updateValues := make(dbx.Params)

	switch {
	case updates.Title != "":
		updateValues["title"] = updates.Title
	case updates.Description != "":
		updateValues["description"] = updates.Description
	case updates.Author != "":
		updateValues["author"] = updates.Author
	case updates.Edition != "":
		updateValues["edition"] = updates.Edition
	case updates.Genre != "":
		updateValues["genre"] = updates.Genre
	case updates.Publisher != "":
		updateValues["publisher"] = updates.Publisher
	case updates.PublishDate != "":
		updateValues["publish_date"] = updates.PublishDate
	}

	_, err := sv.DB.Update("book", updateValues, dbx.HashExp{"id": bookID}).Execute()

	return err
}

// updates collection information in the database.
func (sv *Service) UpdateCollection(collectionName string, updates models.Collection) error {
	updateValues := dbx.Params{
		"description": updates.Description,
	}

	condition := dbx.HashExp{"id": collectionName}

	_, err := sv.DB.Update("collection", updateValues, condition).Execute()
	return err
}
