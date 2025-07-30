package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var output strings.Builder
	output.WriteString(`package main

import (
	"os"
	"log"
	"strconv"
)

func main(){`)
	output.WriteString(`
	argsRaw := os.Args
	arg := argsRaw[1:]
	number, err := strconv.Atoi(arg[0])
	if err != nil {
		log.Fatalln("Arg must be number")
	}
	if number == 0 {
		println("even")`)

	for i := range 100000 {
		if i%2 == 0 {
			output.WriteString(fmt.Sprintf(`
	}else if number == %d {
		println(%s)`, i+1, `"odd"`))
		} else {
			output.WriteString(fmt.Sprintf(`
	}else if number == %d {
		println(%s)`, i+1, `"even"`))
		}
	}

	output.WriteString("\n\t}\n}")
	writeFile(output.String())
}

func writeFile(content string) {
	err := os.WriteFile("output/program.go", []byte(content), 0755)
	if err != nil {
		log.Fatalln(err)
	}
}
