package techpalace
import "strings"
func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    msg := strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
    return msg
}

func CleanupMessage(oldMsg string) string {
    return strings.Trim(oldMsg, "\n\r\t *")
}
