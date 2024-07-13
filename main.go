package main

import (
	"log"
	"os"

	"example.com/switchbot-cli/cfg"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	configs *cfg.Cfg
	sb      *SwitchBot
)

func main() {
	var err error
	configs, err = cfg.New("config.json")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	if configs.APIToken.Token == "" {
		viewInputToken()
	}
	if configs.APIToken.Secret == "" {
		viewInputToken()
	}
	sb = NewSwitchBot(configs)
	sb.SearchDevice()
	if err := configs.Save(); err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	for {
		viewDeviceSelect()
	}
}

func inputText(s string) string {
	app := tview.NewApplication()
	inputField := tview.NewInputField().
		SetLabel(s + ": ").
		SetFieldWidth(100).
		SetDoneFunc(func(key tcell.Key) {
			app.Stop()
		})
	if err := app.SetRoot(inputField, true).SetFocus(inputField).Run(); err != nil {
		panic(err)
	}
	buf := inputField.GetText()
	if buf == "" {
		os.Exit(0)
	}
	return buf
}

func viewInputToken() {
	configs.APIToken.Token = inputText("SwitchBot トークンを入力")
	configs.APIToken.Secret = inputText("SwitchBot クライアントシークレットを入力")
	configSave()
}

func configSave() {
	if err := configs.Save(); err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}

func viewDeviceSelect() {
	app := tview.NewApplication()
	stop := func() {
		app.Stop()
	}
	list := tview.NewList().
		AddItem("エアコン", "エアコンを操作します", '1', stop).
		AddItem("テレビ", "テレビを操作します", '2', stop).
		AddItem("終了", "アプリケーションを終了します", '0', stop)
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		log.Fatal(err)
	}
	switch list.GetCurrentItem() {
	case 0:
		viewAC()
	case 1:
		viewTV()
	case 2:
		os.Exit(0)
	}
}
