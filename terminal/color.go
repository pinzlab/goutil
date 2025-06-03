package terminal

type BGColor string
type FGColor string
type Style string

const (
	BgGreen   BGColor = "\033[97;42m"
	BgWhite   BGColor = "\033[90;47m"
	BgYellow  BGColor = "\033[90;43m"
	BgRed     BGColor = "\033[97;41m"
	BgBlue    BGColor = "\033[97;44m"
	BgMagenta BGColor = "\033[97;45m"
	BgCyan    BGColor = "\033[97;46m"

	FgBlack   FGColor = "\033[30m"
	FgRed     FGColor = "\033[31m"
	FgGreen   FGColor = "\033[32m"
	FgYellow  FGColor = "\033[33m"
	FgBlue    FGColor = "\033[34m"
	FgMagenta FGColor = "\033[35m"
	FgCyan    FGColor = "\033[36m"
	FgWhite   FGColor = "\033[37m"

	Reset Style = "\033[0m"
	Bold  Style = "\033[1m"
)
