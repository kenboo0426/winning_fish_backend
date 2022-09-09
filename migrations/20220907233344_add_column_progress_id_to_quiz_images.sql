-- +goose Up
alter table quiz_images add progress_id integer;

-- +goose Down
alter table quiz_images drop column progress_id;
