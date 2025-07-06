package main

import (
	//"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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
	openMenu := fyne.NewMenuItem("Open...", func() {})

	saveMenuItem := fyne.NewMenuItem("Save...", func() {})
	app.SaveMenuItem = saveMenuItem
	app.SaveMenuItem.Disabled = true
	saveAsMenu := fyne.NewMenuItem("Save as...", app.saveAsFunc(win))

	fileMenu := fyne.NewMenu("File", openMenu, saveMenuItem, saveAsMenu)

	menu := fyne.NewMainMenu(fileMenu)
	win.SetMainMenu(menu)
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

			win.SetTitle(win.Title() + " - " + write.URI().Name())
			app.SaveMenuItem.Disabled = false
		}, win)
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
