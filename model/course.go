package model

import (
	"fmt"
)

type RankCount map[int]int
type Ranking map[string]int

type Course struct {
	id          string
	isFull      bool
	ranking     Ranking
	jointCourse *JointCourse
}

func NewCourse(
	id string,
	jointCourse *JointCourse,
	ranking Ranking,
) *Course {
	isFull := jointCourse.AvailableSpots() == 0
	course := &Course{
		id,
		isFull,
		ranking,
		jointCourse,
	}

	jointCourse.RegisterCourse(course)
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

func (course *Course) JointCourse() *JointCourse {
	return course.jointCourse
}

func (course *Course) Ranking() Ranking {
	return course.ranking
}

func (course *Course) Apply(student *Student) bool {
	if course.jointCourse.Limit() == 0 {
		return false
	}

	rank, ok := course.ranking[student.CitizenId()]
	if !ok {
		return false
	}

	rankedStudent := NewRankedStudent(student, rank, 0)

	if !course.jointCourse.Apply(rankedStudent) {
		return false
	}

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
