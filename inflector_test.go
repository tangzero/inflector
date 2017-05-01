package inflector_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tangzero/inflector"
)

var SingularToPlural = map[string]string{
	"search":      "searches",
	"switch":      "switches",
	"fix":         "fixes",
	"box":         "boxes",
	"process":     "processes",
	"address":     "addresses",
	"case":        "cases",
	"stack":       "stacks",
	"wish":        "wishes",
	"fish":        "fish",
	"jeans":       "jeans",
	"funky jeans": "funky jeans",
	"my money":    "my money",
	"category":    "categories",
	"query":       "queries",
	"ability":     "abilities",
	"agency":      "agencies",
	"movie":       "movies",
	"archive":     "archives",
	"index":       "indices",
	"wife":        "wives",
	"safe":        "saves",
	"half":        "halves",
	"move":        "moves",
	"salesperson": "salespeople",
	"person":      "people",
	"spokesman":   "spokesmen",
	"man":         "men",
	"woman":       "women",
	"basis":       "bases",
	"diagnosis":   "diagnoses",
	"diagnosis_a": "diagnosis_as",
	"datum":       "data",
	"medium":      "media",
	"stadium":     "stadia",
	"analysis":    "analyses",
	"my_analysis": "my_analyses",
	"node_child":  "node_children",
	"child":       "children",
	"experience":  "experiences",
	"day":         "days",
	"comment":     "comments",
	"foobar":      "foobars",
	"newsletter":  "newsletters",
	"old_news":    "old_news",
	"news":        "news",
	"series":      "series",
	"miniseries":  "miniseries",
	"species":     "species",
	"quiz":        "quizzes",
	"perspective": "perspectives",
	"ox":          "oxen",
	"photo":       "photos",
	"buffalo":     "buffaloes",
	"tomato":      "tomatoes",
	"dwarf":       "dwarves",
	"elf":         "elves",
	"information": "information",
	"equipment":   "equipment",
	"bus":         "buses",
	"status":      "statuses",
	"status_code": "status_codes",
	"mouse":       "mice",
	"louse":       "lice",
	"house":       "houses",
	"octopus":     "octopi",
	"virus":       "viri",
	"alias":       "aliases",
	"portfolio":   "portfolios",
	"vertex":      "vertices",
	"matrix":      "matrices",
	"matrix_fu":   "matrix_fus",
	"axis":        "axes",
	"taxi":        "taxis",
	"testis":      "testes",
	"crisis":      "crises",
	"rice":        "rice",
	"shoe":        "shoes",
	"horse":       "horses",
	"prize":       "prizes",
	"edge":        "edges",
	"database":    "databases",
}

var CamelToUnderscore = map[string]string{
	"Product":               "product",
	"SpecialGuest":          "special_guest",
	"ApplicationController": "application_controller",
	"Area51Controller":      "area51_controller",
}

var CamelToDashes = map[string]string{
	"Product":               "product",
	"SpecialGuest":          "special-guest",
	"ApplicationController": "application-controller",
	"Area51Controller":      "area51-controller",
}

var ToTables = map[string]string{
	"Product":      "products",
	"AdminUser":    "admin_users",
	"user-account": "user_accounts",
	"specialOrder": "special_orders",
}
var ToForeignKey = map[string]string{
	"Product":      "product_id",
	"AdminUser":    "admin_user_id",
	"user-account": "user_account_id",
	"specialOrder": "special_order_id",
}

func TestPluralize(t *testing.T) {
	inflector.ClearCache()
	inflector.ShouldCache = false
	for singular, plural := range SingularToPlural {
		assert.Equal(t, plural, inflector.Pluralize(singular))
	}
}

func TestCachedPluralize(t *testing.T) {
	inflector.ClearCache()
	inflector.ShouldCache = true
	for singular, plural := range SingularToPlural {
		assert.Equal(t, plural, inflector.Pluralize(singular))
	}
}

func TestSingularize(t *testing.T) {
	inflector.ClearCache()
	inflector.ShouldCache = false
	for singular, plural := range SingularToPlural {
		assert.Equal(t, singular, inflector.Singularize(plural))
	}
}

func TestCachedSingularize(t *testing.T) {
	inflector.ClearCache()
	inflector.ShouldCache = true
	for singular, plural := range SingularToPlural {
		assert.Equal(t, singular, inflector.Singularize(plural))
	}
}

func TestCamelize(t *testing.T) {
	for camel, underscore := range CamelToUnderscore {
		assert.Equal(t, camel, inflector.Camelize(underscore))
	}
}

func TestUnderscorize(t *testing.T) {
	for camel, underscore := range CamelToUnderscore {
		assert.Equal(t, underscore, inflector.Underscorize(camel))
	}
}

func TestDasherize(t *testing.T) {
	for camel, dash := range CamelToDashes {
		assert.Equal(t, dash, inflector.Dasherize(camel))
	}
}

func TestTableize(t *testing.T) {
	for term, table := range ToTables {
		assert.Equal(t, table, inflector.Tableize(term))
	}
}

func TestForeignKey(t *testing.T) {
	for term, key := range ToForeignKey {
		assert.Equal(t, key, inflector.ForeignKey(term))
	}
}
