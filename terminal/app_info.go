package terminal

import "log"

// Prints formatted log messages.
func AppInfo(title string, msg string) {
	log.Printf("%s%s%-12s%s %s", FgGreen, Bold, title+":", Reset, msg)
}
