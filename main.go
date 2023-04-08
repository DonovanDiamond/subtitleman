package main

import (
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Include the path to the folder where subtitles should be merged.")
	}
	path := os.Args[1]
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	output := ""

	num := 1
	for _, f := range files {
		data, err := os.ReadFile(path + f.Name())
		if err != nil {
			panic(err)
		}

		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			switch {
			case line == "", strings.Contains(line, "WEBVTT"), strings.Contains(line, "X-TIMESTAMP"):
				continue
			case regexp.MustCompile(`^[0-9]*$`).MatchString(line):
				output += "\n" + fmt.Sprint(num) + "\n"
				num++
			case strings.Contains(line, "-->"):
				output += fmt.Sprint(line) + "\n"
			default:
				output += fmt.Sprint(line) + "\n"
			}
		}
	}

	os.WriteFile(path+"out.srt", []byte(output), fs.ModeAppend)
}
