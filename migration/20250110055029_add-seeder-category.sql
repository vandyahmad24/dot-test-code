-- +migrate Up
INSERT INTO categories (name) VALUES
('Fiction'),
('Non-Fiction'),
('Science'),
('Biography'),
('Fantasy'),
('Mystery'),
('Romance'),
('Horror'),
('Self-Help'),
('History');

-- +migrate Down
truncate table categories;

