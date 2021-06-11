package controller

import (
	"crypto/rand"
	"io"
	"math/big"
	"strings"
)

// Generator is a struct for arithmetic operations.
type Generator struct {
	Length *int
	Numbers,
	Lowercase,
	Uppercase,
	Symbols bool
	ExtraSymbols string
	charset      string
	Reader       io.Reader
}

const (
	// LowerLetters is the list of lowercase letters.
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	// UpperLetters is the list of uppercase letters.
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Numbers is the list of permitted digits.
	Numbers = "0123456789"

	// Symbols is the list of symbols.
	Symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

// NewGenerator is a factory for Generator.
func NewGenerator() *Generator {
	return &Generator{
		Reader: rand.Reader,
	}
}

// Generate generates random password.
func (g Generator) Generate() (string, error) {
	err := g.initCharset()
	if err != nil {
		return "", err
	}

	ret := make([]byte, *g.Length)
	for i := 0; i < *g.Length; i++ {
		num, err := rand.Int(g.Reader, big.NewInt(int64(len(g.charset))))
		if err != nil {
			return "", err
		}
		ret[i] = g.charset[num.Int64()]
	}

	return string(ret), nil
}

func (g *Generator) initCharset() error {
	var sb strings.Builder

	if g.Numbers {
		sb.WriteString(Numbers)
	}

	if g.Lowercase {
		sb.WriteString(LowerLetters)
	}

	if g.Uppercase {
		sb.WriteString(UpperLetters)
	}

	if g.Symbols {
		sb.WriteString(Symbols)
	}

	if len(g.ExtraSymbols) > 0 {
		sb.WriteString(g.ExtraSymbols)
	}

	g.charset = sb.String()

	if len(g.charset) == 0 {
		return CharsetEmptyError{}
	}

	return nil
}
