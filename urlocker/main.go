// package urlocker  // after complete code use this package
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"go.mongodb.org/mongo-driver/mongo"

)

func createEntry() {
	fmt.Println("inset entry")
}
func updateEntry() {
	fmt.Println("update entry")
}
func readEntry() {
	fmt.Println("read entry")
}
func deleteEntry() {
	fmt.Println("delete entry")
}

type MongoFields struct {
	FieldStr  string `json:"Field Str"`
	FieldInt  int    `json:"Field Int"`
	FieldBool bool   `json:"Field Bool"`
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Button Widget")

	name := widget.NewEntry()
	name.SetPlaceHolder("url title")
	email := widget.NewEntry()
	email.SetPlaceHolder("full url address")
	password := widget.NewEntry()
	password.SetPlaceHolder("discription")

	add := widget.NewButton("ADD URL", func() {
		log.Println("tapped")
	})

	show := widget.NewButtonWithIcon("Show All", theme.HomeIcon(), func() {
		log.Println("tapped home")
	})

	myWindow.SetContent(widget.NewVBox(name, email, password, add, show))
	myWindow.ShowAndRun()
// ----------------------------------------------------------------------------------------------------------

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://maninder:147105@mike-cluster.lobxs.mongodb.net/urlocker?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err, client)
	}

}
