package testing

import (
	"fmt"
	"testing"

	"github.com/gic-vikash/golang-advance/testing/add"
)

// testing file name should be as xxx_test.go
// in testing file name of that function should be func TestXXX(t *testing.T){}

func TestAdd(t *testing.T){
	Got := add.Addition(100,500)
	Want := 600
	if Got!=Want{
		t.Errorf("Got %v But Want %v",Got,Want)
	}
}
// run Test By go test command 
/*
PASS
ok      github.com/gic-vikash/golang-advance/testing/Test       0.251s
*/

// run test in robust mode
// go test -v

/*
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
PASS
ok      github.com/gic-vikash/golang-advance/testing/Test       0.254s
*/
// Table Tests

func TestIntMinTableDriven(t *testing.T) {
    var tests = []struct {
        a, b int
        want int
    }{
        {50, 17,67 },
        {1, 0, 1},
        {14, -2, 12},
        {-1, -1, -2},
        {-1, 0, -1},
    }
// t.Run enables running “subtests”, one for each table entry. These are shown separately when executing go test -v.

    for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            ans := add.Addition(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}


