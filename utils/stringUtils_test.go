package utils

import (
	"log"
	"testing"
)

func TestSubBetween(t *testing.T) {
	s := "<a>this is a test</a>"
	tmp := SubstrBetween(s, "<a>", "</a>")
	log.Println(tmp)

	tmp = SubstrBefore(s, "<a>")
	log.Println(tmp)

	tmp = SubstrAfter(s, "</a>")
	log.Println(tmp)
}
