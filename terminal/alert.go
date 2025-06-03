package terminal

import (
	"fmt"
	"log"
)

// Alert returns a formatted string with a background color, bold title, and message.
// This is used internally by the Alert log functions.
func Alert(bg BGColor, title, msg string) string {
	return fmt.Sprintf("%s%s %-8s%s %s", bg, Bold, title, Reset, msg)
}

// Success logs a message with a green "Success" label using a styled format.
func Success(msg string) {
	log.Println(Alert(BgGreen, "Success", msg))
}

// Warning logs a message with a yellow "Warning" label using a styled format.
func Warning(msg string) {
	log.Println(Alert(BgYellow, "Warning", msg))
}

// Info logs a message with a blue "Info" label using a styled format.
func Info(msg string) {
	log.Println(Alert(BgBlue, "Info", msg))
}

// Info logs a message with a cyan "about" label using a styled format.
func About(title, msg string) {
	if title == "" {
		log.Println(Alert(BgCyan, "About", msg))
	} else {
		log.Println(Alert(BgCyan, title, msg))
	}
}

// Error logs a message with a red "Error" label using a styled format.
func Error(msg string) {
	log.Println(Alert(BgRed, "Error", msg))
}
