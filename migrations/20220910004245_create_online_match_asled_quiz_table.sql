-- +goose Up
create table online_match_asked_quizzes (
  id integer primary key autoincrement,
  quiz_id integer,
  online_match_id integer,
  created_at datetime,
  updated_at datetime
);

-- +goose Down
drop table online_match_asked_quizzes;
