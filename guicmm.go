package main

import (
	//"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"io"
	"strings"
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

// var myApp App
var cfg config

func main() {
	a := app.NewWithID("GUI")
	win := a.NewWindow("Mark down")

	edit, preview := cfg.makeUI()
	cfg.createMenuItems(win)
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

func (app *config) createMenuItems(win fyne.Window) {
	openMenuItem := fyne.NewMenuItem("Open...", app.openFunc(win))
	saveMenuItem := fyne.NewMenuItem("Save...", func() {})
	app.SaveMenuItem = saveMenuItem
	app.SaveMenuItem.Disabled = true

	// create a file menu, and add the three items to it
	saveAsMenu := fyne.NewMenuItem("Save as...", app.saveAsFunc(win))

	// create a main menu, and add the file menu to it
	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenu)

	menu := fyne.NewMainMenu(fileMenu)
	win.SetMainMenu(menu)
}

var filter = storage.NewExtensionFileFilter([]string{".md", ".MD"})

func (app *config) openFunc(win fyne.Window) func() {
	return func() {
		startLocation, _ := storage.ListerForURI(storage.NewFileURI("C:/Temp"))
		openDialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if read == nil {
				return
			}

			fileURI := read.URI()
			defer read.Close()

			data, err := io.ReadAll(read)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			app.EditWidget.SetText(string(data))
			app.CurrentFile = fileURI
			win.SetTitle(win.Title() + " - " + fileURI.Name())
			app.SaveMenuItem.Disabled = false
		}, win)

		openDialog.SetLocation(startLocation)
		openDialog.SetFilter(filter)
		openDialog.Show()
	}
}

func (app *config) saveAsFunc(win fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if write == nil {
				// user cancel
				return
			}
			write.Write([]byte(app.EditWidget.Text))
			app.CurrentFile = write.URI()
			defer write.Close()
			if !strings.HasSuffix(strings.ToLower(write.URI().String()), ".md") {
				dialog.ShowInformation("Error", "Please name your file .md extension", win)
				return
			}
			win.SetTitle(win.Title() + " - " + write.URI().Name())
			app.SaveMenuItem.Disabled = false
		}, win)
		saveDialog.SetFileName("untitled.md")
		saveDialog.SetFilter(filter)
		saveDialog.Show()
	}
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
