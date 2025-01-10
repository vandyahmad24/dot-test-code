-- +migrate Up
INSERT INTO books (title, author, category_id) VALUES
('Book Title 1', 'Author 1', 1),
('Book Title 2', 'Author 2', 2),
('Book Title 3', 'Author 3', 3),
('Book Title 4', 'Author 4', 4),
('Book Title 5', 'Author 5', 5),
('Book Title 6', 'Author 6', 6),
('Book Title 7', 'Author 7', 7),
('Book Title 8', 'Author 8', 8),
('Book Title 9', 'Author 9', 9),
('Book Title 10', 'Author 10', 10),
('Book Title 11', 'Author 11', 1),
('Book Title 12', 'Author 12', 2),
('Book Title 13', 'Author 13', 3),
('Book Title 14', 'Author 14', 4),
('Book Title 15', 'Author 15', 5),
('Book Title 16', 'Author 16', 6),
('Book Title 17', 'Author 17', 7),
('Book Title 18', 'Author 18', 8),
('Book Title 19', 'Author 19', 9),
('Book Title 20', 'Author 20', 10);

-- +migrate Down
truncate table books;