package common

type Student interface {
	CitizenId() string
	SetCourse(Course)
}

type Course interface {
	Apply(Student) bool
}
