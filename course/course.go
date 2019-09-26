package course

import "fmt"

type Student interface {
	CitizenId() string
	SetCourse(*Course)
}

type RankCount map[int]int
type Ranking map[string]int

type Course struct {
	id      string
	isFull  bool
	ranking Ranking
}

func NewCourse(
	id string,
	ranking Ranking,
) *Course {
	course := &Course{
		id,
		false,
		ranking,
	}

	return course
}

func (course *Course) Id() string {
	return course.id
}

func (course *Course) IsFull() bool {
	return course.isFull
}

func (course *Course) SetIsFull(isFull bool) {
	course.isFull = isFull
}

func (course *Course) Ranking() Ranking {
	return course.ranking
}

func (course *Course) Apply(student Student) bool {
	_, ok := course.ranking[student.CitizenId()]
	if !ok {
		return false
	}

	// some business logics here

	student.SetCourse(course)
	return true
}

func (course *Course) String() string {
	return fmt.Sprintf(
		"{id: %v, isFull: %v, ranking: %v}",
		course.id,
		course.isFull,
		course.ranking,
	)
}
