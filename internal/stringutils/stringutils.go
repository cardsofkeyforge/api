package stringutils

import (
	"regexp"
	"strings"
)

var stopwords = []string{
	"a", "as", "à", "às", "ao", "aos", "e", "é", "o", "os",
	"da", "das", "de", "des", "do", "dos",
	"em", "essa", "essas", "esse", "esses",
	"ou", "outra", "outras", "outro", "outros",
	"me", "para", "por", "se",
	"um", "uma",
	"v.",
}

var abbrvs = []string{
	"10ny",
	"1n",
	"346e",
	"35a",
	"5ub",
	"5c077",
	"[Censurado]",
	"ant1",
	"calv",
	"edai",
	"egs",
	"eslr",
	"fixed",
	"l30",
	"m3r50",
	"oic",
	"p4rd0",
	"pem",
	"t3r",
}

var specials = map[string]string{
	"ieq": "IeQ",
}

func EspecialTitle(value string) string {
	replaceHead := func(word string) string {
		if Contains(abbrvs, word) {
			return strings.ToUpper(word)
		}
		return strings.Title(word)
	}

	replaceTail := func(word string) string {
		if Contains(stopwords, word) {
			return word
		} else if Contains(abbrvs, word) {
			return strings.ToUpper(word)
		} else {
			special := specials[word]
			if special != "" {
				return special
			}
		}
		return strings.Title(word)
	}

	r := regexp.MustCompile(`[^\s,“”!-]+`)
	newValue := r.ReplaceAllStringFunc(strings.ToLower(value), replaceTail)
	r = regexp.MustCompile(`^[^\s,“”!-]+`)
	return r.ReplaceAllStringFunc(newValue, replaceHead)
}

func Contains(list []string, value string) bool {
	if len(list) == 0 {
		return false
	}

	for _, v := range list {
		if v == value {
			return true
		}
	}

	return false
}
