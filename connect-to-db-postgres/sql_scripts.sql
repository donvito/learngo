CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL
);

INSERT INTO users (username, email, password)
VALUES
    ('john_doe', 'john.doe@example.com', 'password123'),
    ('jane_smith', 'jane.smith@example.com', 'securepass'),
    ('alex_johnson', 'alex.johnson@example.com', 'pass123word'),
    ('susan_miller', 'susan.miller@example.com', 'millerpass'),
    ('michael_lee', 'michael.lee@example.com', 'p@ssw0rd'),
    ('emily_williams', 'emily.williams@example.com', 'test123'),
    ('david_smith', 'david.smith@example.com', 'abc123'),
    ('linda_jones', 'linda.jones@example.com', 'userpass'),
    ('ryan_adams', 'ryan.adams@example.com', 'newpass'),
    ('sophia_nguyen', 'sophia.nguyen@example.com', 'qwerty');