package database

// CreateQuestion uses a QuestionCreator to add a Question to database
func CreateQuestion(qc QuestionCreator, q Question) bool {
	return qc.Create(q)
}

// ReadQuestion uses a QuestionReader to read a Question from database
func ReadQuestion(qc QuestionReader, id string) *Question {
	return qc.Read(id)
}

// UpdateQuestion uses a QuestionUpdater to update a Question in database
func UpdateQuestion(qc QuestionUpdater, q Question) bool {
	return qc.Update(q)
}

// DeleteQuestion uses a QuestionDeleter to delete a Question from database
func DeleteQuestion(qc QuestionDeleter, id string) bool {
	return qc.Delete(id)
}
