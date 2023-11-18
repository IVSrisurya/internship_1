package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"

	"github.com/tebeka/selenium"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's showtime!", name)
}

func (a *App) DisableWindowsUpdates() string {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("powershell", "-Command", "Set-ItemProperty -Path 'HKLM:\\SOFTWARE\\Policies\\Microsoft\\Windows\\WindowsUpdate' -Name 'AU' -Value 1")
		err := cmd.Run()
		if err != nil {
			return fmt.Sprintf("Error disabling Windows Updates: %v", err)
		}
		return fmt.Sprintf("Windows Updates disabled successfully.")
	}
	return fmt.Sprintf("Windows Updates can only be disabled on Windows.")
}

func (a *App) DisableCommandPrompt() string {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("reg", "add", "HKCU\\Software\\Policies\\Microsoft\\Windows\\System", "/v", "DisableCMD", "/t", "REG_DWORD", "/d", "2", "/f")
		err := cmd.Run()
		if err != nil {
			return fmt.Sprintf("Error disabling Command Prompt: %v", err)

		}
		return fmt.Sprintf("Command Prompt disabled successfully.")
	} else {
		return fmt.Sprintf("Disabling Command Prompt is only supported on Windows.")
	}
}

func (a *App) DisableFileDownloads() string {
	caps := selenium.Capabilities{
		"browserName": "chrome", // or "firefox"
	}

	// Set up WebDriver
	wd, err := selenium.NewRemote(caps, "")
	if err != nil {
		return fmt.Sprintf("Failed to open browser: %v", err)

	}
	defer wd.Quit()

	// Navigate to a sample web page
	err = wd.Get("https://www.google.com")
	if err != nil {
		return fmt.Sprintf("Failed to navigate: %v", err)

	}

	// Execute JavaScript to disable file downloads
	script := `
		Object.defineProperty(window, 'saveAs', { value: undefined });
		Object.defineProperty(navigator, 'msSaveOrOpenBlob', { value: undefined });
		Object.defineProperty(HTMLAnchorElement.prototype, 'download', { value: '' });
	`
	_, err = wd.ExecuteScript(script, nil)
	if err != nil {
		return fmt.Sprintf("Failed to execute JavaScript: %v", err)

	}

	return fmt.Sprintf("File downloads in browsers disabled.")
}

func (a *App) BlockWebsite(website string) string {
	if runtime.GOOS == "windows" {
		err := blockWebsiteOnWindows(website)
		if err != nil {
			return fmt.Sprintf("Error blocking website: %v\n", err)

		}
	} else {
		return fmt.Sprintf("Blocking websites is only supported on Windows in this example.")
	}
	return fmt.Sprintf("Access to %s blocked.\n", website)
}

func blockWebsiteOnWindows(website string) error {
	// Assuming the hosts file is located at C:\Windows\System32\drivers\etc\hosts
	hostsFilePath := "C:\\Windows\\System32\\drivers\\etc\\hosts"
	file, err := os.OpenFile(hostsFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("failed to open hosts file: %v", err)
	}
	defer file.Close()

	// Add an entry to block the website
	entry := fmt.Sprintf("127.0.0.1 %s\n", website)
	_, err = file.WriteString(entry)
	if err != nil {
		return fmt.Errorf("failed to write to hosts file: %v", err)
	}

	return nil
}

const (
	// SPI_SETSCREENSAVETIMEOUT is a constant used for setting the screen saver timeout on Windows.
	SPI_SETSCREENSAVETIMEOUT = 15
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	systemParametersInfo = user32.NewProc("SystemParametersInfoW")
)

// SetLockScreenTimeout sets the lock screen timeout on Windows.
func (a *App) SetLockScreenTimeout(timeout time.Duration) string {
	// Convert the timeout to seconds
	timeoutInSeconds := uint32(timeout.Seconds())

	// Set the lock screen timeout using the Windows API
	result, _, err := systemParametersInfo.Call(
		uintptr(SPI_SETSCREENSAVETIMEOUT),
		0,
		uintptr(timeoutInSeconds),
		0,
	)

	if result == 0 {
		return fmt.Sprintf("Error setting lock screen timeout: %v\n", err)

	}

	return fmt.Sprintf("Lock screen timeout set to %s.\n", timeout.String())
}
func (a *App) GetContext() context.Context {
	return a.ctx
}

func (a *App) startup(ctx context.Context) {
	fmt.Println("Application is starting up...")
	a.ctx = ctx
}
