package main

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func viewAC() {
	for {
		app := tview.NewApplication()
		stop := func() {
			app.Stop()
		}

		jdFanSpeed := []string{"", "自動", "弱", "中", "強"}
		jdMode := []string{"", "自動", "冷房", "乾燥", "送風", "暖房"}

		fanSpeed := jdFanSpeed[configs.DeviceAC.Speed]
		mode := jdMode[configs.DeviceAC.Mode]

		list := tview.NewList().
			AddItem("電源入", "電源を入れます", '1', stop).
			AddItem("電源切", "電源を切ります", '2', stop).
			AddItem("温度", fmt.Sprintf("温度を変更します (%d度)", configs.DeviceAC.Temp), '3', stop).
			AddItem("風速", fmt.Sprintf("風速を切り替えます (%s)", fanSpeed), '4', stop).
			AddItem("モード", fmt.Sprintf("冷暖房を切り替えます (%s)", mode), '5', stop).
			AddItem("戻る", "機器選択に戻ります", '0', stop)
		if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
			panic(err)
		}
		switch list.GetCurrentItem() {
		case 0:
			sb.ACOn(true)
		case 1:
			sb.ACOn(false)
		case 2:
			viewACTemp()
		case 3:
			viewSpeed()
		case 4:
			viewMode()
		case 5:
			return
		}
	}
}

func viewSpeed() {
	for {
		app := tview.NewApplication()
		stop := func() {
			app.Stop()
		}
		list := tview.NewList().
			AddItem("自動", "", '1', stop).
			AddItem("弱", "", '2', stop).
			AddItem("中", "", '3', stop).
			AddItem("強", "", '4', stop).
			AddItem("戻る", "エアコンに戻ります", '0', stop)
		if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
			panic(err)
		}
		switch list.GetCurrentItem() {
		case 0:
			configs.DeviceAC.Speed = 1
			sb.ACOn(true)
			configSave()
			return

		case 1:
			configs.DeviceAC.Speed = 2
			sb.ACOn(true)
			configSave()
			return

		case 2:
			configs.DeviceAC.Speed = 3
			sb.ACOn(true)
			configSave()
			return

		case 3:
			configs.DeviceAC.Speed = 4
			sb.ACOn(true)
			configSave()
			return

		default:
			return
		}
	}
}

func viewMode() {
	for {
		app := tview.NewApplication()
		stop := func() {
			app.Stop()
		}
		list := tview.NewList().
			AddItem("自動", "", '1', stop).
			AddItem("冷房", "", '2', stop).
			AddItem("乾燥", "", '3', stop).
			AddItem("送風", "", '4', stop).
			AddItem("暖房", "", '5', stop).
			AddItem("戻る", "エアコンに戻ります", '0', stop)
		if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
			panic(err)
		}
		switch list.GetCurrentItem() {
		case 0:
			configs.DeviceAC.Mode = 1
			sb.ACOn(true)
			configSave()
			return

		case 1:
			configs.DeviceAC.Mode = 2
			sb.ACOn(true)
			configSave()
			return

		case 2:
			configs.DeviceAC.Mode = 3
			sb.ACOn(true)
			configSave()
			return

		case 3:
			configs.DeviceAC.Mode = 4
			sb.ACOn(true)
			configSave()
			return

		case 4:
			configs.DeviceAC.Mode = 4
			sb.ACOn(true)
			configSave()
			return

		default:
			return
		}
	}
}

func viewACTemp() {
	app := tview.NewApplication()
	inputField := tview.NewInputField().
		SetLabel("エアコン 設定温度(摂氏・数値2桁): ").
		SetAcceptanceFunc(tview.InputFieldInteger).
		SetFieldWidth(2).
		SetDoneFunc(func(key tcell.Key) {
			app.Stop()
		})
	if err := app.SetRoot(inputField, true).SetFocus(inputField).Run(); err != nil {
		panic(err)
	}
	buf := inputField.GetText()
	if buf == "" {
		return
	}
	configs.DeviceAC.Temp, _ = strconv.Atoi(buf)
	sb.ACOn(true)
	configSave()
}
