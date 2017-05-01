package inflector_test

import (
	"fmt"

	"github.com/tangzero/inflector"
)

func ExamplePluralize() {
	fmt.Println(inflector.Pluralize("post"))
	fmt.Println(inflector.Pluralize("octopus"))
	fmt.Println(inflector.Pluralize("sheep"))
	fmt.Println(inflector.Pluralize("words"))
	fmt.Println(inflector.Pluralize("CamelOctopus"))
	// Output:
	// posts
	// octopi
	// sheep
	// words
	// CamelOctopi
}

func ExampleSingularize() {
	fmt.Println(inflector.Singularize("posts"))
	fmt.Println(inflector.Singularize("octopi"))
	fmt.Println(inflector.Singularize("sheep"))
	fmt.Println(inflector.Singularize("word"))
	fmt.Println(inflector.Singularize("CamelOctopi"))
	// Output:
	// post
	// octopus
	// sheep
	// word
	// CamelOctopus
}

func ExampleCamelize() {
	fmt.Println(inflector.Camelize("my_account"))
	fmt.Println(inflector.Camelize("user-profile"))
	// Output:
	// MyAccount
	// UserProfile
}

func ExampleUnderscorize() {
	fmt.Println(inflector.Underscorize("MyAccount"))
	fmt.Println(inflector.Underscorize("user-profile"))
	// Output:
	// my_account
	// user_profile
}

func ExampleDasherize() {
	fmt.Println(inflector.Dasherize("MyAccount"))
	fmt.Println(inflector.Dasherize("user_profile"))
	// Output:
	// my-account
	// user-profile
}

func ExampleTableize() {
	fmt.Println(inflector.Tableize("RawScaledScorer"))
	fmt.Println(inflector.Tableize("ham_and_egg"))
	fmt.Println(inflector.Tableize("fancyCategory"))
	// Output:
	// raw_scaled_scorers
	// ham_and_eggs
	// fancy_categories
}

func ExampleForeignKey() {
	fmt.Println(inflector.ForeignKey("Message"))
	fmt.Println(inflector.ForeignKey("AdminPost"))
	// Output:
	// message_id
	// admin_post_id
}

func ExampleOrdinal() {
	fmt.Println(inflector.Ordinal(1))
	fmt.Println(inflector.Ordinal(2))
	fmt.Println(inflector.Ordinal(14))
	fmt.Println(inflector.Ordinal(1002))
	fmt.Println(inflector.Ordinal(1003))
	fmt.Println(inflector.Ordinal(-11))
	fmt.Println(inflector.Ordinal(-1021))
	// Output:
	// st
	// nd
	// th
	// nd
	// rd
	// th
	// st
}

func ExampleOrdinalize() {
	fmt.Println(inflector.Ordinalize(1))
	fmt.Println(inflector.Ordinalize(2))
	fmt.Println(inflector.Ordinalize(14))
	fmt.Println(inflector.Ordinalize(1002))
	fmt.Println(inflector.Ordinalize(1003))
	fmt.Println(inflector.Ordinalize(-11))
	fmt.Println(inflector.Ordinalize(-1021))
	// Output:
	// 1st
	// 2nd
	// 14th
	// 1002nd
	// 1003rd
	// -11th
	// -1021st
}
