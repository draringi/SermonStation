CREATE TABLE "users" if not exists (
	uid SERIAL PRIMARY KEY,
	username varchar(128) UNIQUE NOT NULL,
	salt char(8) NOT NULL,
	algo varchar(16) NOT NULL,
	parameters hstore
);

CREATE TABLE "preachers" IF NOT EXISTS (
	pid SERIAL PRIMARY KEY,
	name varchar(128) UNIQUE NOT NULL,
);

CREATE TABLE "recordings" IF NOT EXISTS (
	rid SERIAL PRIMARY KEY,
	date TIMESTAMP NOT NULL,
	title varchar(32) NOT NULL,
	preacher int FORIEGN KEY(preachers) NOT NULL,
	path varchar(128) NOT NULL,
);
