CREATE DATABASE IF NOT EXISTS todo_db;
USE todo_db;
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL,
    UNIQUE(username)
);
CREATE TABLE lists (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(50) NOT NULL,
    user_id int not null, foreign key (user_id) references users(id)
);
CREATE TABLE tasks (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(50) NOT NULL,
    description VARCHAR(250) NOT NULL,
    tags VARCHAR(250),
    position INT,
    completed boolean not null default 0,
    list_id int not null, foreign key (list_id) references lists(id)
);