package advent20201221_test

import (
	advent "advent20201221"
	"testing"
)

func TestRecordsFromFile(t *testing.T) {
	advent.FoodsFromFile("sample.txt")
	// t.Fail()
}
func TestSafeCount(t *testing.T) {
	foods := advent.FoodsFromFile("input.txt")
	amap := advent.BuildAllergenCountMap(foods)
	potentialAllergens := advent.PotentialAllergens(amap)
	safeCt := advent.SafeCount(foods, potentialAllergens)
	if safeCt != 1945 {
		t.Fail()
	}
}
func TestPrintPossibles(t *testing.T) {
	// I just hand-sudoku'd it here
	foods := advent.FoodsFromFile("input.txt")
	amap := advent.BuildAllergenCountMap(foods)
	advent.PrintPossibleMatches(amap)
	// t.Fail()
}

// func TestPart1(t *testing.T) {
// 	fmt.Println(advent.Part1("sample.txt"))
// 	fmt.Println(advent.Part1("input.txt"))
// 	// t.Fail()
// }

// func TestPart2(t *testing.T) {
// 	fmt.Println(advent.Part2("sample.txt"))
// 	fmt.Println(advent.Part2("input.txt"))
// 	t.Fail()
// }
