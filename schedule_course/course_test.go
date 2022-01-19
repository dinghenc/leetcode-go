package schedule_course

import "testing"

type Course [][]int

func TestNewCourse(t *testing.T) {
	courses := [][]int{
		{1, 2},
		{2, 3},
	}

	c := Course(courses)

	t.Log(c)
}