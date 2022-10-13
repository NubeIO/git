package git

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestService(t *testing.T) {

	parts := strings.Split("rubix-service-0.0.1-eb71da61.amd64.zip", "-")
	for _, p := range parts {
		match, _ := regexp.MatchString(`^(\d+\.)?(\d+\.)?(\*|\d+)$`, p)
		if match {
			fmt.Println(p)
		}
	}

	parts = strings.Split("flow-framework-0.5.0-340c0ad8.armv7.zip", "-")
	for _, p := range parts {
		match, _ := regexp.MatchString(`^(\d+\.)?(\d+\.)?(\*|\d+)$`, p)
		if match {
			fmt.Println(p)
		}

	}

	parts = strings.Split("rubix-bacnet-master-1.0.3-8aba04b4.amd64.zip", "-")
	for _, p := range parts {
		match, _ := regexp.MatchString(`^(\d+\.)?(\d+\.)?(\*|\d+)$`, p)
		if match {
			fmt.Println(p)
		}
	}
}
