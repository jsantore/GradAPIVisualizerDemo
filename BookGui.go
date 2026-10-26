package main

import (
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	books           []Book
	authorsDisplay  *widget.Entry
	languageDisplay *widget.Entry
	publishDisplay  *widget.Entry
)

func buildBookGUI(bookData []Book) *fyne.Window {
	guiApp := app.New()
	books = bookData
	mainWindow := guiApp.NewWindow("Book GUI")
	directions := widget.NewLabel("Select a book")
	directions.Alignment = fyne.TextAlignCenter
	directions.TextStyle = fyne.TextStyle{Bold: true}
	titleList := widget.NewList(func() int { return len(books) },
		func() fyne.CanvasObject {
			return widget.NewLabel("Placeholder")
		},
		func(i widget.ListItemID, selected fyne.CanvasObject) {
			selected.(*widget.Label).SetText(books[i].Title)
		})
	titleList.OnSelected = bookSelected
	detailsPanel := makeDetailsPanel()
	mainLayout := container.NewBorder(directions, nil, nil, detailsPanel, titleList)
	mainWindow.SetContent(mainLayout)
	mainWindow.Resize(fyne.NewSize(800, 800))
	return &mainWindow
}

func makeDetailsPanel() *fyne.Container {
	detailsPanel := container.New(layout.NewVBoxLayout())
	authorsPanel := container.New(layout.NewVBoxLayout())
	languagePanel := container.New(layout.NewVBoxLayout())
	a := widget.NewLabel("Authors displayed below:")
	authorsPanel.Add(a)
	authorsDisplay = widget.NewEntry()
	authorsDisplay.MultiLine = true

	authorsDisplay.PlaceHolder = "Authors will appear here"
	authorsDisplay.Resize(fyne.NewSize(200, 100))
	authorsPanel.Add(authorsDisplay)
	detailsPanel.Add(authorsPanel)

	a = widget.NewLabel("Available in these Languages:")
	languagePanel.Add(a)
	languageDisplay = widget.NewEntry()
	languageDisplay.MultiLine = true
	languageDisplay.PlaceHolder = "Languages will appear here"
	languageDisplay.Resize(fyne.NewSize(200, 100))
	languagePanel.Add(languageDisplay)
	detailsPanel.Add(languagePanel)
	publishPanel := container.New(layout.NewVBoxLayout())
	a = widget.NewLabel("Publication Date:")
	publishPanel.Add(a)
	publishDisplay = widget.NewEntry()
	publishDisplay.PlaceHolder = "Publication Date will appear here"
	publishPanel.Add(publishDisplay)

	detailsPanel.Add(publishPanel)
	return detailsPanel
}

func bookSelected(i widget.ListItemID) {
	bookAuthors := books[i].AuthorName //AuthorName is a misnomer, it is a possibly empty list of authors
	if len(bookAuthors) == 0 {
		authorsDisplay.SetText("")
		return
	}
	var authorsNames = strings.Builder{}
	for _, author := range bookAuthors {
		authorsNames.WriteString(author + "\n")
	}
	authorsDisplay.SetText(authorsNames.String())
	pubDate := books[i].FirstPublishYear
	publishDisplay.SetText(strconv.Itoa(pubDate))
	languages := books[i].Language //this is again a lie is a possibly nil list of strings
	if languages == nil {
		languageDisplay.SetText("Languages not available")
		return
	}
	var languagesNames = strings.Builder{}
	for _, language := range languages {
		languagesNames.WriteString(language + "\n")
	}
	languageDisplay.SetText(languagesNames.String())
}
