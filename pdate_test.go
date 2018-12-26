package pgo_test

import (
	"testing"
	"pgo"
	"time"
)

func TestDate(t *testing.T) {
	if time.Now().Format("2006-02-01 15:04:05") != pgo.Date("Y-m-d H:i:s") {
		t.Fatal("Time formats in map doesn't match")
	}

	if time.Now().Format("2006-02-01T15:04:05") != pgo.Date("Y-m-dTH:i:s") {
		t.Fatal("Time formats in map doesn't match")
	}

	if time.Now().Format("Mon, Jan") != pgo.Date("D, M") {
		t.Fatal("Time formats of week days and month in map doesn't match")
	}
}

func TestTime(t *testing.T) {

}

func TestDateTime(t *testing.T) {

}

func TestSpecCases(t *testing.T) {
	if time.Now().Weekday().String() != pgo.Date("l") {
		t.Fatal("Weekday has not been matched")
	}

	//yearDay, _ := strconv.Atoi(pgo.Date("z"))
	//if time.Now().YearDay() != yearDay {
	//	fmt.Println(pgo.Date("z"), time.Now().YearDay(), yearDay)
	//	t.Fatal("Year day has not been matched")
	//}
	//
	//monthDay, _ := strconv.Atoi(pgo.Date("j"))
	//if time.Now().Day() != monthDay {
	//	fmt.Println(pgo.Date("j"), monthDay)
	//	t.Fatal("Year day has not been matched")
	//}
}