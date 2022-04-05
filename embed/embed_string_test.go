package embed

import (
	_ "embed"
	"fmt"
	"testing"
)

// embed file string

//go:embed sinopsis.txt
var marin string

// kamu bisa membuat embed ini lebih dari 1 variable

func TestEmbedString(t *testing.T) {
	fmt.Println(marin)
}
