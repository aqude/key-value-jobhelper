package main

import (
	"fmt"
	"runtime"

	"github.com/rodrigocfd/windigo/ui"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

func main() {
	runtime.LockOSThread()

	myWindow := MainWindow()
	myWindow.window.RunAsMain()
}

type MainWindowStruct struct {
	window         ui.WindowMain
	lblName        ui.Static
	txtName        ui.Edit
	btnShow        ui.Button
	btnDevSettings ui.Button
}

func MainWindow() *MainWindowStruct {
	wnd := ui.NewWindowMain(
		ui.WindowMainOpts().
			Title("write key").
			ClientArea(win.SIZE{Cx: 500, Cy: 400}),
	)

	me := &MainWindowStruct{
		window: wnd,

		lblName: ui.NewStatic(wnd,
			ui.StaticOpts().
				Text("write key ").
				Position(win.POINT{X: 10, Y: 22}),
		),
		txtName: ui.NewEdit(wnd,
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
		msg := fmt.Sprintf("Hello, %s!", me.txtName.Text())
		me.window.Hwnd().MessageBox(msg, "Saying hello", co.MB_ICONINFORMATION)
	})

	me.btnDevSettings.On().BnClicked(func() {
		WindowSettings().window.RunAsMain()
	})

	return me
}
