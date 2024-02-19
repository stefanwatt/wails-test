package main

import (
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) GetBufLines() []string {
	// return some dummy lines representing the context of a file
	// the content should be the same as this file
	return []string{
		"// Path: app.go",
		"package main",
		"",
		"// App struct",
		"type App struct {",
		"	ctx context.Context",
		"}",
		"",
		"// NewApp creates a new App application struct",
		"func NewApp() *App {",
		"	return &App{}",
		"}",
		"",
		"// startup is called when the app starts. The context is saved",
		"// so we can call the runtime methods",
		"func (a *App) startup(ctx context.Context) {",
		"	a.ctx = ctx",
		"}",
		"",
		"// Greet returns a greeting for the given name",
	}
}
