-- +goose Up
-- +goose StatementBegin
create table quiz_images (
  id integer primary key autoincrement,
  name string,
  quiz_id integer,
  created_at datetime,
  updated_at datetime
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table quiz_images;
-- +goose StatementEnd