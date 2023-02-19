CREATE TABLE users(
	ID VARCHAR(50) primary key,
	username VARCHAR(25) NOT NULL,
	email VARCHAR(50) NOT NULL,
	password VARCHAR(255) NOT NULL,
	address VARCHAR(255) NOT NULL,
	age INT NOT NULL
);

INSERT INTO users VALUES(
	'asd',
	'asokdoask',
	'oaskdok@gmail.com',
	'aowkdoawkd',
	'asodkoasdk',
	4
);