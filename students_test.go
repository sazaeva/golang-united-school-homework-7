package coverage

import (
	_ "github.com/stretchr/testify/assert"
	"os"
	"reflect"
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
		{name: "3 person",
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
				Person{
					firstName: "Chris",
					lastName:  "Hemsworth",
					birthDay:  time.Date(1983, 8, 11, 10, 10, 10, 0, time.UTC),
				},
			},
			expected: 3,
		},
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
			if result := tc.people.Len(); result != tc.expected {
				t.Errorf("Incorrect result. Expect %d, got %d", tc.expected, result)
			}
		})
	}
}

func TestPeople_Less(t *testing.T) {
	type val struct {
		i, j int
	}

	testCases := []struct {
		name     string
		people   People
		val      val
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
			val:      val{0, 1},
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
			val:      val{0, 1},
			expected: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := tc.people.Less(tc.val.j, tc.val.i); result != tc.expected {
				t.Errorf("Incorrect result. Expect %v, got %v", tc.expected, result)
			}
		})
	}
}
func TestPeople_Swap(t *testing.T) {
	type val struct {
		i, j int
	}

	testCases := []struct {
		name     string
		people   People
		val      val
		expected People
	}{
		{name: "Swap Chris, Johnny",
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
			val: val{0, 1},
			expected: People{
				Person{
					firstName: "Johnny",
					lastName:  "Depp",
					birthDay:  time.Date(1963, 06, 9, 10, 10, 10, 0, time.UTC),
				},
				Person{
					firstName: "Chris",
					lastName:  "Evans",
					birthDay:  time.Date(1981, 06, 13, 10, 10, 10, 0, time.UTC),
				}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.people.Swap(tc.val.i, tc.val.j)
		})
	}
}

func TestMatrix_Rows(t *testing.T) {
	type Table struct {
		rows, cols int
		data       []int
	}

	testCases := []struct {
		name     string
		table    Table
		expected [][]int
	}{
		{name: "Matrix: 4x4",
			table: Table{
				rows: 4,
				cols: 4,
				data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},
			expected: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
		},
		{name: "Matrix: 2x2",
			table: Table{
				rows: 2,
				cols: 2,
				data: []int{1, 2, 3, 4}},
			expected: [][]int{{1, 2}, {3, 4}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := &Matrix{
				rows: tc.table.rows,
				cols: tc.table.cols,
				data: tc.table.data,
			}
			if result := m.Rows(); !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Incorrect result. Expect %v, got %v", tc.expected, result)
			}

		})
	}
}
func TestMatrix_Cols(t *testing.T) {
	type Table struct {
		rows, cols int
		data       []int
	}

	testCases := []struct {
		name     string
		table    Table
		expected [][]int
	}{
		{name: "Matrix: 4x4",
			table: Table{
				rows: 4,
				cols: 4,
				data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},
			expected: [][]int{{1, 5, 9, 13}, {2, 6, 10, 14}, {3, 7, 11, 15}, {4, 8, 12, 16}},
		},
		{name: "Matrix: 2x2",
			table: Table{
				rows: 2,
				cols: 2,
				data: []int{1, 2, 3, 4}},
			expected: [][]int{{1, 3}, {2, 4}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := &Matrix{
				rows: tc.table.rows,
				cols: tc.table.cols,
				data: tc.table.data,
			}
			if result := m.Cols(); !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Incorrect result. Expect %v, got %v", tc.expected, result)
			}

		})
	}
}
func TestMatrix_Set(t *testing.T) {
	matrix := &Matrix{4, 4, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 17}

	result := matrix.Set(3, 3, 17)

	if !result || !reflect.DeepEqual(matrix.data, expected) {
		t.Errorf("Incorrect result. Expect %v, got %v", expected, result)
	}

	result = matrix.Set(-1, 2, 17)
	if result {
		t.Errorf("Incorrect result")
	}
}
func TestNew(t *testing.T) {
	type val struct {
		str string
	}
	testCases := []struct {
		name     string
		val      val
		expected *Matrix
	}{
		{
			name: "Create a matrix from a string",
			val: val{
				str: "1 2 \n3 4"},
			expected: &Matrix{
				rows: 2,
				cols: 2,
				data: []int{1, 2, 3, 4}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result, err := New(tc.val.str); !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Incorrect result. Expect %v, got %v, %v", tc.expected, result, err)
			}
		})
	}
}
