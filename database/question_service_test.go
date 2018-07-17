package database

import (
	"testing"
)

type MockDB struct {
}

const (
	mockQuestionID     = "BOB123"
	mockQuestionText   = "Why do we do?"
	mockQuestionAnswer = "Cause we do"
	mockQuestionVerse  = "It happened at some point in time"

	mockUpdatedAnswer = "Cause we can and will"
)

var (
	mockQuestionOption = []string{"Cause we can't", "Cause we are", "Cause we aien't"}
	mockQuestion       = &Question{ID: mockQuestionID, Text: mockQuestionText, Answer: mockQuestionAnswer, Options: mockQuestionOption, Verse: mockQuestionVerse}
)

func (m *MockDB) Create(q Question) bool {
	// return true for valid Question
	if q.ID == mockQuestionID {
		return true
	}
	return false
}

func (m *MockDB) Read(id string) *Question {
	// return a valid Question for valid read
	if id == mockQuestionID {
		return &Question{ID: mockQuestionID, Text: mockQuestionText, Answer: mockQuestionAnswer, Options: mockQuestionOption, Verse: mockQuestionVerse}
	}
	return nil
}

func (m *MockDB) Update(q Question) bool {
	// return an updated Question for valid request
	if q.ID == mockQuestionID {
		return true
	}
	return false
}

func (m *MockDB) Delete(id string) bool {
	// return true when for valid id
	if id == mockQuestionID {
		return true
	}
	return false
}

func TestCreateQuestion(t *testing.T) {
	mockDB := &MockDB{}

	q := *mockQuestion

	if !CreateQuestion(mockDB, q) {
		t.Fail()
	}
}

func TestReadQuestion(t *testing.T) {
	mockDB := &MockDB{}

	q := ReadQuestion(mockDB, mockQuestionID)

	if q == nil {
		t.FailNow()
	}
	if q.Answer != mockQuestionAnswer {
		t.Fail()
	}
	if q.Text != mockQuestionText {
		t.Fail()
	}
}

func TestUpdateQuestion(t *testing.T) {
	mockDB := &MockDB{}

	q := *mockQuestion

	if !UpdateQuestion(mockDB, q) {
		t.Fail()
	}
}

func TestDeleteQuestion(t *testing.T) {
	mockDB := &MockDB{}

	q := *mockQuestion

	if !DeleteQuestion(mockDB, q.ID) {
		t.Fail()
	}
}
