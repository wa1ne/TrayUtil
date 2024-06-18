package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/getlantern/systray"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Utils    []UtilsEntry `yaml:"utils"`
	Language string       `yaml:"language"`
}

type QuitButton struct {
	QuitButton string `yaml:"quit_button"`
}
type UtilsEntry struct {
	Path string `yaml:"path"`
	Name string `yaml:"name"`
}

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	// Set icon, title, and tooltip for system tray
	setApp()

	// Add menu items
	loadHud()
}

func loadConfig() Config {
	var config Config
	data, err := os.ReadFile("src/config.yaml")
	if err != nil {
		fmt.Println("Error reading config:", err)
		return Config{}
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error unmarshalling config:", err)
		return Config{}
	}
	return config
}

func loadHud() {
	config := loadConfig()
	for _, util := range config.Utils {
		addUtilityMenuItem(util.Name, util.Path)
	}
	systray.AddSeparator()

	menuQuit := systray.AddMenuItem(getQuitButtonText().QuitButton, getQuitButtonText().QuitButton)
	go func() {
		<-menuQuit.ClickedCh
		systray.Quit()
	}()
}

func addUtilityMenuItem(name string, path string) {
	menuItem := systray.AddMenuItem(name, name)
	go func() {
		for {
			<-menuItem.ClickedCh
			executeScript(path)
		}
	}()
}
func setApp() {
	systray.SetIcon(getIcon())
	systray.SetTitle("TrayUtil")
	systray.SetTooltip("TrayUtil by wa1ne")
}

func getIcon() []byte {
	data, err := os.ReadFile("src/icon.ico")
	if err != nil {
		fmt.Println("Error reading icon:", err)
		return nil
	}
	return data
}

func executeScript(scriptPath string) {
	ext := filepath.Ext(scriptPath)
	switch ext {
	case ".go":
		cmd := exec.Command("cmd", "/c", "start", "cmd", "/k", "go run", scriptPath)
		err := cmd.Start()
		if err != nil {
			fmt.Println("Error starting Golang script:", err)
		}
	case ".py":
		cmd := exec.Command("cmd", "/c", "start", "cmd", "/k", "python", scriptPath)
		err := cmd.Start()
		if err != nil {
			fmt.Println("Error starting Python script:", err)
		}
	default:
		err := exec.Command(scriptPath).Start()
		if err != nil {
			fmt.Println("Error starting command:", err)
		}
	}
	/*
		You can make working commands for your types of files by changing "your terminal command"
		to whatever you want to execute
			if ext == ".type" {
				cmd := exec.Command("cmd", "/c", "start", "cmd", "/k", "your terminal command",
				 scriptPath)
				err := cmd.Start()
					if err != nil {
				fmt.Println("Error starting Python script:", err)
			}
	*/
}

func getQuitButtonText() QuitButton {
	var quitbutton QuitButton
	data, err := os.ReadFile("src/lang/" + loadConfig().Language + ".yaml")
	if err != nil {
		fmt.Println("Error reading lang config:", err)
		return QuitButton{}
	}

	err = yaml.Unmarshal(data, &quitbutton)
	if err != nil {
		fmt.Println("Error unmarshalling lang config:", err)
		return QuitButton{}
	}

	return quitbutton
}

func onExit() {
	// Not needed
}
