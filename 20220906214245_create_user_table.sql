-- +goose Up
CREATE TABLE users (
  	id integer primary key autoincrement,
		uuid string not null unique,
		name string,
		emain string,
		password string,
		rating float,
		role ineter,
		createdt datetime
);

-- +goose Down
DROP TABLE users;