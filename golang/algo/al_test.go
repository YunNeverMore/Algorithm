package algo_test

import (
	"testing"

	"example.com/algo"

	"github.com/stretchr/testify/assert"
)

func TestParse1(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(algo.Add(3, 5), 8)
	res, _ := algo.ParseLang("fr-CA, fr-FR", []string{"en-US", "fr-FR"})
	assert.Equal([]string{"fr-FR"}, res)

	res, _ = algo.ParseLang("en-US", []string{"en-US", "fr-CA"})
	assert.Equal([]string{"en-US"}, res)
}

func TestParse2(t *testing.T) {
	assert := assert.New(t)

	res, _ := algo.ParseLang2("en", []string{"en-US", "fr-CA", "fr-FR"})
	assert.Equal([]string{"en-US"}, res)

	res, _ = algo.ParseLang2("fr", []string{"en-US", "fr-CA", "fr-FR"})
	assert.Equal([]string{"fr-CA", "fr-FR"}, res)

	res, _ = algo.ParseLang2("fr-FR, fr", []string{"en-US", "fr-CA", "fr-FR"})
	assert.Equal([]string{"fr-FR", "fr-CA"}, res)
}

func TestParse3(t *testing.T) {
	assert := assert.New(t)

	res, _ := algo.ParseLang3("en-US, *", []string{"en-US", "fr-CA", "fr-FR"})
	assert.Equal([]string{"en-US", "fr-CA", "fr-FR"}, res)

	res, _ = algo.ParseLang3("fr-FR, fr, *", []string{"en-US", "fr-CA", "fr-FR"})
	assert.Equal([]string{"fr-FR", "fr-CA", "en-US"}, res)
}

func TestParse4(t *testing.T) {
	assert := assert.New(t)

	res, _ := algo.ParseLang4("fr-FR;q=1, fr-CA;q=0, fr;q=0.5", []string{"fr-FR", "fr-CA", "fr-BG"})
	assert.Equal([]string{"fr-FR", "fr-BG", "fr-CA"}, res)

	res, _ = algo.ParseLang4("fr-FR;q=1, fr-CA;q=0, *;q=0.5", []string{"fr-FR", "fr-CA", "fr-BG", "en-US"})
	assert.Equal([]string{"fr-FR", "fr-BG", "en-US", "fr-CA"}, res)

	res, _ = algo.ParseLang4("fr-FR;q=1, fr-CA;q=0.8, *;q=0.5", []string{"fr-FR", "fr-CA", "fr-BG", "en-US"})
	assert.Equal([]string{"fr-FR", "fr-CA", "fr-BG", "en-US"}, res)
}
