package main

import (
	"strings"

	"github.com/JacieChao/what2do/pkg"
	"github.com/JacieChao/what2do/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&theme.MyTheme{})
	w := a.NewWindow("Let's plan what to do today!")

	place := widget.NewLabel("")
	todo := widget.NewLabel("")
	w.SetContent(container.NewVBox(
		place,
		todo,
		widget.NewButton("Let's Go!", func() {
			result := pkg.Random()
			for key, value := range result {
				place.SetText(key)
				todo.SetText(strings.Join(value, ","))
			}
		}),
	))
	w.Resize(fyne.NewSize(480, 360))
	w.ShowAndRun()
}

func Submit() {

}
