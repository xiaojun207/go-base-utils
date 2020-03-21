package utils

import (
	"fmt"
	"testing"
)

func TestExternalIP(t *testing.T) {
	ip, err := ExternalIP()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ip.String())
}
