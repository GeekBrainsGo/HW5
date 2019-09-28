DROP TABLE IF EXISTS posts;
CREATE TABLE posts (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(100),
  short VARCHAR(255),
  body TEXT
);

INSERT INTO posts (title, short, body) VALUES ('Title1', 'Short Description', 'Full text post');