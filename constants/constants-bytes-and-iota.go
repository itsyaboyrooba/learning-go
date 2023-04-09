package main

import "fmt"

const (
	isAdmin = 1 << iota
	isHeadquaters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
)

func main() {

	/*
		Immutable, but can be shadowed.

		Replaced by the compiler at compile time:
			- Value must be calcuable at compile time.

		Named like variables:
			- PascalCase for exported constants.
			- camelCase for internal constants.

		Typed constants work like immutable variables.
			- Can interoperate only with same type.

		Untyped constants work like literals.
			- Can interoperate with similar types.

		Enumerated constants.
			- Special symbol iota allows related constants to be created easily.
			- Iota starts at 0 in each const block and increments by one.
			- Watch out of constant values that match zero values for variables.

	*/

	// Toots on Admin, ei ole HQs, näeb finantsi ja tixub euroopas.
	var Martin byte = isAdmin | canSeeFinancials | canSeeEurope

	// 100101:
	//
	// 1 - canSeeEurope
	// 0 - canSeeAsia
	// 0 - canSeeAfrica
	// 1 - canSeeFinancials
	// 0 - isHeadquaters
	// 1 - isAdmin

	// Mathias on Admin, on orjapidamise HQs, ei näe finantsi ja tixub Aafrikas põllul.
	var Mathias byte = isAdmin | isHeadquaters | canSeeAfrica

	// 001011:
	//
	// 0 - canSeeEurope
	// 0 - canSeeAsia
	// 1 - canSeeAfrica
	// 0 - canSeeFinancials
	// 1 - isHeadquaters
	// 1 - isAdmin

	fmt.Println("")
	fmt.Printf("Martini õigused: %b\n", Martin)
	// 									isAdmin väärtus byte'st 'Martin' võrdub väärtusega isAdmin? truuuueeee.
	fmt.Printf("On Martin Admin? %v\n", isAdmin&Martin == isAdmin)
	fmt.Printf("On Martin HQs? %v\n", isHeadquaters&Martin == isHeadquaters)
	fmt.Printf("Näeb Martin finantse? %v\n", canSeeFinancials&Martin == canSeeFinancials)

	fmt.Println("")
	fmt.Printf("Mathiase õigused: %b\n", Mathias)
	fmt.Printf("On Mathias Admin? %v\n", isAdmin&Mathias == isAdmin)
	fmt.Printf("On Mathias HQs? %v\n", isHeadquaters&Mathias == isHeadquaters)
	fmt.Printf("Näeb Mathias finantse? %v\n", canSeeFinancials&Mathias == canSeeFinancials)

}
