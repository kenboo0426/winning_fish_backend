-- +goose Up
-- +goose StatementBegin
create table options (
  	id integer primary key autoincrement,
		name string,
		quiz_id integer
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table options;
-- +goose StatementEnd
