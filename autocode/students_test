package coverage

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestPeople_Len(t *testing.T) {
	testCases := []struct {
		name     string
		people   People
		expected int
	}{
		{name: "2 person",
			people: People{
				Person{
					firstName: "Chris",
					lastName:  "Evans",
					birthDay:  time.Date(1981, 06, 13, 10, 10, 10, 0, time.UTC),
				},
				Person{
					firstName: "Johnny",
					lastName:  "Depp",
					birthDay:  time.Date(1963, 06, 9, 10, 10, 10, 0, time.UTC),
				},
			},
			expected: 2,
		},
		{name: "1 person",
			people: People{
				Person{
					firstName: "Chris",
					lastName:  "Evans",
					birthDay:  time.Date(1981, 06, 13, 10, 10, 10, 0, time.UTC),
				},
			},
			expected: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Len(t, tc.people, tc.expected)
			//if result := tc.people.Len(); result != tc.expected {
			//	t.Errorf("Incorrect result. Expect %d, got %d", tc.expected, result)

		})
	}
}
func TestPeople_Less(t *testing.T) {
	testCases := []struct {
		name     string
		people   People
		birthDay time.Time
		expected bool
	}{
		{name: "false",
			people: People{
				Person{
					firstName: "Chris",
					lastName:  "Evans",
					birthDay:  time.Date(1981, 06, 13, 10, 10, 10, 0, time.UTC),
				},
				Person{
					firstName: "Johnny",
					lastName:  "Depp",
					birthDay:  time.Date(1963, 06, 9, 10, 10, 10, 0, time.UTC),
				},
			},
			expected: false,
		},
		{name: "true",
			people: People{
				Person{
					firstName: "Chris",
					lastName:  "Evans",
					birthDay:  time.Date(1981, 06, 13, 10, 10, 10, 0, time.UTC),
				},
				Person{
					firstName: "Chris",
					lastName:  "Brown",
					birthDay:  time.Date(1981, 06, 13, 10, 10, 10, 0, time.UTC),
				},
			},
			expected: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Less(t, tc.people, tc.expected)
		})
	}
}
func TestPeople_Swap(t *testing.T) {

}

func TestMatrix_Cols(t *testing.T) {

}

func TestMatrix_Rows(t *testing.T) {

}
func TestMatrix_Set(t *testing.T) {

}

func TestNew(t *testing.T) {

}
