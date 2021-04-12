package main

import (
	"MCSkinStealer/downloader"

	ui "github.com/VladimirMarkelov/clui"
)

func CreateForm() {
	form := ui.AddWindow(0, 0, 40, 10, "Stiahnuť Minecraft skin")

	form.SetPack(ui.Vertical)
	form.SetPaddings(2, 2)
	form.SetGaps(ui.AutoSize, 1)
	form.SetActiveBackColor(ui.ColorGreen)
	form.SetActiveTextColor(ui.ColorWhite)
	form.SetBackColor(ui.ColorBlue)
	form.SetBorder(ui.BorderThick)

	frmPlayer := ui.CreateFrame(form, ui.AutoSize, 1, ui.BorderNone, 1)
	frmPlayer.SetActiveBackColor(ui.ColorGreen)
	frmPlayer.SetActiveTextColor(ui.ColorWhite)
	frmPlayer.SetBackColor(ui.ColorBlue)
	ui.CreateLabel(frmPlayer, ui.AutoSize, ui.AutoSize, "Hráč: ", 1)
	playerInput := ui.CreateEditField(frmPlayer, ui.AutoSize, "", 1)

	frmControl := ui.CreateFrame(form, ui.AutoSize, 4, ui.BorderNone, 1)
	frmControl.SetActiveBackColor(ui.ColorGreen)
	frmControl.SetActiveTextColor(ui.ColorWhite)
	frmControl.SetBackColor(ui.ColorBlue)
	btnSkin := ui.CreateButton(frmControl, ui.AutoSize, 4, "Stiahnuť skin", 1)
	btnBody := ui.CreateButton(frmControl, ui.AutoSize, 4, "Stiahnuť telo", 1)
	btnQuit := ui.CreateButton(frmControl, ui.AutoSize, 4, "Koniec", 1)

	btnQuit.OnClick(func(ev ui.Event) {
		go ui.Stop()
	})

	btnSkin.OnClick(func(ev ui.Event) {
		fileUrl := "https://minotar.net/skin/" + playerInput.Title()

		err := DownloadFile("skin_"+playerInput.Title()+".png", fileUrl)

		if err != nil {
			ui.CreateAlertDialog("Chyba", "Chyba sťahovania súboru.", "OK")
		}
	})

	btnBody.OnClick(func(ev ui.Event) {
		fileUrl := "https://minotar.net/body/" + playerInput.Title() + "/100.png"
		err := DownloadFile("body_"+playerInput.Title()+".png", fileUrl)

		if err != nil {
			ui.CreateAlertDialog("Chyba", "Chyba sťahovania súboru.", "OK")
		}
	})

}

func main() {
	ui.InitLibrary()

	defer ui.DeinitLibrary()

	CreateForm()

	ui.MainLoop()
}
