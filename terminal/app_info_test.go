package terminal_test

import (
	"testing"

	"github.com/pinzlab/goutil/terminal"
	"github.com/stretchr/testify/assert"
)

func TestAppInfo(t *testing.T) {
	title := "Startup"
	msg := "Application is running"

	output := captureOutput(func() {
		terminal.AppInfo(title, msg)
	})

	assert.Contains(t, output, title+":", "Output should include the formatted title")
	assert.Contains(t, output, msg, "Output should include the message")
	assert.Contains(t, output, terminal.Bold, "Output should contain ANSI bold sequence")
	assert.Contains(t, output, terminal.FgGreen, "Output should use green text")
	assert.Contains(t, output, terminal.Reset, "Output should reset formatting")
}
