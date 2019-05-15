package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
)

type file struct {
	name string
	path string
}

type commit struct {
	sync.Mutex
	files []string
}

type dirs struct {
	sync.Mutex
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
	if len(os.Args) != 2 {
		fmt.Println("saiyan-finder.exe [search term]")
		os.Exit(0)
	}
	searchTerm := os.Args[1]

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Current working directory:", cwd)

	addFile := make(chan *file)
	addDirPath := make(chan string)
	done := make(chan bool)

	commit := commit{}
	dirs := dirs{}
	dirs.queue = append(dirs.queue, cwd)

	var wg sync.WaitGroup
	go func() {
		for {
			if len(dirs.queue) <= 0 {
				done <- true
			}
			for i := 0; i < len(dirs.queue); i++ {
				wg.Add(1)

				var dir string
				dir, dirs.queue = dirs.queue[0], dirs.queue[1:]
				go func() {
					defer wg.Done()
					//var dir string
					//dirs.Lock()
					//if len(dirs.queue) > 0 {
					//	dir, dirs.queue = dirs.queue[0], dirs.queue[1:]
					//}
					//dirs.Unlock()

					//if len(dir) <= 0 {
					//	return
					//}

					files, err := readDir(dir)
					if err != nil {
						fmt.Println(err)
					}
					//fmt.Println("Working in:", dir)
					for _, f := range files {
						name := f.Name()
						path := filepath.FromSlash(fmt.Sprintf("%s/%s", dir, name))
						stat, err := os.Stat(path)
						if err != nil {
							fmt.Println(err)
							continue
						}
						switch mode := stat.Mode(); {
						case mode.IsDir():
							addDirPath <- path
						case mode.IsRegular():
							addFile <- &file{name: name, path: path}
						}
					}
				}()
			}
			wg.Wait()
		}
	}()

	fileIndexed := 0

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case dirPath := <-addDirPath:
			dirs.queue = append(dirs.queue, dirPath)
		case file := <-addFile:
			fileIndexed++
			commit.files = append(commit.files, file.path)
			if file.name == searchTerm {
				fmt.Println(file.path)
			}
			//fmt.Println(file.name, file.path)

		case <-interrupt:
			fmt.Println("Interrupt detected. Exiting program.")
			fmt.Println("Number of files indexed:", fileIndexed)
			os.Exit(1)
			return
		case <-done:
			fmt.Println("Number of files indexed:", fileIndexed)
			os.Exit(0)
			return
		}
	}
}
