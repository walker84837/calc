package gui

import (
	"calc/pkg/calculator"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type CalculatorUI struct {
	expression *widget.Label
	Window     fyne.Window
}

func NewCalculatorUI(app fyne.App) *CalculatorUI {
	w := app.NewWindow("Calcolatrice")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(400, 400))

	expression := widget.NewLabel("")
	expression.Alignment = fyne.TextAlignTrailing

	buttons := container.NewGridWithColumns(4,
		widget.NewButton("C", func() { expression.SetText("") }),
		widget.NewButton("(", func() { expression.SetText(expression.Text + "(") }),
		widget.NewButton(")", func() { expression.SetText(expression.Text + ")") }),
		widget.NewButton("/", func() { expression.SetText(expression.Text + "/") }),
		widget.NewButton("7", func() { expression.SetText(expression.Text + "7") }),
		widget.NewButton("8", func() { expression.SetText(expression.Text + "8") }),
		widget.NewButton("9", func() { expression.SetText(expression.Text + "9") }),
		widget.NewButton("*", func() { expression.SetText(expression.Text + "*") }),
		widget.NewButton("4", func() { expression.SetText(expression.Text + "4") }),
		widget.NewButton("5", func() { expression.SetText(expression.Text + "5") }),
		widget.NewButton("6", func() { expression.SetText(expression.Text + "6") }),
		widget.NewButton("-", func() { expression.SetText(expression.Text + "-") }),
		widget.NewButton("1", func() { expression.SetText(expression.Text + "1") }),
		widget.NewButton("2", func() { expression.SetText(expression.Text + "2") }),
		widget.NewButton("3", func() { expression.SetText(expression.Text + "3") }),
		widget.NewButton("+", func() { expression.SetText(expression.Text + "+") }),
		widget.NewButton("0", func() { expression.SetText(expression.Text + "0") }),
		widget.NewButton(".", func() { expression.SetText(expression.Text + ".") }),
		widget.NewButton("=", func() {
			result, err := calculator.Evaluate(expression.Text)
			if err != nil {
				expression.SetText("Error")
			} else {
				expression.SetText(result)
			}
		}),
	)

	w.SetContent(container.NewVBox(
		expression,
		buttons,
	))

	return &CalculatorUI{
		expression: expression,
		Window:     w,
	}
}
