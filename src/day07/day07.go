package day07

import (
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

type Day07 struct {
	fs  map[string]int
	dir string
	ds  []string
}

func (d *Day07) ChangeDir(path string) {
	if strings.HasPrefix(path, "/") {
		d.dir = path
	} else {
		d.dir = filepath.Join(d.dir, path)
	}
}

func (d Day07) Pwd() string {
	return d.dir
}

func (d *Day07) Discover(file string) {
	if d.fs == nil {
		d.fs = map[string]int{}
	}
	if d.ds == nil {
		d.ds = []string{}
	}

	fParts := strings.Split(file, " ")
	if strings.HasPrefix(file, "dir") {
		// File is a directory, do nothing
		d.ds = append(d.ds, d.dir)
	} else {
		size, _ := strconv.Atoi(fParts[0])
		d.fs[filepath.Join(d.dir, fParts[1])] = size
	}
}

func (d Day07) List() []string {
	fs := []string{}
	for dir := range d.fs {
		if filepath.Dir(dir) == d.dir {
			fs = append(fs, filepath.Base(dir))
		}
	}
	return fs
}

func (d Day07) SumOfDirsMaxNBytes(n int) int {
	sum := 0
	log.Printf("%+v", d.ds)
	return sum
}
