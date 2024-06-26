


DROP TABLE IF EXISTS users;


CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    registration_date TIMESTAMP WITHOUT TIME ZONE,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE authentication(
	id SERIAL PRIMARY KEY,
	user_email varchar(255),
	token_auth varchar(255),
	expire TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE cat(
	id SERIAL PRIMARY KEY,
	cat_name varchar(30),
	cat_race varchar(40),
	cat_sex varchar(6),
	cat_age int,
	description varchar(200),
	image varchar(1000)
);