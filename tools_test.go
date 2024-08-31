package toolkit

import (
	"testing"
)

func TestTool_Option(t *testing.T) {
	var testTools Tools

	testTools.SetName("testing")
	if len(testTools.GetName()) > 10 {
		t.Error("name String too long, Max is 10 chars")
	}
}
