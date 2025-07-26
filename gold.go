package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"log"
	"os"
)

type Config struct {
	App        fyne.App
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
	MainWindow fyne.Window
}

var myApp Config

func main() {
	fyneApp := app.NewWithID("gold")
	myApp.App = fyneApp
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.InfoLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	myApp.MainWindow = fyneApp.NewWindow("Gold")
	myApp.MainWindow.Resize(fyne.NewSize(300, 200))
	myApp.MainWindow.FixedSize()
	myApp.MainWindow.SetMaster()

	myApp.makeUI()
	
	myApp.MainWindow.ShowAndRun()
}
