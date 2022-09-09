-- +goose Up
-- +goose StatementBegin
create table quizzes (
  	id integer primary key autoincrement,
		image string,
		correct_id integer,
		correct_rate float,
		level integer,
		created_at datetime
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table quizzes;
-- +goose StatementEnd
