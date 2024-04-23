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
	titleElement         ui.Static
	window               ui.WindowMain
	buttonImportJsonData ui.Button
	txtJsonDataPicker    ui.Button
	jsonData             Interactions
}

func (s *WindowSettingsStruct) setTitle(title string) {
	s.title = title
}
func (s *WindowSettingsStruct) getTitle() string {
	return s.title
}

type Interactions struct {
	InteractionsTable []InteractionsTable `json:"interactionsTable"`
}
type InteractionsTable struct {
	AttributeKey   string `json:"attributeKey"`
	AttributeValue string `json:"attributeValue"`
}

func WindowSettings() *WindowSettingsStruct {
	p := &WindowSettingsStruct{}
	p.setTitle("Settings")

	wnd := ui.NewWindowMain(
		ui.WindowMainOpts().
			Title(p.getTitle()).
			ClientArea(win.SIZE{Cx: 300, Cy: 200}),
	)

	me := &WindowSettingsStruct{
		titleElement: ui.NewStatic(wnd,
			ui.StaticOpts().
				Text(p.getTitle()).
				Position(win.POINT{X: 130, Y: 10}),
		),

		window: wnd,

		txtJsonDataPicker: ui.NewButton(wnd,
			ui.ButtonOpts().
				Text("Json Data Picker").
				Position(win.POINT{X: 10, Y: 19}),
		),
		buttonImportJsonData: ui.NewButton(wnd,
			ui.ButtonOpts().
				Text("&Import").
				Position(win.POINT{X: 190, Y: 19}),
		),
	}

	me.txtJsonDataPicker.On().BnClicked(func() {
		filepath, err := dialog.File().Filter("JSON Files", "json").Load()
		if err != nil {
			me.window.Hwnd().MessageBox("error file pick: "+err.Error(), p.getTitle(), co.MB_ICONERROR)
		}
		file, err := os.Open(filepath)
		if err != nil {
			me.window.Hwnd().MessageBox("error open file: "+err.Error(), p.getTitle(), co.MB_ICONERROR)
		}
		defer file.Close()

		bytes, err := os.ReadFile(filepath)
		if err != nil {
			me.window.Hwnd().MessageBox("error read file: "+err.Error(), p.getTitle(), co.MB_ICONERROR)
		}
		data := Interactions{}
		if err := json.Unmarshal(bytes, &data); err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}
		me.jsonData = data
	})

	me.buttonImportJsonData.On().BnClicked(func() {
		fmt.Println(me.jsonData)
		err := me.window.Hwnd().DestroyWindow()
		if err != nil {
			me.window.Hwnd().MessageBox("error destroy window: "+err.Error(), p.getTitle(), co.MB_ICONERROR)
		}
	})
	return me
}
