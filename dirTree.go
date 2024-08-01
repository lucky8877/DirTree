package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DirTree(pathName string, st int, mod string) error {

	dir, err := os.Open(pathName)
	if err != nil {
		return err
	}
	defer dir.Close()

	files, err := dir.ReadDir(-1)
	if err != nil {
		return err
	}
	for _, fi := range files {
		fl, _ := os.Open(pathName + "/" + fi.Name())
		fk, _ := fl.Stat()
		if mod == "-f" && fk.IsDir() == false {
			fmt.Println(strings.Repeat("\t", st) + "├───" + fi.Name() + "(" + strconv.Itoa(int(fk.Size())) + "b)")
		} else if fk.IsDir() == true {
			fmt.Println(strings.Repeat("\t", st) + "├───" + fi.Name())
			DirTree(dir.Name()+"/"+fi.Name(), st+1, "")
		}
	}
	return nil
}

func main() {
	args := os.Args[1:]
	if len(args) != 0 {
		if args[0] == "info" {
			fmt.Println("specify the path to the directory to display the directory tree\nspecify the -f option to display files and their sizes")
		} else if len(args) > 0 && len(args) == 2 && args[1] == "-f" {
			DirTree(args[0], 0, "-f")
		} else if 0 < len(args) && len(args) < 2 {
			DirTree(args[0], 0, ".")
		} else {
			err := errors.New("incorrect command")
			fmt.Println(err)

		}
	} else {
		err := errors.New("path not specified")
		fmt.Println(err)
	}

}
