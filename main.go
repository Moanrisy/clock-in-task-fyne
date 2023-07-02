package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
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
	a.Settings().SetTheme(&myTheme{})
	w := a.NewWindow("Clock in Task")
	filePath, file, err := setupTmpFile()

	cit := canvas.NewText("Please clock in your task", theme.TextColor())
	cit.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	cit.TextSize = 15.0
	container := container.NewVBox(cit)

	w.SetContent(container)

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

				bufString := string(buf)
				bufString = strings.Replace(bufString, "\n", "", -1)

				cit.Text = bufString
				cit.Color = color.RGBA{R: 50, G: 205, B: 50, A: 255}

				cit.Refresh()

				initialStat, err = os.Stat(filePath)
				checkErr(err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	w.Resize(fyne.NewSize(200, 50))
	w.ShowAndRun()
}
