package main

import (
	"context"
	"embed"

	"github.com/x-foby/kakafka/internal/application"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := application.New()

	// Create application with options
	if err := wails.Run(&options.App{
		Title:            "kakafka",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	}); err != nil {
		runtime.LogFatal(context.Background(), err.Error())
	}
}
