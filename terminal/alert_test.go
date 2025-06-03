package terminal_test

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/pinzlab/goutil/terminal"
	"github.com/stretchr/testify/assert"
)

// captureOutput temporarily redirects log output and returns what was logged
func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	f()
	return buf.String()
}

func TestAlertFormatting(t *testing.T) {
	result := terminal.Alert(terminal.BgGreen, "Success", "All good")
	assert.Contains(t, result, "Success", "should contain title")
	assert.Contains(t, result, "All good", "should contain message")
	assert.Contains(t, result, terminal.Bold, "should contain bold")
	assert.Contains(t, result, terminal.BgGreen, "should contain background color")
}

func TestSuccessLog(t *testing.T) {
	output := captureOutput(func() {
		terminal.Success("Operation completed")
	})
	assert.Contains(t, output, "Success")
	assert.Contains(t, output, "Operation completed")
}

func TestWarningLog(t *testing.T) {
	output := captureOutput(func() {
		terminal.Warning("Watch out")
	})
	assert.Contains(t, output, "Warning")
	assert.Contains(t, output, "Watch out")
}

func TestInfoLog(t *testing.T) {
	output := captureOutput(func() {
		terminal.Info("Just so you know")
	})
	assert.Contains(t, output, "Info")
	assert.Contains(t, output, "Just so you know")
}

func TestAboutLog(t *testing.T) {
	output := captureOutput(func() {
		terminal.About("About", "First migration")
	})
	assert.Contains(t, output, "About")
	assert.Contains(t, output, "First migration")
}
