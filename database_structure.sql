CREATE TABLE Books (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255),
    publish_date VARCHAR(50),
    genre VARCHAR(100),
    edition VARCHAR(50),
    publisher VARCHAR(255),
    description TEXT,
    CONSTRAINT unique_book_uuid UNIQUE (book_uuid)
);

CREATE TABLE Collections (
    name VARCHAR(255) PRIMARY KEY,
    description TEXT
);

CREATE TABLE CollectionBooks (
    collection_uuid UUID REFERENCES Collections(collection_uuid),
    book_uuid UUID REFERENCES Books(book_uuid),
    PRIMARY KEY (collection_uuid, book_uuid)
);
