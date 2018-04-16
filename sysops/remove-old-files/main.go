package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

func main() {
	pth := "/tmp/test"
	if _, err := os.Stat(pth); os.IsNotExist(err) {
		os.MkdirAll(pth, 0755)
	}

	for {
		now := time.Now()
		filename := now.Format("file-2006-01-02T15-04-05.txt")
		f, err := os.Create(path.Join(pth, filename))
		if err != nil {
			log.Fatal(err)
		}
		f.Close()

		fileInfo, err := ioutil.ReadDir(pth)
		if err != nil {
			log.Fatal(err)
		}
		for _, info := range fileInfo {
			if diff := now.Sub(info.ModTime()); diff > 2*time.Minute {
				fmt.Println("Deleting", path.Join(pth, info.Name()))
				err = os.Remove(path.Join(pth, info.Name()))
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		time.Sleep(10 * time.Second)
	}
}
