CREATE TABLE if not exists notes (
	id int auto_increment NOT NULL,
	title varchar(200) NOT NULL,
	description text NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	CONSTRAINT notes_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4;

INSERT INTO notes (title, description) VALUES ('Note 1', 'Description 1');
INSERT INTO notes (title, description) VALUES ('Note 2', 'Description 2');
INSERT INTO notes (title, description) VALUES ('Note 3', 'Description 3');
INSERT INTO notes (title, description) VALUES ('Note 4', 'Description 4');
INSERT INTO notes (title, description) VALUES ('Note 5', 'Description 5');

