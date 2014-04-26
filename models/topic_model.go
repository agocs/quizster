package models

type Topic struct {
	Id   int
	Name string
}

func FindTopic(id int) Topic {
	return Topic{1, "Potpourri"}
}
