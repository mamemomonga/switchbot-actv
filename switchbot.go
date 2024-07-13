package main

import (
	"context"
	"log"

	"example.com/switchbot-cli/cfg"
	// "github.com/davecgh/go-spew/spew"

	goswitchbot "github.com/nasa9084/go-switchbot"
)

type SwitchBot struct {
	configs *cfg.Cfg
	client  *goswitchbot.Client
	device  *goswitchbot.DeviceService
}

func NewSwitchBot(c *cfg.Cfg) *SwitchBot {
	t := new(SwitchBot)
	t.configs = c
	t.client = goswitchbot.New(t.configs.APIToken.Token, t.configs.APIToken.Secret)
	t.device = t.client.Device()
	return t
}

func (t *SwitchBot) connect() *goswitchbot.DeviceService {
	c := goswitchbot.New(t.configs.APIToken.Token, t.configs.APIToken.Secret)
	d := c.Device()
	return d
}

func (t *SwitchBot) SearchDevice() {
	log.Println("SwitchBot IRデバイスを検索します")
	_, irdev, _ := t.device.List(context.Background())
	for i := range irdev {
		switch irdev[i].Type {
		case "TV":
			log.Println("テレビを発見しました")
			t.configs.DeviceTV.ID = irdev[i].ID
			t.configs.DeviceTV.Name = irdev[i].Name
			log.Printf("  Name: %s / ID: %s", t.configs.DeviceTV.Name, t.configs.DeviceTV.ID)
		case "Air Conditioner":
			log.Println("エアコンを発見しました")
			t.configs.DeviceAC.ID = irdev[i].ID
			t.configs.DeviceAC.Name = irdev[i].Name
			log.Printf("  Name: %s / ID: %s", t.configs.DeviceAC.Name, t.configs.DeviceAC.ID)
			if t.configs.DeviceAC.Speed == 0 {
				t.configs.DeviceAC.Speed = 1
			}
			if t.configs.DeviceAC.Mode == 0 {
				t.configs.DeviceAC.Mode = 1
			}
			if t.configs.DeviceAC.Temp == 0 {
				t.configs.DeviceAC.Temp = 20
			}
		}
	}
}

func (t *SwitchBot) ACOn(on bool) {
	pwr := goswitchbot.PowerOn
	if !on {
		pwr = goswitchbot.PowerOff
	}

	ctx := context.Background()
	err := t.device.Command(
		ctx,
		t.configs.DeviceAC.ID,
		goswitchbot.ACSetAllCommand(
			t.configs.DeviceAC.Temp,
			goswitchbot.ACMode(t.configs.DeviceAC.Mode),
			goswitchbot.ACFanSpeed(t.configs.DeviceAC.Speed),
			pwr,
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	//	log.Printf("設定温度: %d度", t.configs.DeviceAC.Temp)
	//	log.Printf("モード:   %d", t.configs.DeviceAC.Mode)
	//	log.Printf("風速:     %d", t.configs.DeviceAC.Speed)

}

func (t *SwitchBot) TVPower() {
	t.connect()
	ctx := context.Background()
	t.device.Command(
		ctx,
		t.configs.DeviceTV.ID,
		goswitchbot.DeviceCommandRequest{
			CommandType: "command",
			Command:     "turnOn",
			Parameter:   "default",
		},
	)
}
func (t *SwitchBot) TVChannelUp(up bool) {
	cmd := "channelSub"
	if up {
		cmd = "channelAdd"
	}
	ctx := context.Background()
	t.device.Command(
		ctx,
		t.configs.DeviceTV.ID,
		goswitchbot.DeviceCommandRequest{
			CommandType: "command",
			Command:     cmd,
			Parameter:   "default",
		},
	)
}
func (t *SwitchBot) TVChannelOne() {
	ctx := context.Background()
	t.device.Command(
		ctx,
		t.configs.DeviceTV.ID,
		goswitchbot.DeviceCommandRequest{
			CommandType: "command",
			Command:     "SetChannel",
			Parameter:   "1",
		},
	)
}

func (t *SwitchBot) TVVolumeUp(up bool) {
	cmd := "volumeSub"
	if up {
		cmd = "volumeAdd"
	}
	ctx := context.Background()
	t.device.Command(
		ctx,
		t.configs.DeviceTV.ID,
		goswitchbot.DeviceCommandRequest{
			CommandType: "command",
			Command:     cmd,
			Parameter:   "default",
		},
	)
}
