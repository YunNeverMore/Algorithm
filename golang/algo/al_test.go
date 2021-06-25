package algo_test

import (
	"fmt"
	"testing"

	"example.com/algo"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	assert := assert.New(t)
	str := "hello"
	assert.Equal("hello", str)
	assert.Equal(algo.Add(3, 5), 8)
}

func TestParseLang(t *testing.T) {
	assert := assert.New(t)
	output := algo.ParseLang("en-US, fr-CA, fr-FR", []string{"fr-FR", "en-US"})
	assert.Equal([]string{"en-US", "fr-FR"}, output)

	output = algo.ParseLang("fr-CA,    fr-FR", []string{"en-US", "fr-FR"})
	assert.Equal([]string{"fr-FR"}, output)

	output = algo.ParseLang("en-US", []string{"en-US", "fr-CA"})
	assert.Equal([]string{"en-US"}, output)

	output = algo.ParseLang("en", []string{"en-US", "fr-CA", "fr-FR"})
	assert.Equal([]string{"en-US"}, output)

	output = algo.ParseLang("fr", []string{"en-US", "fr-CA", "fr-FR"})
	assert.Equal([]string{"fr-CA", "fr-FR"}, output)

	output = algo.ParseLang("fr-FR, fr, fr-HE", []string{"en-US", "fr-CA", "fr-FR"})
	assert.Equal([]string{"fr-FR", "fr-CA"}, output)

	output = algo.ParseLang("en-US, *", []string{"en-US", "fr-CA", "fr-FR"})
	assert.Equal([]string{"en-US", "fr-CA", "fr-FR"}, output)

	output = algo.ParseLang("fr-FR, fr, *", []string{"en-US", "fr-CA", "fr-FR"})
	assert.Equal([]string{"fr-FR", "fr-CA", "en-US"}, output)

	fmt.Println("test finished")

}
