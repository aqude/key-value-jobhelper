package main

import (
	"github.com/rodrigocfd/windigo/ui"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"runtime"
)

func main() {
	updateMainWindow(Interactions{})
}

type MainWindowStruct struct {
	window            ui.WindowMain
	lblName           ui.Static
	txtKey            ui.Edit
	btnShow           ui.Button
	btnDevSettings    ui.Button
	statusLoadData    ui.Static
	txtStatusLoadData string
}

func updateMainWindow(data Interactions) {
	var text string
	if data.InteractionsTable == nil {
		text = "None set data"
	} else {
		text = "App ready"
	}
	runtime.LockOSThread()
	myWindow := MainWindow(data, text)
	myWindow.window.RunAsMain()

	if myWindow.window.Hwnd() != 0 {
		myWindow.window.Hwnd().SendMessage(co.WM_CLOSE, 0, 0)
	}
}
func MainWindow(data Interactions, text string) *MainWindowStruct {

	wnd := ui.NewWindowMain(
		ui.WindowMainOpts().
			Title("write key").
			ClientArea(win.SIZE{Cx: 500, Cy: 100}),
	)

	me := &MainWindowStruct{
		statusLoadData: ui.NewStatic(wnd, ui.StaticOpts().Text(text).Position(win.POINT{X: 10, Y: 2})),
		window:         wnd,
		lblName: ui.NewStatic(wnd,
			ui.StaticOpts().
				Text("write key ").
				Position(win.POINT{X: 10, Y: 22}),
		),
		txtKey: ui.NewEdit(wnd,
			ui.EditOpts().
				Position(win.POINT{X: 80, Y: 20}).
				Size(win.SIZE{Cx: 150}),
		),
		btnShow: ui.NewButton(wnd,
			ui.ButtonOpts().
				Text("&Show").
				Position(win.POINT{X: 250, Y: 19}),
		),
		btnDevSettings: ui.NewButton(wnd,
			ui.ButtonOpts().
				Text("&Settings").
				Position(win.POINT{X: 350, Y: 19}),
		),
	}

	me.btnShow.On().BnClicked(func() {
		key := me.txtKey.Text()

		if key == "" {
			me.window.Hwnd().MessageBox("Write key", "Error", co.MB_ICONERROR)
		} else {
			for i := 0; i < len(data.InteractionsTable); i++ {
				if data.InteractionsTable[i].AttributeKey == key {
					me.window.Hwnd().MessageBox(data.InteractionsTable[i].AttributeValue, "Key", co.MB_ICONINFORMATION)
				}
			}
		}
	})

	me.btnDevSettings.On().BnClicked(func() {
		WindowSettings().window.RunAsMain()
	})

	return me
}
