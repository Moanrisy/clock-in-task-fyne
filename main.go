package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func setupTmpFile() (string, *os.File, error) {

	dirname := os.TempDir()
	filePath := dirname + "/clock-in-title"
	var file *os.File
	var err error

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File does not exist, proceeds to create /tmp/clock-in-title")
		file, err = os.Create(filePath)
		checkErr(err)
	} else {
		file, err = os.Open(filePath)
		checkErr(err)
	}

	return filePath, file, err
}

func main() {
	a := app.New()
	w := a.NewWindow("Clock in Task")
	filePath, file, err := setupTmpFile()

	cit := widget.NewLabel("Please clock in your task")
	w.SetContent(cit)

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	initialStat, err := os.Stat(filePath)
	checkErr(err)
	go func() {
		for {
			stat, err := os.Stat(filePath)
			checkErr(err)
			if stat.Size() != initialStat.Size() || stat.ModTime() != initialStat.ModTime() {

				buf := make([]byte, stat.Size())
				_, err := file.Read(buf)
				fmt.Println(string(buf))
				cit.SetText(string(buf))

				initialStat, err = os.Stat(filePath)
				checkErr(err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	w.Resize(fyne.NewSize(200, 50))
	w.ShowAndRun()
}
