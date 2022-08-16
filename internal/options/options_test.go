package options

import (
	"github.com/Incarnation-p-lee/cachalot/pkg/assert"
	"os"
	"testing"
)

func TestCreateOptions(t *testing.T) {
	os.Args = append(os.Args, "-decode")
	os.Args = append(os.Args, "0x1111")
	testOptions := CreateOptions()

	assert.IsNotNil(t, testOptions.Options, "options should not be nil")
	assert.IsEqual(t, 1, len(testOptions.Options), "no arguments should have emtpy options")
}
