// greet_test.go
// run test from this directory: 
//     go test -v

package greet

import "testing"

func Test_1_Greet_add(t *testing.T) {
    res := Greet_add(1,2)
    if res == int64(3) {
        return
    } else {
        t.Error("1+2 = 3 test failed, got ", res)
        //t.Fail()
    }
}

func Test_2_Greet_add(t *testing.T) {
    res := Greet_add(-1,1)
    if res == int64(0) {
        return
    } else {
        t.Error("-1 + 1 = 0 test failed, got ", res)
        //t.Fail()
    }
}

