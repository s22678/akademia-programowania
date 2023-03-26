package academy

import "math"

type Student struct {
	Name       string
	Grades     []int
	Project    int
	Attendance []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {
	if len(grades) != 0 {
		var sum int
		for _, v := range grades {
			sum += v
		}
		avg := float64(sum) / float64(len(grades))
		return int(math.Round(avg))
	} else {
		return 0
	}
}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from 0 to 1.
func AttendancePercentage(attendance []bool) float64 {
	if len(attendance) != 0 {
		var sum float64 = 0
		for _, v := range attendance {
			if v == true {
				sum += 1
			}
		}
		return sum / float64(len(attendance))
	} else {
		return 0
	}
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.

// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.
func FinalGrade(s Student) int {
	var average_grade int
	if s.Project == 1 {
		return 1
	} else {
		average_grade = AverageGrade(s.Grades)
		if average_grade == 1 {
			return 1
		}
		var all_grades []int
		all_grades = append(all_grades, average_grade)
		all_grades = append(all_grades, s.Project)
		final_grade := AverageGrade(all_grades)

		switch attendance := AttendancePercentage(s.Attendace); {
		case attendance < 0.6:
			return 1
		case attendance < 0.8:
			final_grade--
			return final_grade
		}

		return final_grade
	}
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {
	m := make(map[string]uint8)
	if len(students) != 0 {
		for _, v := range students {
			m[v.Name] = uint8(FinalGrade(v))
		}
	}
	return m
}
