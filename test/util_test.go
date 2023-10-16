package test

import (
	"regexp"
	"strings"
	"testing"
)

func TestTrimRight(t *testing.T) {
	data := []string{"/", "/user/add/", "//", "/user////add///"}
	reg := regexp.MustCompile(`/+`)
	for _, v := range data {
		base := strings.TrimRight(reg.ReplaceAllString(v, "/"), "/")
		t.Logf("base: >>>%s<<<", base)
	}
}
