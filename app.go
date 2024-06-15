package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

type Contact struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
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
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) OpenFile() []Contact {
	selection, err := rt.OpenFileDialog(a.ctx, rt.OpenDialogOptions{
		Title: "Select File",
		Filters: []rt.FileFilter{
			{
				DisplayName: "Comma-separated values text file (*.csv)",
				Pattern:     "*.csv",
			},
		},
	})
	if err != nil {
		log.Println("Error selecting a file")
	}
	cont := readCsvFile(selection)
	return getBirthdayContacts(cont)
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func getBirthdayContacts(data [][]string) (bd []Contact) {
	for i, line := range data {
		if i > 0 { // omit header line
			var rec Contact
			for j, field := range line {
				if j == 0 {
					rec.Name = field
				} else if j == 15 {
					rec.Birthday = field
				}
			}
			if rec.Birthday != "" {
				bd = append(bd, rec)
			}
		}
	}
	return bd
}
