-- +migrate Up
CREATE TABLE categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS categories;
