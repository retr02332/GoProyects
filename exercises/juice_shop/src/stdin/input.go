package stdin

import (
	"bufio"
	"fmt"
	"os"
)

// Input ...
func Input(prompt string) string {
	fmt.Printf("%s", prompt)
	captureInput := bufio.NewScanner(os.Stdin)
	captureInput.Scan()
	return captureInput.Text()
}
