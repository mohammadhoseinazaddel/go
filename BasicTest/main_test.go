// +build main

package main

import (
	"io/ioutil"
	"testing"
)

type sumResult struct {
	num1, num2 int
	out        int
}

var Res = []sumResult{
	{1, 3, 4},
	//{1, 3, 5},
}

// func TestSumCal(t *testing.T) {
// 	if res := SumCal(2, 2); res != 4 {
// 		t.Error(" sum of 2 and 2 is != 4")
// 	}
// }

func TestSumArray(t *testing.T) {
	for _, item := range Res {
		res := SumCal(item.num1, item.num2)
		if res != item.out {
			t.Error("Sum Error : res != out")
		}
	}
}

func TestFile(t *testing.T) {
	data, err := ioutil.ReadFile("data/toplearmData.data")
	if err != nil {
		t.Error("Could not open File")
	}
	if string(data) != "toplearn.com" {
		t.Error("string Content Do not Match!")
	}
}

// func TestSumCalNumbers(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		if i == 2 || i == 5 || i == 7 {
// 			t.Error(" sum of 2 and 2 is != 4")
// 		}
// 	}
// }
