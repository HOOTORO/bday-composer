package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:           "Bday Composer",
		Width:           512,
		Height:          1244,
		AlwaysOnTop:     true,
		Frameless:       true,
		CSSDragProperty: "widows",
		CSSDragValue:    "1",
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 19, G: 17, B: 44, A: 17},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			&Contact{},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
