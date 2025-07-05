package main

import (
	//"fmt"

	"fyne.io/fyne/v2"
	//"fmt"
	//"github.com/spf13/cobra"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	output *widget.Label
}

var myApp App

func main() {
	//var rootCmd = &cobra.Command{
	//	Use:   "hello",
	//	Short: "Hello CLI",
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("Hello, Cobra CLI!")
	//	},
	//}
	a := app.New()
	w := a.NewWindow("Hello World")
	output, entry, btn := myApp.makeUI()
	w.SetContent(container.NewVBox(output, entry, btn))
	w.Resize(fyne.Size{Width: 500, Height: 500})
	w.ShowAndRun()
}

func (app *App) makeUI() (*widget.Label, *widget.Entry, *widget.Button) {
	output := widget.NewLabel("Hello world")
	entry := widget.NewEntry()
	btn := widget.NewButton("Enter", func() {
		app.output.SetText(entry.Text)
	})
	btn.Importance = widget.HighImportance
	app.output = output

	return output, entry, btn
}
