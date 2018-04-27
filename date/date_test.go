package date

import "fmt"
import "testing"

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

func TestDate(t *testing.T) {
	var d = Today()
	fmt.Printf("Today: %T\t%v\n", d, d)

	var yearMapDays = make(map[int]int)
	yearMapDays[2018] = 28
	yearMapDays[2016] = 29
	yearMapDays[1900] = 28
	yearMapDays[2000] = 29

	for year, daysExpect := range yearMapDays {
		leap := IsLeap(year)
		daysGot, _ := EndDay(year, Feb)
		fmt.Printf("Feb %d year has %d days (leap:%t).\n", year, daysGot, leap)
		if daysGot != daysExpect {
			t.Errorf("Feb %d year has %d days, while got %d\n",
				year, daysExpect, daysGot)
		}
	}
}

func TestMain(m *testing.M) {
	TestDate(nil)
	ExampleDate()
}
