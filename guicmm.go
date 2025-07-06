package main

import (
	//"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	output *widget.Label
}

type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

var myApp App
var cfg config

func main() {
	//var rootCmd = &cobra.Command{
	//	Use:   "hello",
	//	Short: "Hello CLI",
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("Hello, Cobra CLI!")
	//	},
	//}
	//a := app.New()
	//w := a.NewWindow("Hello World")
	//output, entry, btn := myApp.makeUI()
	//w.SetContent(container.NewVBox(output, entry, btn))
	//w.Resize(fyne.Size{Width: 500, Height: 500})
	//w.ShowAndRun()
	a := app.New()
	win := a.NewWindow("Mark down")

	edit, preview := cfg.makeUI()
	win.SetContent(container.NewHSplit(edit, preview))
	win.Resize(fyne.Size{Width: 800, Height: 500})
	win.CenterOnScreen()
	win.ShowAndRun()
}

func (app *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	app.EditWidget = edit
	app.PreviewWidget = preview
	edit.OnChanged = preview.ParseMarkdown
	return edit, preview
}

//func (app *App) makeUI() (*widget.Label, *widget.Entry, *widget.Button) {
//	output := widget.NewLabel("Hello world")
//	entry := widget.NewEntry()
//	btn := widget.NewButton("Enter", func() {
//		app.output.SetText(entry.Text)
//	})
//	btn.Importance = widget.HighImportance
//	app.output = output
//
//	return output, entry, btn
//}
