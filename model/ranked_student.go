package model

import (
	"fmt"
)

type RankedStudent struct {
	student *Student
	rank    int
	index   int
}

func NewRankedStudent(student *Student, rank, index int) *RankedStudent {
	return &RankedStudent{
		student,
		rank,
		index,
	}
}

func (rs *RankedStudent) SetIndex(index int) {
	rs.index = index
}

func (rs *RankedStudent) Student() *Student {
	return rs.student
}

func (rs *RankedStudent) Rank() int {
	return rs.rank
}

func (rs *RankedStudent) String() string {
	return fmt.Sprintf("{student: %v, rank: %v}", rs.student, rs.rank)
}
