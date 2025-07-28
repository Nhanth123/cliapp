package main

import "fyne.io/fyne/v2/container"

func (app *Config) makeUI() {
	// get current price
	openPrice, currentPrice, priceChange := app.getPriceText()

	//put infor in container
	priceContent := container.NewGridWithColumns(3, openPrice, currentPrice, priceChange)
}
