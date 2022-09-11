-- +goose Up
-- +goose StatementBegin
alter table users rename column emain to email;
alter table users drop column password;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table users rename column email to emain;
alter table users add password string;
-- +goose StatementEnd
