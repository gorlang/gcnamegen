package gcnamegen

import (
	"fmt"
	"testing"
)

func TestDedupe(t *testing.T) {

	ctx := Context{}
	config(&ctx)
	test := "gorraapaaa"
	expected := "gorapa"
	got := dedupe(test, &ctx)

	if expected != got {
		t.Error("Expected ", expected, "but got", got)
	}
}

func TestDedupeType(t *testing.T) {

	ctx := Context{}
	config(&ctx)
	test := "orelaie"
	expected := "orele"
	got := dedupeType(test, &ctx)

	if expected != got {
		t.Error("Expected ", expected, "but got", got)
	}
}

func TestGetRandomName(t *testing.T) {

	ctx := Context{}
	ctx.FilterConsonants = "fhjvqwxz"
	ctx.FilterVowels = "uyoi"
	ctx.Prefix = []string{"go"}
	ctx.Suffix = []string{"apa"}
	ctx.NameCount = 10
	ctx.Pattern = []string{PREFIX, SYLLABLE, SYLLABLE, SUFFIX}
	ctx.Dedup = false

	config(&ctx)
	for i, _ := range []int{1, 2, 3, 4} {
		name := getRandomName(&ctx)
		fmt.Println(i, name)
		l := len(name)
		expected_suf := ctx.Suffix[0]
		got_suf := name[l-3 : l]
		if expected_suf != got_suf {
			t.Error("Expected suffix ", expected_suf, "but got", got_suf)
		}
		expected_pre := ctx.Prefix[0]
		got_pre := name[0:2]
		if expected_pre != got_pre {
			t.Error("Expected prefix ", expected_pre, "but got", got_pre)
		}
		// tests go here, for example check if filters, pattern, prefix, suffix, dedup works
	}
}

func TestIsVowel(t *testing.T) {

	ctx := Context{}
	config(&ctx)
	vows := "aeouyi"
	for _, v := range vows {
		if !isVowel(v, &ctx) {
			t.Error("Expected vowel, but got none.")
		}
	}

}

func TestFilterList(t *testing.T) {

	list := []rune{'q', 'r', 'd', 'f', 'g', 'ä'}
	filter := "qdä"

	filtered := filterList(list, filter)

	res_str := ""
	for _, c := range filtered {
		res_str += string(c)
	}

	expected := "rfg"

	if expected != res_str {
		t.Error("Expected ", expected, " got ", res_str)
	}

}
