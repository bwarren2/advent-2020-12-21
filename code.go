package advent20201221

import "fmt"

type Food struct {
	Ingredients []string
	Allergens   []string
}

type AllergenIngredientCountMap map[string]map[string]int

func (f Food) String() (repr string) {
	repr += fmt.Sprintf("<Ingredients: %v ", f.Ingredients)
	repr += fmt.Sprintf("Allergens: %v >", f.Allergens)
	return
}

func BuildAllergenCountMap(foods []Food) AllergenIngredientCountMap {
	allergenMap := make(map[string]map[string]int)
	for _, food := range foods {
		for _, allergen := range food.Allergens {
			if _, ok := allergenMap[allergen]; !ok {
				allergenMap[allergen] = make(map[string]int)
			}
			allergenMap[allergen]["foodCount"]++
			for _, ingredient := range food.Ingredients {
				allergenMap[allergen][ingredient]++
			}
		}
	}
	return allergenMap
}

func PotentialAllergens(amap AllergenIngredientCountMap) (allergens []string) {
	for _, submap := range amap {
		totalFoods := submap["foodCount"]
		for k, ct := range submap {
			if k == "foodCount" {
				continue
			}
			if ct == totalFoods {
				allergens = append(allergens, k)
			}
		}
	}
	return
}

func SafeCount(foods []Food, potentialAllergens []string) (ct int) {
	for _, food := range foods {
		for _, ingredient := range food.Ingredients {
			if !Contains(potentialAllergens, ingredient) {
				ct++
			}
		}
	}
	return
}

func Contains(lst []string, target string) bool {
	// My kingdom for a "in" operator.
	for _, potential := range lst {
		if target == potential {
			return true
		}
	}
	return false
}

func PrintPossibleMatches(amap AllergenIngredientCountMap) {
	for allergen, submap := range amap {
		totalFoods := submap["foodCount"]
		fmt.Printf("%v could be...", allergen)
		fmt.Println()
		for k, ct := range submap {
			if k == "foodCount" {
				continue
			}
			if ct == totalFoods {
				fmt.Println(k)
			}
		}
	}
	return
}

// // Part1 solves part1
// func Part1(filename string) int {

// }

// // Part2 solves part2
// func Part2(filename string) int {

// }
