package cfg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Cfg struct {
	filename string
	APIToken string
	DeviceTV CfgDeviceTV
	DeviceAC CfgDeviceAC
}

type CfgDeviceTV struct {
	ID    string
	Name  string
	State bool
}

type CfgDeviceAC struct {
	ID    string
	Name  string
	State bool
	Speed int
	Mode  int
	Temp  int
}

func New(filename string) (t *Cfg, err error) {
	t = &Cfg{filename: filename}
	if !t.exists() {
		if err := t.Save(); err != nil {
			return t, err
		}
	} else {
		if err := t.Load(); err != nil {
			return t, err
		}
	}
	return t, nil
}

func (t *Cfg) exists() bool {
	_, err := os.Stat(t.filename)
	return err == nil
}

func (t *Cfg) Save() (err error) {
	buf, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		return
	}
	err = ioutil.WriteFile(t.filename, buf, 0644)
	if err != nil {
		return
	}
	log.Printf("debug: [SAVE] %s", t.filename)
	return nil
}

func (t *Cfg) Load() (err error) {
	buf, err := ioutil.ReadFile(t.filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(buf, t)
	if err != nil {
		return
	}
	log.Printf("debug: [LOAD] %s", t.filename)
	return nil
}
