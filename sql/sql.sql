CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
) ENGINE=INNODB;

CREATE TABLE posts(
    id int AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(50) not NULL,
    content VARCHAR(300) not NULL,
    author_id int NOT NULL,
    likes int DEFAULT 0,
    createdAt timestamp DEFAULT current_timestamp(),
    foreign key (author_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=INNODB;

CREATE TABLE followers(
    user_id INT NOT NULL, 
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    
    follower_id INT NOT NULL,
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,

    PRIMARY KEY(user_id, follower_id)
) ENGINE=INNODB;

INSERT INTO users (name, email, nick, password) 
VALUES 
("User1", "user1@gmail.com", "u1", "$2a$10$590VHQlO3L0Eyywl.QIBQOb3NkL0biKRjL1c0kfdrRedmPwKe7uSq"),
("User2", "user2@gmail.com", "u2", "$2a$10$590VHQlO3L0Eyywl.QIBQOb3NkL0biKRjL1c0kfdrRedmPwKe7uSq"),
("User3", "user3@gmail.com", "u3", "$2a$10$590VHQlO3L0Eyywl.QIBQOb3NkL0biKRjL1c0kfdrRedmPwKe7uSq");

INSERT INTO followers(user_id, follower_id)
VALUES 
(1, 2),
(2, 3),
(1, 3);

INSERT INTO posts(title, content, author_id)
VALUES
("post1", "post1 content by user 1", 1),
("post2", "post2 content by user 2", 2),
("post3", "post3 content by user 3", 3),
("post4", "post1 content by user 1", 1);