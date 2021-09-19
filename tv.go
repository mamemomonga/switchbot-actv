package main

import (
	"github.com/rivo/tview"
)

func viewTV() {
	for {
		app := tview.NewApplication()
		stop := func() {
			app.Stop()
		}
		list := tview.NewList().
			AddItem("電源入/切", "電源ボタンを押します", '1', stop).
			AddItem("チャンネル1", "リモコンのチャンネル1を選択します", '2', stop).
			AddItem("次のチャンネル", "次のチャンネルに移動します", '3', stop).
			AddItem("前のチャンネル", "前のチャンネルに移動します", '4', stop).
			AddItem("音量を上げる", "ボリュームを上げます", '5', stop).
			AddItem("音量を下げる", "ボリュームを下げます", '6', stop).
			AddItem("戻る", "機器選択に戻ります", '0', stop)
		if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
			panic(err)
		}
		switch list.GetCurrentItem() {
		case 0:
			sb.TVPower()
		case 1:
			sb.TVChannelOne()
		case 2:
			sb.TVChannelUp(true)
		case 3:
			sb.TVChannelUp(false)
		case 4:
			sb.TVVolumeUp(true)
		case 5:
			sb.TVVolumeUp(false)
		default:
			return
		}
	}
}
