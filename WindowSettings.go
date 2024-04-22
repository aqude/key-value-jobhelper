package main

import (
	"encoding/json"
	"fmt"
	"github.com/rodrigocfd/windigo/ui"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/sqweek/dialog"
	"os"
)

type WindowSettingsStruct struct {
	title                string
	window               ui.WindowMain
	buttonImportJsonData ui.Button
	txtJsonDataPicker    ui.Button
	jsonData             Interactions
}

type Interactions struct {
	InteractionsTable []InteractionsTable `json:"interactionsTable"`
}
type InteractionsTable struct {
	AttributeKey   string `json:"attributeKey"`
	AttributeValue string `json:"attributeValue"`
}

func WindowSettings() *WindowSettingsStruct {
	wnd := ui.NewWindowMain(
		ui.WindowMainOpts().
			Title("Settings").
			ClientArea(win.SIZE{Cx: 500, Cy: 200}),
	)

	me := &WindowSettingsStruct{
		title: "Settings",

		window: wnd,

		txtJsonDataPicker: ui.NewButton(wnd,
			ui.ButtonOpts().
				Text("Json Data Picker").
				Position(win.POINT{X: 10, Y: 19}),
		),
		buttonImportJsonData: ui.NewButton(wnd,
			ui.ButtonOpts().
				Text("&Import").
				Position(win.POINT{X: 240, Y: 19}),
		),
	}

	me.txtJsonDataPicker.On().BnClicked(func() {
		filepath, err := dialog.File().Filter("JSON", ".json").Load()
		if err != nil {
			me.window.Hwnd().MessageBox("error file pick: "+err.Error(), me.title, co.MB_ICONERROR)
		}
		file, err := os.Open(filepath)
		if err != nil {
			me.window.Hwnd().MessageBox("error open file: "+err.Error(), me.title, co.MB_ICONERROR)
		}
		defer file.Close()

		bytes, err := os.ReadFile(filepath)
		if err != nil {
			me.window.Hwnd().MessageBox("error read file: "+err.Error(), me.title, co.MB_ICONERROR)
		}
		data := Interactions{}
		if err := json.Unmarshal(bytes, &data); err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}
		me.jsonData = data
		fmt.Println(me.jsonData)
	})
	return me
}
