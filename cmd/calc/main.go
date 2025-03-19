package main

import (
	"calc/gui"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	c := gui.NewCalculatorUI(a)
	c.Window.ShowAndRun()
}
