package main

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jhead/phantom/internal/proxy"
	"strconv"
)

func main() {
	a := app.New()
	w := a.NewWindow("MobilePhantom")

	var _ fyne.Theme = (*myTheme)(nil)
	a.Settings().SetTheme(&myTheme{})

	var proxyServer *proxy.ProxyServer
	var err error
	var loading = false

	hello := widget.NewLabel("MobilePhantom")
	status := widget.NewLabel("Waiting for input")
	status.Wrapping = fyne.TextWrapWord
	ipInput := widget.NewEntry()
	ipInput.SetPlaceHolder("Server ip")
	ipInput.Validator = inputValidator
	portInput := widget.NewEntry()
	portInput.Validator = numberInputValidator
	portInput.SetPlaceHolder("Server ip")
	portInput.Text = "19132"

	var disableInput = func() {
		ipInput.Disable()
		portInput.Disable()
	}

	var enableInput = func() {
		ipInput.Enable()
		portInput.Enable()
	}

	var disableButton = func() {}
	var enableButton = func() {}

	startButton := widget.NewButton("start/stop server", func() {
		status.SetText("Attempting to connect")
		disableInput()
		disableButton()
		if proxyServer != nil {
			proxyServer.Close()
			status.SetText("Server closed")
			proxyServer = nil
		} else if portInput.Validate() == nil && ipInput.Validate() == nil {
			proxyServer, err = proxy.New(proxy.ProxyPrefs{
				BindAddress:  "0.0.0.0",
				BindPort:     0,
				RemoteServer: fmt.Sprintf("%s:%s", ipInput.Text, portInput.Text),
				IdleTimeout:  60,
				EnableIPv6:   false,
				RemovePorts:  false,
				NumWorkers:   1,
			})
			if err != nil {
				status.SetText(fmt.Sprintf("Failed to init server: %s", err))
				proxyServer = nil
				loading = false
				enableInput()
				enableButton()
				return
			}

			go func() {
				if err := proxyServer.Start(); err != nil {
					status.SetText(fmt.Sprintf("Failed to start server: %s", err))
					proxyServer = nil
					enableInput()
				}
			}()
			status.SetText("Server started")
		} else {
			status.SetText("Invalid input")
		}
		enableInput()
		enableButton()
	})

	disableButton = func() {
		startButton.Disable()
	}

	enableButton = func() {
		startButton.Enable()
	}

	w.SetContent(container.NewVBox(
		container.NewCenter(
			hello,
		),
		ipInput,
		portInput,
		startButton,
		status,
	))
	w.Resize(fyne.NewSize(300, 400))
	w.ShowAndRun()
}

func inputValidator(s string) error {
	if len(s) == 0 {
		return errors.New("empty value")
	}
	return nil
}

func numberInputValidator(s string) error {
	if len(s) == 0 {
		return errors.New("empty value")
	}
	_, err := strconv.Atoi(s)
	return err
}
