package main

import "fmt"

func main() {

	/*
			Maps:
		 		- Collections of value types that are accessed via keys.
				- Created via literals or via make() function.
				- Members accessed via [key] syntax
					- myMap["key"] = "value"
				- Check for presence with "value, ok" form of result.
	*/

	{

		// Example #1 (Creation)

		// Map makes key value pairs.
		// First value has to be string, second integer.
		statePopulations := map[string]int{
			"California":   39250017,
			"Texas":        27862596,
			"Florida":      20612439,
			"New York":     19745289,
			"Pennsylvania": 12802503,
			"Illinois":     12801539,
			"Ohio":         11614373,
		}

		fmt.Println(statePopulations)

	}

	{

		// Example #2 (Appending)

		// If you want to populate the map in a loop or later, define an empty map with length of 10.
		statePopulations := make(map[string]int, 10)

		statePopulations = map[string]int{
			"California":   39250017,
			"Texas":        27862596,
			"Florida":      20612439,
			"New York":     19745289,
			"Pennsylvania": 12802503,
			"Illinois":     12801539,
			"Ohio":         11614373,
		}

		// Fetch a single value from the map.
		fmt.Println("Population of Ohio:", statePopulations["Ohio"])

		// Append the map.
		statePopulations["Georgia"] = 10310371
		fmt.Printf("Population of Georgia: %v\n\n", statePopulations["Georgia"])

	}

	{

		// Example #3 (Removal)

		statePopulations := make(map[string]int, 10)
		statePopulations = map[string]int{
			"California":   39250017,
			"Texas":        27862596,
			"Florida":      20612439,
			"New York":     19745289,
			"Pennsylvania": 12802503,
			"Illinois":     12801539,
			"Ohio":         11614373,
		}

		statePopulations["Georgia"] = 10310371

		// Remove "Georgia" keypair from the map.
		delete(statePopulations, "Georgia")
		fmt.Printf("Population of Georgia: %v\n\n", statePopulations["Georgia"]) // Returns 0.

	}

	{

		// Example #4 (Booleans and checks)

		statePopulations := make(map[string]int, 10)
		statePopulations = map[string]int{
			"California":   39250017,
			"Texas":        27862596,
			"Florida":      20612439,
			"New York":     19745289,
			"Pennsylvania": 12802503,
			"Illinois":     12801539,
			"Ohio":         11614373,
		}

		// Check for presence of Texas in the map:
		_, ok := statePopulations["Texas"]
		fmt.Printf("Texas exists? %v\n", ok)

		// Return the value and presence of the key:
		// Utah doesn't exist in map.
		utahPop, ok := statePopulations["Utah"]
		fmt.Printf("Population and existance of Utah: %v, %v\n", utahPop, ok)

		// Ohio exists in map.
		ohioPop, ok := statePopulations["Ohio"]
		fmt.Printf("Population and existance of Ohio: %v, %v\n\n", ohioPop, ok)

	}

	{

		// Example #5 (Map slices)

		statePopulations := make(map[string]int, 10)
		statePopulations = map[string]int{
			"California":   39250017,
			"Texas":        27862596,
			"Florida":      20612439,
			"New York":     19745289,
			"Pennsylvania": 12802503,
			"Illinois":     12801539,
			"Ohio":         11614373,
		}

		// If you slice 'statePopulations' map to map 'sp' they are linked.
		sp := statePopulations
		delete(sp, "Ohio")
		fmt.Printf("%v", (sp))
		fmt.Printf("%v", (statePopulations))

	}

}
