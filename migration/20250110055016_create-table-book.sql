-- +migrate Up
CREATE TABLE books (
    id INT AUTO_INCREMENT PRIMARY KEY,
     title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    category_id INT NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories (id)
);

-- +migrate Down
DROP TABLE IF EXISTS books;
