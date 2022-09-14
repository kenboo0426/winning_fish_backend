-- +goose Up
create table online_matches (
  id integer primary key autoincrement,
  person_number integer,
  participants_number integer,
  started_at datetime,
  finished_at datetime,
  created_at datetime,
  updated_at datetime
);

-- +goose Down
drop table online_matches;
