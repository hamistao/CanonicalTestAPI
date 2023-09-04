package models

// Represents a Book in the Virtual Library.
type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Edition     string `json:"edition"`
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
	Title      string
	Collection string
	Author     string
	Genre      string
	Publisher  string
	Edition    string
	From       string
	To         string
	Max        int64
}
