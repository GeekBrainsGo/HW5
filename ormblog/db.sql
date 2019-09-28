CREATE DATABASE IF NOT EXISTS ormblog;
USE ormblog;
DROP TABLE IF EXISTS posts;
CREATE TABLE posts (
	id SERIAL PRIMARY KEY,
	title VARCHAR(16) NOT NULL,
	author VARCHAR(16) NOT NULL,
	content TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	deleted_at DATETIME DEFAULT NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
COLLATE=utf8_general_ci;

INSERT INTO 
posts (title, author, content)
VALUES
("Post One", "Vasia Pupkine", "This is my very first post"),
("Post Two", "Джон Сноу", "Дедлайн завтра ОГОНЬ!");

