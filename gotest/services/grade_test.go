package services_test

import (
	"fmt"
	"gotest/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

// to create a unit test 
// create a func "Test"

// to run the unit test with command line
// First Approch
// cd services
// go test -v

// Second Approch
// go test <modulename>/folder
// in this case go test gotest/services -v
// note that module name is in go.mod file

// run unit test check cover case
// go test gotest/services -cover
func TestCheckGrade(t *testing.T) {

	type testCase struct {
		name     string
		score    int
		expected string
	}

	cases := []testCase{
		{name: "a", score: 80, expected: "A"},
		{name: "b", score: 70, expected: "B"},
		{name: "c", score: 60, expected: "C"},
		{name: "d", score: 50, expected: "D"},
		{name: "f", score: 0, expected: "F"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			grade := services.CheckGrade(c.score)

			assert.Equal(t,c.expected,grade)
			// if grade != c.expected {
			// 	t.Errorf("got %v expected %v", grade, c.expected)
			// }
		})
	}
}

func BenchmarkCheckGrade(b *testing.B) {

	for i := 0; i < b.N; i++ {
		services.CheckGrade(80)
	}
	for i := 0; i < b.N; i++ {
		services.CheckGrade(70)
	}
	for i := 0; i < b.N; i++ {
		services.CheckGrade(60)
	}
	for i := 0; i < b.N; i++ {
		services.CheckGrade(50)
	}
	for i := 0; i < b.N; i++ {
		services.CheckGrade(10)
	}

}

func ExampleCheckGrade() {
	grade := services.CheckGrade(80)
	fmt.Println(grade)
	// Output: A
}