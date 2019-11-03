package filtering_test

import (
	"gg.gov.revenue.gonfluence/files/filtering"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatching(t *testing.T) {

	overThreeChars := func(s string) bool { return len(s) > 3 }
	matchingNames := filtering.Matching([]string{"foo", "bar", "bazz", "quxxx"}, overThreeChars)

	assert.Equal(t,  matchingNames, []string{"bazz", "quxxx"})
}

func TestNotMatching(t *testing.T) {

	overThreeChars := func(s string) bool { return len(s) > 3 }
	matchingNames := filtering.NotMatching([]string{"foo", "bar", "bazz", "quxxx"}, overThreeChars)

	assert.Equal(t,  matchingNames, []string{"foo", "bar"})
}

func TestAny(t *testing.T) {

	names := []string{"foo", "bar", "bazz", "quxxx"}
	overThreeChars := func(s string) bool { return len(s) > 3 }
	overSixChars := func(s string) bool { return len(s) > 6 }

	assert.True(t, filtering.Any(names, overThreeChars))
	assert.False(t,  filtering.Any(names, overSixChars))
}