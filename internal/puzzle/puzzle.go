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
	if content, err := os.ReadFile(path); err != nil {
		panic(err)
	} else {
		return strings.Split(strings.TrimSpace(string(content)), "\n")
	}
}
