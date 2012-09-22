package mafsa

import (
	"fmt"
	"testing"
)

func TestEmpty(t *testing.T) {
	b := &builder{}
	b.InsertWord("abc")
	b.InsertWord("abcde")
	b.InsertWord("abcd")
	b.InsertWord("cde")

	fmt.Println(b.Contains("ab"))
	fmt.Println(b.Contains("abc"))
	fmt.Println(b.Contains("abe"))
	fmt.Println(b.Contains("abcde"))
	fmt.Println(b.Contains("abd"))
	fmt.Println(b.Contains("abe"))
	fmt.Println(b.Contains("cdef"))
	fmt.Println(b.Contains("cde"))
}
