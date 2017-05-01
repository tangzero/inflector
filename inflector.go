package inflector

import (
	"regexp"
	"strings"
)

const (
	regexFlags = "(?i)"
	regexEnd   = "$"
)

var (
	uncountables words
	plurals      rules
	singulars    rules
)

var (
	pluralsCache   = cache{}
	singularsCache = cache{}
)

var (
	camelizeRegex   = regexp.MustCompile(`(?:^|[_-])(.)`)
	upperWordsRegex = regexp.MustCompile(`([A-Z]+)([A-Z][a-z])`)
	lowerWordsRegex = regexp.MustCompile(`([a-z\d])([A-Z])`)
)

// ShouldCache set if the inflector should (or not) cache the inflections
var ShouldCache = false

// ClearCache clear the inflection cache. Both for singulars and plurals.
func ClearCache() {
	singularsCache = cache{}
	pluralsCache = cache{}
}

type cache map[string]string

func (c cache) get(key string, r rules) string {
	value, ok := c[key]
	if !ok {
		value = r.convert(key)
		c[key] = value
	}
	return value
}

type rules []rule

func (r rules) convert(text string) string {
	if len(text) > 0 && !uncountables.contains(text) {
		for _, rule := range r {
			if newtext, ok := rule.apply(text); ok {
				return newtext
			}
		}
	}
	return text
}

type words []string

func (w words) contains(value string) bool {
	for _, word := range w {
		if strings.HasSuffix(value, word) {
			return true
		}
	}
	return false
}

type rule struct {
	regex       *regexp.Regexp
	replacement string
	same        bool
}

func (r rule) apply(term string) (string, bool) {
	if r.regex.MatchString(term) {
		return r.regex.ReplaceAllString(term, r.replacement), true
	}
	return term, false
}

func plural(regex, replacement string) {
	r := rule{regex: regexp.MustCompile(regexFlags + regex + regexEnd), replacement: replacement}
	plurals = append(rules{r}, plurals...)
}

func singular(regex, replacement string) {
	r := rule{regex: regexp.MustCompile(regexFlags + regex + regexEnd), replacement: replacement}
	singulars = append(rules{r}, singulars...)
}

func irregular(s, p string) {
	plural(regexp.QuoteMeta(s), p)
	plural(regexp.QuoteMeta(p), p)
	singular(regexp.QuoteMeta(s), s)
	singular(regexp.QuoteMeta(p), s)
}

func uncountable(words ...string) {
	uncountables = append(uncountables, words...)
}

// Pluralize returns the plural form of the word in the string.
func Pluralize(singular string) string {
	if ShouldCache {
		return pluralsCache.get(singular, plurals)
	}
	return plurals.convert(singular)
}

// Singularize returns the singular form of a word in a string.
func Singularize(plural string) string {
	if ShouldCache {
		return singularsCache.get(plural, singulars)
	}
	return singulars.convert(plural)
}

// Camelize converts strings to UpperCamelCase.
func Camelize(term string) string {
	return camelizeRegex.ReplaceAllStringFunc(term, func(match string) string {
		return strings.Title(strings.Replace(strings.Replace(match, "-", "", 1), "_", "", 1))
	})
}

// Underscorize converts strings to underscored, lowercase form.
func Underscorize(term string) string {
	replacement := "${1}_${2}"
	term = upperWordsRegex.ReplaceAllString(term, replacement)
	term = lowerWordsRegex.ReplaceAllString(term, replacement)
	term = strings.Replace(term, "-", "_", -1)
	return strings.ToLower(term)
}

// Dasherize converts strings to dashed, lowercase form.
func Dasherize(term string) string {
	return strings.Replace(Underscorize(term), "_", "-", -1)
}
