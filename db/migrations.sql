--name: users-table
CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT,
    first_name VARCHAR(16) NOT NULL,
    last_name VARCHAR(16) NOT NULL,
    email VARCHAR(64) NOT NULL UNIQUE,
    is_admin BOOLEAN NOT NULL,
PRIMARY KEY (id))

--name: user-auth-table
CREATE TABLE user_auth (
    id INT NOT NULL AUTO_INCREMENT,
    hash VARCHAR(256) NOT NULL,
    user_id INT NOT NULL,
FOREIGN KEY (user_id) REFERENCES users (id),
PRIMARY KEY (id))

--name: articles-table
CREATE TABLE articles (
    id int NOT NULL AUTO_INCREMENT,
    title VARCHAR(256) NOT NULL,
    description VARCHAR(512) NOT NULL,
    creator_id int NOT NULL,
    body TEXT NOT NULL,
    thumbnail_url VARCHAR(256) NOT NULL,
    date_created DATETIME NOT NULL,
    is_authorized BOOLEAN NOT NULL,
FOREIGN KEY (creator_id) REFERENCES users (id),
PRIMARY KEY (id))

--name: saved-articles-table
CREATE TABLE saved_articles (
    id INT NOT NULL AUTO_INCREMENT,
    article_id INT NOT NULL,
    user_id INT NOT NULL,
FOREIGN KEY (article_id) REFERENCES articles (id),
FOREIGN KEY (user_id) REFERENCES users (id),
PRIMARY KEY (id))

--name: events-table
CREATE TABLE events (
    id INT NOT NULL AUTO_INCREMENT,
    event_name VARCHAR(256) NOT NULL,
    event_description TEXT NOT NULL,
    date DATETIME NOT NULL,
    location VARCHAR(128) NOT NULL,
PRIMARY KEY (id))
