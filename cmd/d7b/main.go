package main

import (
	"advent-of-code-22/internal/puzzle"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const puzzleDay = 7 // DAY
const puzzleInputPart = puzzle.PART_A
const puzzleInputFile = puzzle.FILE_IN
const verbose = true

const totalDiskSpace = 70_000_000
const requiredFreeDiskSpace = 30_000_000

var changeDirectoryRegexp = regexp.MustCompile(`^\$\s+cd\s+(\S+)$`)
var commandRegexp = regexp.MustCompile(`^\$ (\w+)\b`)
var fileRegexp = regexp.MustCompile(`^(\d+) (\S+)$`)
var dirRegexp = regexp.MustCompile(`^dir (\S+)$`)

type Directory struct {
	Name       string
	ChildNodes []int
	ParentNode int
	TotalSize  int
}

func NewDirectoryFromString(description string) (*Directory, error) {
	match := dirRegexp.FindAllStringSubmatch(description, -1)

	if len(match) < 1 {
		return nil, fmt.Errorf("invalid directory description: %s", description)
	}

	return &Directory{Name: match[0][1], ChildNodes: make([]int, 0), ParentNode: -1}, nil
}

func (d *Directory) GetSize() int {
	return 0
}

type File struct {
	Name       string
	Size       int
	ParentNode int
}

func NewFileFromString(description string) (*File, error) {
	match := fileRegexp.FindAllStringSubmatch(description, -1)

	if len(match) < 1 {
		return nil, fmt.Errorf("invalid file description: %s", description)
	}

	size, err := strconv.Atoi(match[0][1])
	if err != nil {
		return nil, fmt.Errorf("invalid file size in description: %s - %s", description, err)
	}

	return &File{Size: size, Name: match[0][2], ParentNode: -1}, nil
}

type DirectoryTree struct {
	current     []string
	nodes       map[string]int
	directories map[int]*Directory
	files       map[int]*File
	sequence    int
}

func NewDirectoryTree() *DirectoryTree {
	return &DirectoryTree{
		current: []string{},
		nodes:   map[string]int{"": 0},
		directories: map[int]*Directory{0: {
			Name:       "",
			ChildNodes: make([]int, 0),
			ParentNode: -1,
		}},
		files:    make(map[int]*File, 0),
		sequence: 1,
	}
}

func (t *DirectoryTree) AddDirectoryAtCurrentPath(d *Directory) {
	d.ParentNode = t.nodes[strings.Join(t.current, "/")]
	t.directories[d.ParentNode].ChildNodes = append(t.directories[d.ParentNode].ChildNodes, t.sequence)
	t.directories[t.sequence] = d
	t.nodes[strings.Join(append(t.current, d.Name), "/")] = t.sequence
	t.sequence++
}

func (t *DirectoryTree) AddFileAtCurrentPath(f *File) {
	f.ParentNode = t.nodes[strings.Join(t.current, "/")]

	parent := t.directories[f.ParentNode]
	parent.ChildNodes = append(t.directories[f.ParentNode].ChildNodes, t.sequence)

	for parent.ParentNode >= 0 {
		parent.TotalSize += f.Size
		parent = t.directories[parent.ParentNode]
	}

	parent.TotalSize += f.Size

	t.files[t.sequence] = f
	t.nodes[strings.Join(append(t.current, f.Name), "/")] = t.sequence
	t.sequence++
}

func (t *DirectoryTree) ChangeDirectory(path string) {
	if path == "/" {
		t.current = []string{}
	} else if path == ".." {
		t.current = t.current[0 : len(t.current)-1]
	} else {
		t.current = append(t.current, path)
	}
}

func (t *DirectoryTree) GetDirectories() []*Directory {
	cursor := 0
	search := 0
	directories := make([]*Directory, len(t.directories))

	for search < t.sequence {
		if d, ok := t.directories[search]; ok {
			directories[cursor] = d
			cursor++
		}
		search++
	}

	return directories
}

func (t *DirectoryTree) GetDirectoryFromPath(p string) *Directory {
	if nodeId, ok := t.nodes[strings.TrimLeft(p, "/")]; !ok {
		return nil
	} else if directory, ok := t.directories[nodeId]; !ok {
		return nil
	} else {
		return directory
	}
}

func (t *DirectoryTree) Dump() {
	t.dumpFromPath([]string{})
}

func (t *DirectoryTree) GetFullDirectoryPath(directory *Directory) string {
	nodeId := -1

	for id, d := range t.directories {
		if d == directory {
			nodeId = id
			break
		}
	}

	if nodeId < 0 {
		panic("failed to find node")
	}

	next := t.directories[nodeId]
	var path []string

	for next.ParentNode >= 0 {
		path = append([]string{next.Name}, path...)
		next = t.directories[next.ParentNode]
	}

	return "/" + strings.Join(path, "/")
}

func (t *DirectoryTree) dumpFromPath(p []string) {
	var name string

	key := strings.Join(p, "/")

	nodeId, exists := t.nodes[key]

	if !exists {
		panic(fmt.Errorf("failed to traverse directory: /%s is not a known path", key))
	}

	directory, isDirectory := t.directories[nodeId]

	if !isDirectory {
		panic(fmt.Errorf("failed to traverse directory: /%s (%d) is not a directory", key, nodeId))
	}

	nodes := directory.ChildNodes
	depth := len(p)

	if len(p) == 0 {
		name = "/"
	} else {
		name = p[len(p)-1]
	}

	fmt.Printf("%s- %s (dir)\n", strings.Repeat(" ", depth*2), name)

	for _, node := range nodes {
		if d, ok := t.directories[node]; ok {
			t.dumpFromPath(append(p, d.Name))
		}
	}

	for _, node := range nodes {
		if f, ok := t.files[node]; ok {
			fmt.Printf("%s- %s (file, size=%d)\n", strings.Repeat(" ", (depth+1)*2), f.Name, f.Size)
		}
	}
}

func main() {
	lines := puzzle.ReadInputLinesOrPanic(puzzle.GetInputFile(puzzleDay, puzzleInputPart, puzzleInputFile))
	listing := false
	dTree := NewDirectoryTree()

	for _, line := range lines {
		if strings.HasPrefix(line, "$ ") {
			cmd := extractCommand(line)

			listing = false

			switch cmd {
			case "cd":
				dTree.ChangeDirectory(extractChangeDirectoryPath(line))
			case "ls":
				listing = true
			default:
				panic(fmt.Errorf("unknown command: %s", cmd))
			}

			continue
		}

		if !listing {
			panic(fmt.Errorf("invalid input: %s - expected command", line))
		}

		if strings.HasPrefix(line, "dir ") {
			if d, err := NewDirectoryFromString(line); err != nil {
				panic(err)
			} else {
				dTree.AddDirectoryAtCurrentPath(d)
			}
		} else {
			if f, err := NewFileFromString(line); err != nil {
				panic(err)
			} else {
				dTree.AddFileAtCurrentPath(f)
			}
		}
	}

	if verbose {
		dTree.Dump()
	}

	var directoryToDelete *Directory

	totalUsedSpace := dTree.GetDirectoryFromPath("/").TotalSize
	additionalFreeSpaceRequired := requiredFreeDiskSpace - (totalDiskSpace - totalUsedSpace)

	println("disk space: required", requiredFreeDiskSpace, "total", totalDiskSpace, "used", totalUsedSpace)
	println("additional space required", additionalFreeSpaceRequired)

	for _, d := range dTree.GetDirectories() {
		println("checking directory size", dTree.GetFullDirectoryPath(d), d.TotalSize)

		if d.TotalSize > additionalFreeSpaceRequired && (directoryToDelete == nil || d.TotalSize < directoryToDelete.TotalSize) {
			directoryToDelete = d
		}
	}

	println("total size of directory to delete:", directoryToDelete.TotalSize)
}

func extractCommand(line string) string {
	match := commandRegexp.FindAllStringSubmatch(line, -1)

	if len(match) > 0 {
		return match[0][1]
	}

	return ""
}

func extractChangeDirectoryPath(line string) string {
	match := changeDirectoryRegexp.FindAllStringSubmatch(strings.TrimSpace(line), -1)

	if len(match) > 0 {
		return match[0][1]
	}

	return ""
}
