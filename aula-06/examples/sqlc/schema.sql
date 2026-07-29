CREATE TABLE users_sqlc_demo (
	id UUID PRIMARY KEY DEFAULT uuidv7(),
	name VARCHAR(255),
	email VARCHAR(255),
	password VARCHAR(255)
);
