package utils

import (
	"log"
	"testing"
)

func TestSubBetween(t *testing.T) {
	s := "<a>this is a test</a>"
	tmp := SubStrBetween(s, "<a>", "</a>")
	log.Println(tmp)

	tmp = SubStrStart(s, "<a>")
	log.Println(tmp)

	tmp = SubStrEnd(s, "</a>")
	log.Println(tmp)
}
