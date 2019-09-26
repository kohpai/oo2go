package main

import (
	cs "github.com/kohpai/oo2go/course"
	st "github.com/kohpai/oo2go/student"
)

func main() {
	student := st.NewStudent("S01")
	ranking := make(cs.Ranking)
	ranking["S01"] = 1
	course := cs.NewCourse("C01", ranking)

	student.SetPreferredCourse(1, course)
}
