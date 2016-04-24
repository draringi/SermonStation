CREATE TABLE IF NOT EXISTS "users" (
	uid SERIAL PRIMARY KEY,
	username varchar(128) UNIQUE NOT NULL CHECK (name <> ''),
	salt char(8) NOT NULL,
	algo varchar(16) NOT NULL CHECK (name <> ''),
	parameters hstore
);

CREATE TABLE IF NOT EXISTS "preachers" (
	pid SERIAL PRIMARY KEY,
	name varchar(128) UNIQUE NOT NULL CHECK (name <> '')
);

CREATE TABLE IF NOT EXISTS "recordings" (
	rid SERIAL PRIMARY KEY,
	recorded_at TIMESTAMP NOT NULL,
	title varchar(32) NOT NULL,
	preacher integer REFERENCES preachers NOT NULL,
	path varchar(128) NOT NULL
);
