package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCorrectNameEmpty(t *testing.T) {
	name := correctName("")
	assert.Equal(t, name, "名無しさん")
}

func TestCorrectNameNotEmpty(t *testing.T) {
	name := correctName("Chino")
	assert.Equal(t, name, "Chino")
}

func TestCorrectNameWithUnicode(t *testing.T) {
	name := correctName("香風智乃")
	assert.Equal(t, name, "香風智乃")
}

func TestCorrectNameBeforeLimit(t *testing.T) {
	generatedName := strings.Repeat("a", 19)
	name := correctName(generatedName)
	assert.Equal(t, name, generatedName)
}

func TestCorrectNameJustLimit(t *testing.T) {
	generatedName := strings.Repeat("a", 20)
	name := correctName(generatedName)
	assert.Equal(t, name, generatedName)
}

func TestCorrectNameJustLimitWithUnicode(t *testing.T) {
	generatedName := strings.Repeat("あ", 20)
	name := correctName(generatedName)
	assert.Equal(t, name, generatedName)
}

func TestCorrectNameOverLimit(t *testing.T) {
	generatedName := strings.Repeat("a", 21)
	name := correctName(generatedName)
	assert.Equal(t, name, generatedName[:20])
}

func TestCorrectNameOverLimitWithUnicode(t *testing.T) {
	generatedName := strings.Repeat("あ", 21)
	name := correctName(generatedName)
	assert.Equal(t, name, (string)(([]rune)(generatedName)[:20]))
}

func BenchmarkCorrectName(b *testing.B) {
	name := strings.Repeat("Chino", 20)
	newName := ""
	_ = newName
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newName = correctName(name)
	}
}
