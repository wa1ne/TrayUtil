# TrayUtil

![TUbanner](https://github.com/wa1ne/TrayUtil/assets/124814881/a0c10304-3355-4e89-b2ce-7023d1bc3312)

TrayUtil is a simple system tray application developed in Go that allows users to easily launch specified utilities or scripts directly from the system tray. The application reads a configuration file to dynamically create menu items in the system tray, which users can click to run their utilities.

![изображение](https://github.com/wa1ne/TrayUtils/assets/124814881/297e50a8-23ec-458e-98d4-5c0ee3607607)

## Installation

* Clone the Repository:
```bash
git clone https://github.com/wa1ne/TrayUtils.git
```

* Update Configuration:
Edit the src/config.yaml file to specify the utilities you want to add to the system tray menu and set the language for the quit button.

Example config.yaml:
```yaml
language: "eng" # Can be changed to "ru"
utils:
  - path: "path\\script.py"
    name: "Run script.py"
  - path: "path\\script.go"
    name: "Run script.go"
  - path: "path\\notepad.exe"
    name: "Run notepad"
```
* Build the Application:
Go to your TrayUtil directory in cmd
```bash
go build -ldflags -H=windowsgui
```
* Run the Application:
Simply run the precompiled TrayUtil.exe file included in the repository.
```bash
./TrayUtil.exe
```

## Configuration

### Config File (src/config.yaml)
The configuration file defines the utilities to be displayed in the system tray menu and the language for the quit button text.

### Adding New Utilities
To add new utilities, edit the src/config.yaml file and add entries under the utils section. Each utility entry requires a path and a name.

Example:
```yaml
- path: "path\\notepad.exe"
  name: "Run notepad"
```

### Language File (src/lang/eng.yaml or src/lang/ru.yaml)
To change the application language replace `language` field in src/config.yaml with your language code. By default its "eng", but also included "ru".
```yaml
language: "eng"
```
To add support for a new language:
* Create a new YAML file in the src/lang directory (e.g., src/lang/es.yaml).
* Define the quit_button text in the new file.
* Update the language field in src/config.yaml to the appropriate language code (e.g., language: "es").

This file contains the text for the quit button in the specified language. You can add more languages by creating new YAML files in the src/lang directory.

Example eng.yaml:

```yaml
quit_button: "Quit"
```

### How to add different formats
You can make working commands for your types of files by making case swithes in `executeScript` function
```go
case ".type":
  cmd := exec.Command("cmd", "/c", "start", "cmd", "/k", "your terminal command", scriptPath)
  err := cmd.Start()
  if err != nil {
    fmt.Println("Error starting your script:", err)
```

## Usage

Run the Application:
```bash
./TrayUtil.exe
```
Interact with the System Tray:
* The application icon will appear in the system tray.
* Click on the tray icon to see the menu items.
* Select any utility from the menu to run it.
* Click the "Quit" button to exit the application.
