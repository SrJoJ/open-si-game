package questpack

import "fmt"

type Category struct {
	Title     string      `json:"title"`
	Questions []*Question `json:"questions"`
}

func NewCategory(title string) *Category {
	return &Category{
		Title: title,
	}
}

func (cat *Category) AddQuestion(question *Question) {
	cat.Questions = append(cat.Questions, question)
}

// MoveQuestion takes 2 indexes and move question from first to target position
func (cat *Category) MoveQuestion(from, to int) error {
	if from >= len(cat.Questions) || to >= len(cat.Questions) {
		return fmt.Errorf("one of the questions is out of range")
	}
	if from == to {
		return nil
	}

	return nil
}
