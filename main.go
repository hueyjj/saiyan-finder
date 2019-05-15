package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
)

var (
	// CurrentWorkingDir is the current working directory
	CurrentWorkingDir = ""
)

type buf struct {
	queue []string
}

// Same as ioutil.ReadDir but no sorting
func readDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func main() {
	CurrentWorkingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current working directory:", CurrentWorkingDir)

	buf := buf{}

	newQueueItem := make(chan string)

	go func() {
		files, err := readDir(".")
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			path := filepath.FromSlash(fmt.Sprintf("%s/%s", CurrentWorkingDir, file.Name()))
			fmt.Println(path)
			newQueueItem <- path
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		select {
		case item := <-newQueueItem:
			buf.queue = append(buf.queue, item)
		case <-interrupt:
			fmt.Print(buf.queue)
			fmt.Println("Ctrl+C detected. Ending program")
			os.Exit(0)
			return
		}
	}
}
