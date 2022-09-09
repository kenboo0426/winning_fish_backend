-- +goose Up
create table online_match_joined_users (
  id integer primary key autoincrement,
  user_id integer,
  rank integer,
  remained_time float,
  miss_answered_count integer,
  score integer
);

-- +goose Down
drop table online_match_joined_users;
