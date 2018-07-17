package database

// QuestionCreator creates a Question in database
type QuestionCreator interface {
	Create(Question) bool
}

// QuestionReader reads a Question from database
type QuestionReader interface {
	Read(string) *Question
}

// QuestionUpdater updates a Question in database
type QuestionUpdater interface {
	Update(Question) bool
}

// QuestionDeleter deletes a Question from database
type QuestionDeleter interface {
	Delete(string) bool
}
