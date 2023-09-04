CREATE TABLE Book (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    genre VARCHAR(100) NOT NULL,
    author VARCHAR(255),
    publish_date VARCHAR(50),
    edition VARCHAR(50),
    publisher VARCHAR(255),
    description TEXT,
    CONSTRAINT unique_book_uuid UNIQUE (id)
);

CREATE TABLE Collection (
    id VARCHAR(255) PRIMARY KEY,
    description TEXT,
    CONSTRAINT unique_collection_name UNIQUE (id)
);

CREATE TABLE CollectionBook (
    collection_name VARCHAR(255) REFERENCES Collection(id),
    book_uuid UUID REFERENCES Book(id),
    PRIMARY KEY (collection_name, book_uuid)
);
