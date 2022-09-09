package models

import "time"

type QuizImage struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	QuizID     int       `json:"quiz_id"`
	ProgressID int       `json:"progress_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (q *Quiz) CreateQuizImages() (err error) {
	cmd := `insert into quiz_images (
		name,
		quiz_id,
		progress_id,
		created_at,
		updated_at
	) values (?, ?, ?, ?, ?)`
	for _, quizImage := range q.QuizImages {
		_, err = Db.Exec(cmd, quizImage.Name, q.ID, quizImage.ProgressID, time.Now(), time.Now())
	}

	return err
}
