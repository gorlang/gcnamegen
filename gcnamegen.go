package gcnamegen

import (
	"math/rand"
	"time"
)

/*
	Constants used in pattern definitions.
*/

const VOWEL string = "V"
const CONSONANT string = "C"
const SYLLABLE string = "S"
const PREFIX string = "PRE"
const SUFFIX string = "SUF"

/*
	Context holds the config and settings.

	Pattern:			The sequence of prefix, syllables, single chars and suffix.
	Prefix:				Array of strings to use as prefix.
	Suffix:				Array of strings to use as suffix.
	FilterConsonants:	A string containing consonants to exclude in final name.
	FilterVowels:		A string containing vowels to exclude in final name.
	Dedup:				True if double letters as aa, ee, uu should be removed.
	NameCount:			Number of names to generate.
	Vowels:				Filtered vowels to use.
	Consonants:			Filtered consonants to use.
	Syllables:			Generated syllables to use.
*/

type Context struct {
	Pattern          []string
	FilterConsonants string
	FilterVowels     string
	Prefix           []string
	Suffix           []string
	Dedup            bool
	DedupType        bool
	NameCount        int
	Consonants       []rune
	Vowels           []rune
	Syllables        []string
}

/*
	Generates random company or product names with some predefined rules that can be specified.
*/

func GenerateNames(ctx *Context) []string {

	config(ctx)

	if len(ctx.Syllables) <= 0 {
		empty := []string{"Vowels or consonants must be present.", "Please, edit Settings under 'Exclude'.", "", "", "", "", "", "", "", ""}
		return empty
	}

	result := []string{}
	for i := 1; i <= ctx.NameCount; i++ {
		result = append(result, getRandomName(ctx))
	}
	return result
}

/*
	Apply filter on an array of runes.
*/

func filterList(list []rune, filter string) []rune {

	res := []rune{}
	for _, l := range list {
		l_match := false
		for _, f := range filter {
			if rune(f) == l {
				l_match = true
				break
			}
		}
		if !l_match {
			res = append(res, l)
		}
	}
	return res
}

/*
	Setup the configuration. Filter and exclude letters not to use etc.
	Generate syllables.
*/

func config(ctx *Context) {

	ctx.Vowels = []rune{'a', 'e', 'i', 'o', 'u', 'y'}
	ctx.Consonants = []rune{'a', 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'z'}

	if ctx.FilterConsonants != "" {
		c_filtered := filterList(ctx.Consonants, ctx.FilterConsonants)
		ctx.Consonants = c_filtered
	}

	if ctx.FilterVowels != "" {
		v_filtered := filterList(ctx.Vowels, ctx.FilterVowels)
		ctx.Vowels = v_filtered
	}

	generateSyllables(ctx)
}

/*
	Generate all combinations of vowel + consonant or vice versa.
*/

func generateSyllables(ctx *Context) {

	for _, c := range ctx.Consonants {
		for _, v := range ctx.Vowels {
			syl1 := string(c) + string(v)
			syl2 := string(v) + string(c)
			ctx.Syllables = append(ctx.Syllables, syl1)
			ctx.Syllables = append(ctx.Syllables, syl2)
		}
	}
}

/*
	Not used at the moment
	Use for wash of the names generated.
*/

func isVowel(c rune, ctx *Context) bool {

	for _, v := range ctx.Vowels {
		if c == v {
			return true
		}
	}
	return false
}

func isDupType(c1 rune, c2 rune, ctx *Context) bool {

	if isVowel(c1, ctx) {
		if isVowel(c2, ctx) {
			return true
		}
	} else {
		if !isVowel(c2, ctx) {
			return true
		}
	}
	return false
}

func getRandomName(ctx *Context) string {

	name := ""
	for _, part_type := range ctx.Pattern {
		name += string(getRandomPart(part_type, ctx))

	}

	// remove duplicate sequences like ee, aa etc.
	if ctx.Dedup {
		name = dedupe(name, ctx)
	}

	if ctx.DedupType {
		name = dedupeType(name, ctx)
	}

	return name
}

/*
	Remove duplicate letters in name like aa,bb,cc etc.
*/

func dedupe(name string, ctx *Context) string {

	deduped := ""
	i := 0
	for j := 0; j < len(name); j++ {
		if i < len(name)-1 { // next last letter
			if name[i] != name[i+1] {
				// dedupe letters
				deduped += string(name[i])
				i++
			} else {
				// duplicate chars in sequence move on
				i++
			}
		} else {
			// add last char
			deduped += string(name[i])
		}
	}
	return deduped
}

/*
	Remove duplicates in type (vowel or cons) like ae,bc,cd etc.
*/

func dedupeType(name string, ctx *Context) string {

	deduped := ""
	i := 0
	for j := 0; j < len(name); j++ {
		if i < len(name)-1 { // next last letter
			if !isDupType(rune(name[i]), rune(name[i+1]), ctx) {
				// dedupe types
				deduped += string(name[i])
				i++
			} else {
				// duplicate types in sequence move on
				i++
			}
		} else {
			// add last char
			deduped += string(name[i])
		}
	}
	return deduped
}

func getRandomPart(part_type string, ctx *Context) string {

	s := rand.NewSource(int64(time.Now().Nanosecond()))
	r := rand.New(s)

	switch string(part_type) {
	case VOWEL:
		ix := r.Intn(len(ctx.Vowels))
		return string(ctx.Vowels[ix])
	case CONSONANT:
		ix := r.Intn(len(ctx.Consonants))
		return string(ctx.Consonants[ix])
	case PREFIX:
		ix := r.Intn(len(ctx.Prefix))
		return string(ctx.Prefix[ix])
	case SUFFIX:
		ix := r.Intn(len(ctx.Suffix))
		return string(ctx.Suffix[ix])
	default:
		ix := r.Intn(len(ctx.Syllables))
		return string(ctx.Syllables[ix])
	}
}
