package models

type Question struct {
	Id           int
	TopicId      int
	QuestionText string
	RightAnswer  string
	WrongAnswer1 string
	WrongAnswer2 string
	WrongAnswer3 string
}

func FindQuestion(id int) Question {
	return Question{id, 1,
		"What what in the butt?",
		"Right answer",
		"Wrong answer",
		"Another wrong answer",
		"You dummy"}
}
