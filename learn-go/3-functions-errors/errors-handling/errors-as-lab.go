package main

import (
	"errors"
	"fmt"
	"log"
)

/*
Create a slice of student IDs (as strings).
Implement an Enrol function which accepts a student ID as an argument. Each time the function is called, it should try to enrol the given student in a course. However, after two students have been enrolled, the function should return an EnrolmentError, simulating the scenario where the course is full.
The EnrolmentError should contain the student ID, a fixed course ID, and an error message.
Create an error method over the struct, in a way that returns a helpful error message.
In the main function, loop over the slice of student IDs and attempt to enrol each student by calling the Enrol function.
If the Enrol function returns an error, use errors.As to check if it is of type EnrolmentError. If it is, print the error and stop processing more students.
*/

var studentSlice = []string{"001", "002", "003", "004", "005"}
var classList = []string{"classID1", "classID2"}
var classEnrollments = make(map[string][]string)

func enroll(student, classId string) error {
	if len(classEnrollments[classId]) < 2 {
		classEnrollments[classId] = append(classEnrollments[classId], student)
		fmt.Println("enrolled", student)
		return nil
	} else {
		e := EnrollmentError{student, classId, ErrEnrollment}
		return &e
	}

}

func main() {
	var ee *EnrollmentError

	for _, class := range classList {
		classEnrollments[class] = []string{}

		for _, student := range studentSlice {
			err := enroll(student, class)
			if err != nil {
				if errors.As(err, &ee) {
					log.Println("Encountered our error!\n", ee.Error())
					break
				}
			}
		}
	}
}

// when making an error type, add Error to end of name so you know what its for
// standard practice
type EnrollmentError struct {
	Id    string
	Class string
	Err   error
}

// has to be a pointer here, in order to use errors.As
func (ee *EnrollmentError) Error() string {
	return "Could not enroll student " + ee.Id + " in course " + ee.Class + ": " + ee.Err.Error()
}

var ErrEnrollment = errors.New("class full")
