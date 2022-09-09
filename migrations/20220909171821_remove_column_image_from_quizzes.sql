-- +goose Up
alter table quizzes drop column image;

-- +goose Down
alter table quizzes add image string;
