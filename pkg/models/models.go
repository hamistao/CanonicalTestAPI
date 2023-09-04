package models

// Represents a Book in the Virtual Library.
type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Edition     string `json:"edition"`
	Genre       string `json:"genre"`
	Publisher   string `json:"publisher"`
	PublishDate string `json:"pub_date"`
}

// Represents a collection of books.
type Collection struct {
	Id          string `json:"name"`
	Description string `json:"description"`
}

// Used as a filter to query books with.
type QueryFilter struct {
	TitleFilter      string `json:"titleFilter"`
	CollectionFilter string `json:"collection"`
	AuthorFilter     string `json:"authorFilter"`
	GenreFilter      string `json:"genreFilter"`
	PubFilter        string `json:"pubFilter"`
	EditionFilter    string `json:"editionFilter"`
	From             string `json:"from"`
	To               string `json:"to"`
	Max              int64  `json:"max"`
}
