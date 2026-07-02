CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    isbn VARCHAR(50) UNIQUE NOT NULL,
    category_id INT REFERENCES categories(id) ON DELETE SET NULL
);

CREATE TABLE loans (
    id SERIAL PRIMARY KEY,
    book_id INT REFERENCES books(id) ON DELETE CASCADE,
    borrower_name VARCHAR(255) NOT NULL,
    loan_date DATE DEFAULT CURRENT_DATE,
    return_date DATE
);

INSERT INTO categories (name) VALUES ('Fiction'), ('Science'), ('History');

INSERT INTO books (title, author, isbn, category_id) VALUES 
('The Hobbit', 'J.R.R. Tolkien', '978-0261102217', 1),
('A Brief History of Time', 'Stephen Hawking', '978-0553380163', 2);

INSERT INTO loans (book_id, borrower_name, loan_date) VALUES 
(1, 'Alice Smith', '2026-06-15');
