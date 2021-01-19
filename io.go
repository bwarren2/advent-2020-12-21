package advent20201221

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var ingredientsRegex = regexp.MustCompile(`^[a-z ]*( \()?`)
var allergensRegex = regexp.MustCompile(`contains [a-z ,]*`)

func FoodsFromFile(filename string) (foods []Food) {
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		ingredients := strings.Split(strings.Trim(ingredientsRegex.FindString(line), " "), " ")
		allergenStr := allergensRegex.FindString(line)
		allergenStr = strings.Replace(allergenStr, "contains ", "", -1)
		allergens := strings.Split(allergenStr, ", ")
		foods = append(foods, Food{Ingredients: ingredients, Allergens: allergens})
	}
	return
}
