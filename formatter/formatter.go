package formatter

import (
	"fmt"
	"os"
	"strings"
)

func FormatInput(unformattedText string, fileName string) string {

	unformattedText = strings.Replace(unformattedText, "\n", ";\n", -1)
	unformattedText = strings.Replace(unformattedText, "{;", "{", -1)
	unformattedText = strings.Replace(unformattedText, "};", "}", -1)
	unformattedText = strings.Replace(unformattedText, ";;", ";", -1)
	unformattedText = strings.Replace(unformattedText, "\n;\n", "\n\n", -1)

	file, err := os.Create("./src/" + fileName)
	if err != nil {
		fmt.Println("Warning error in auto formatting input file!")
	}
	defer file.Close()

	file.WriteString(unformattedText)
	fmt.Println("Auto formatting successful\n")
	return strings.Replace(unformattedText, "\n", "", -1)
}
