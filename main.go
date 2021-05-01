package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"stats/CPU"
	"stats/Error"
	"stats/Memory"
	"strconv"
	"strings"

	"github.com/therecipe/qt/widgets"
)

type Sendable struct {
	CpuStats CPU.CPU       `json:"cpu_stats"`
	MemStat  Memory.Memory `json:"mem_stat"`
}

func main() {

	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(250, 200)
	window.SetWindowTitle("Connect to adress")

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Write something ...")
	widget.Layout().AddWidget(input)

	button := widgets.NewQPushButton2("Connect", nil)
	button.ConnectClicked(func(bool) {

		if err := ValidateAddress(input.Text()); err != nil {
			widgets.QMessageBox_Information(nil, "OK", "Invalid Address", widgets.QMessageBox__Ok,
				widgets.QMessageBox__Ok)
		}

		//TODO check response for specific errors
		resp, err := http.Get(input.Text())
		Error.HTTPError(err)

		if err := ValidateAnswer(resp); err != nil {
			widgets.QMessageBox_Information(nil, "OK", "Invalid Address", widgets.QMessageBox__Ok,
				widgets.QMessageBox__Ok)
		}

		widgets.QMessageBox_Information(nil, "OK", "connected to: "+input.Text(), widgets.QMessageBox__Ok,
			widgets.QMessageBox__Ok)
	})
	widget.Layout().AddWidget(button)

	window.Show()

	app.Exec()

}

func GetInfo() Sendable {

	cpuInfo := make(chan CPU.CPU)
	memInfo := make(chan Memory.Memory)

	go CPU.GetCPUStats(cpuInfo)
	go Memory.GetMemStats(memInfo)

	sendablePackage := Sendable{
		CpuStats: <-cpuInfo,
		MemStat:  <-memInfo,
	}

	return sendablePackage
}

func ValidateAddress(addr string) error {

	blocks := strings.Split(addr, ".")
	if len(blocks) != 4 {
		return errors.New("address validation: invalid address")
	}

	if n, err := strconv.Atoi(blocks[0]); n != 192 || err != nil {
		return errors.New("address validation: invalid address")
	}

	if n, err := strconv.Atoi(blocks[1]); n != 168 || err != nil {
		return errors.New("address validation: invalid address")
	}

	if n, err := strconv.Atoi(blocks[2]); n > 255 || err != nil {
		return errors.New("address validation: invalid address")
	}

	if n, err := strconv.Atoi(blocks[3]); n > 255 || err != nil {
		return errors.New("address validation: invalid address")
	}

	return nil
}

func ValidateAnswer(response *http.Response) error {

	if body, err := ioutil.ReadAll(response.Body); string(body) != "OK" || err != nil {
		return errors.New("connection error: bad response")
	}

	return nil
}
