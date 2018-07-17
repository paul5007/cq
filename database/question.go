package database

// Question is used as source for questionaire questions
type Question struct {
	ID      string
	Text    string
	Answer  string
	Options []string
	Verse   string
}
