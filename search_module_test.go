package wf_test

import (
	"testing"

	"github.com/handlename/alfred-metacpan-workflow"
)

func TestSearchModule(t *testing.T) {
	xmlStr := wf.SearchModule("test")

	t.Log(xmlStr)
}
