-- CREATE TABLE
CREATE TABLE users(id INT AUTO_INCREMENT NOT NULL PRIMARY KEY, name VARCHAR(50), token VARCHAR(100));

-- INITIAL DATA
INSERT INTO users (name, token) VALUES ("John", "Sm9obgo="), ("Rose", "Um9zZQo=");
