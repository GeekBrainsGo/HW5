DROP DATABASE IF EXISTS blog;
CREATE DATABASE blog;
USE blog;

DROP TABLE IF EXISTS posts;
CREATE TABLE posts (
                        id SERIAL PRIMARY KEY,
                        title VARCHAR ( 255 ),
                        text TEXT,
                        created_at DATETIME,
                        updated_at DATETIME,
                        deleted_at DATETIME
);

DROP TABLE IF EXISTS labels;
CREATE TABLE labels (
                     id SERIAL PRIMARY KEY,
                     name VARCHAR ( 255 ),
                     created_at DATETIME,
                     updated_at DATETIME,
                     deleted_at DATETIME
);

DROP TABLE IF EXISTS post_label;
CREATE TABLE post_label (
                               id SERIAL PRIMARY KEY,
                               orderId INT UNSIGNED,
                               productId INT UNSIGNED,
                               created_at DATETIME,
                               updated_at DATETIME,
                               deleted_at DATETIME
);

INSERT INTO posts VALUES (DEFAULT, 'Post 1', '1 Post Text', DEFAULT, DEFAULT, DEFAULT);
INSERT INTO posts VALUES (DEFAULT, 'Post 2', '2 Post Text', DEFAULT, DEFAULT, DEFAULT);
INSERT INTO posts VALUES (DEFAULT, 'Post 3', '3 Post Text', DEFAULT, DEFAULT, DEFAULT);

INSERT INTO labels VALUES (DEFAULT, 'Label 1', DEFAULT, DEFAULT, DEFAULT);
INSERT INTO labels VALUES (DEFAULT, 'Label 2', DEFAULT, DEFAULT, DEFAULT);
INSERT INTO labels VALUES (DEFAULT, 'Label 3', DEFAULT, DEFAULT, DEFAULT);
INSERT INTO labels VALUES (DEFAULT, 'Label 4', DEFAULT, DEFAULT, DEFAULT);

INSERT INTO post_label VALUES (DEFAULT, 1, 1, DEFAULT, DEFAULT, DEFAULT);
INSERT INTO post_label VALUES (DEFAULT, 1, 2, DEFAULT, DEFAULT, DEFAULT);
INSERT INTO post_label VALUES (DEFAULT, 2, 3, DEFAULT, DEFAULT, DEFAULT);
INSERT INTO post_label VALUES (DEFAULT, 2, 4, DEFAULT, DEFAULT, DEFAULT);
