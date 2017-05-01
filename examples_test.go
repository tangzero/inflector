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
