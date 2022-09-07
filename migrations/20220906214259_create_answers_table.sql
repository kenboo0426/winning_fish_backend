-- +goose Up
-- +goose StatementBegin
create table answers (
  	id integer primary key autoincrement,
		user_id integer,
		quiz_id integer,
		correct boolean,
		answered_option_id integer
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table answers;
-- +goose StatementEnd
