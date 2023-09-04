-- Insert sample books with UUIDs
INSERT INTO Book (id, title, genre, author, publish_date, edition, publisher, description)
VALUES
    ('550e8400-e29b-41d4-a716-446655440000', 'Book 1', 'Fiction', 'Author 1', '2022-01-01', 'First Edition', 'Publisher 1', 'Description 1'),
    ('550e8400-e29b-41d4-a716-446655440001', 'Book 2', 'Science Fiction', 'Author 2', '2022-02-01', 'Second Edition', 'Publisher 2', 'Description 2'),
    ('550e8400-e29b-41d4-a716-446655440002', 'Book 3', 'Mystery', 'Author 3', '2022-03-01', 'Third Edition', 'Publisher 3', 'Description 3'),
    ('550e8400-e29b-41d4-a716-446655440003', 'Book 4', 'Fantasy', 'Author 4', '2022-04-01', 'Fourth Edition', 'Publisher 4', 'Description 4'),
    ('550e8400-e29b-41d4-a716-446655440004', 'Book 5', 'Thriller', 'Author 5', '2022-05-01', 'Fifth Edition', 'Publisher 5', 'Description 5');

-- Insert sample collections
INSERT INTO Collection (id, description)
VALUES
    ('Collection 1', 'Collection Description 1'),
    ('Collection 2', 'Collection Description 2');

-- Insert books into collections
INSERT INTO CollectionBook (collection_name, book_uuid)
VALUES
    ('Collection 1', '550e8400-e29b-41d4-a716-446655440000'),
    ('Collection 1', '550e8400-e29b-41d4-a716-446655440001'),
    ('Collection 2', '550e8400-e29b-41d4-a716-446655440001'),
    ('Collection 2', '550e8400-e29b-41d4-a716-446655440002'),
    ('Collection 2', '550e8400-e29b-41d4-a716-446655440003');