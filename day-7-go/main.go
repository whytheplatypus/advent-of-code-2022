package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var Root = &dir{
	subdirs: map[string]*dir{},
}

type dir struct {
	size    int
	subdirs map[string]*dir
	parent  *dir
}

func (d *dir) cd(name string) (*dir, error) {
	if name == ".." {
		return d.parent, nil
	}
	if name == "/" {
		return Root, nil
	}
	next, ok := d.subdirs[name]
	if !ok {
		return nil, fmt.Errorf("failed to find subdir %s", name)
	}
	return next, nil
}

func (d *dir) ls(contents string) error {
	log.Println("ls called with", contents)
	list := strings.Split(contents, "\n")
	for _, item := range list {
		parts := strings.Split(item, " ")
		if len(parts) != 2 {
			continue
		}
		log.Println("handeling", item, parts)
		if parts[0] == "dir" {
			d.subdirs[parts[1]] = &dir{
				parent:  d,
				subdirs: map[string]*dir{},
			}
			log.Println("added dir", parts[1])
			continue
		}
		size, err := strconv.Atoi(parts[0])
		if err != nil {
			return err
		}
		d.size += size
	}
	return nil
}

func (d *dir) Size() int {
	total := 0
	for _, sub := range d.subdirs {
		total += sub.Size()
	}
	total += d.size
	return total
}

func walk(d *dir) int {
	total := 0

	size := d.Size()
	log.Println("size", size)
	if size < 100000 {
		total += size
	}

	for _, sub := range d.subdirs {
		total += walk(sub)
	}

	return total
}

func walkAtLeast(d *dir, need int, have int) int {
	size := d.Size()
	if size >= need && size < have {
		have = size
	}

	for _, sub := range d.subdirs {
		have = walkAtLeast(sub, need, have)
	}
	return have
}

var cmd = regexp.MustCompile(`(?m:^\$ ((?:cd)|(?:ls))(?: (.+))?$)`)

const FileSystemSize = 70000000

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(input))
	commands := cmd.FindAllIndex(input, -1)
	// pop off root
	commands = commands[1:]
	log.Println(commands)

	d := Root
	for i, command := range commands {
		parts := cmd.FindAllSubmatch(input[command[0]:command[1]], -1)
		log.Printf("%+q", parts)
		if string(parts[0][1]) == "ls" {
			log.Println("calling ls")
			endResults := len(input)
			if i+1 < len(commands) {
				endResults = commands[i+1][0]
			}
			// everything until the next command...?
			d.ls(string(input[command[1]:endResults]))
		}
		if string(parts[0][1]) == "cd" {
			d, err = d.cd(string(parts[0][2]))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	log.Println(walk(Root))

	need := 30000000 - (FileSystemSize - Root.Size())
	log.Println("need", need)
	log.Println(walkAtLeast(Root, need, Root.Size()))
}
