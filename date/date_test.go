package date

import "fmt"

// import "testing"

func ExampleDate() {
	var d = NewDate(2018, 4, 15)
	fmt.Println(d)

	SepField(SEP_PATH)
	fmt.Println(d)

	fmt.Println(d.IntNum())
	fmt.Println(d.IntNum() + 1)
	// Output: 2018-04-15
	// 2018/04/15
	// 20180415
	// 20180416
}
