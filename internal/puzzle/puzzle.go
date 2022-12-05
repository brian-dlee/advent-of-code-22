package puzzle

import (
	"fmt"
	"os"
	"path"
	"strings"
)

const (
	PART_A Part = "a"
	PART_B = "b"
)

type Part string

const (
	FILE_S1 InputFile = "s1"
	FILE_S2 = "s2"
	FILE_S3 = "s3"
	FILE_IN = "input"
)

type InputFile string

func GetInputFile(day int, part Part, file InputFile) string {
	return path.Join("data", fmt.Sprintf("d%d", day), fmt.Sprintf("%s.%s.txt", part, file))
}

func ReadInputLinesOrPanic(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	strip := 0

	for len(lines)-strip-1 > 0 {
		if strings.TrimSpace(lines[len(lines)-strip-1]) == "" {
			strip += 1
		} else {
			break
		}
	}

	return lines[0:len(lines)-strip]
}
