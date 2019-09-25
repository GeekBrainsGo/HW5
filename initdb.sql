DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id int(11) NOT NULL AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    last_name varchar(100) NOT NULL,
    created_at datetime NOT NULL DEFAULT now(),
    updated_at datetime NOT NULL DEFAULT now(),
    deleted_at datetime DEFAULT NULL,
    access varchar(100) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO users 
(name, last_name, access) 
VALUES 
('Andrew', 'Ivanov', 'guest'),
('Ivan', 'Rogov', ''),
('Alexey', 'Petrovski', 'admin'),
('Michail', 'Melnichenko', 'admin'),
('Yaroslav', 'Krasniy', 'moderator');