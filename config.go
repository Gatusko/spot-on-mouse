package main

import (
	"encoding/json"
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/Gatusko/spot-on-mouse/constants"
)

// profile is one saved look of the spot.
type profile struct {
	Color  uint32 `json:"color"` // COLORREF, 0x00BBGGRR — not RGB
	Size   int32  `json:"size"`
	Circle bool   `json:"circle"`
}

// config is the whole file: the look restored at startup plus the four slots.
type config struct {
	Enabled  bool                            `json:"enabled"`
	Current  profile                         `json:"current"`
	Profiles [constants.ProfileCount]profile `json:"profiles"`
}

var cfg config

func defaultProfile() profile {
	return profile{
		Color:  constants.SpotDefaultColor,
		Size:   constants.SpotDefaultSize,
		Circle: true,
	}
}

func defaultConfig() config {
	c := config{Current: defaultProfile()}
	for i := range c.Profiles {
		c.Profiles[i] = defaultProfile()
	}
	return c
}

// configPath returns %APPDATA%\spot-on-mouse\config.json.
// https://pkg.go.dev/os#UserConfigDir
func configPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, constants.ConfigDirName, constants.ConfigFileName), nil
}

// loadConfig reads the config file, creating it with defaults on first run.
// Must run before any window is created: it only sets the package variables
// that createSpot and the controls read.
func loadConfig() {
	cfg = defaultConfig()

	path, err := configPath()
	if err != nil {
		log.Printf("config path: %v (using defaults)", err)
		applyProfileVars(cfg.Current)
		return
	}

	data, err := os.ReadFile(path)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		// First run: write the defaults so the file exists from now on.
		if errIn := saveConfig(); errIn != nil {
			log.Printf("config create: %v", errIn)
			return
		}
		log.Printf("config created: %s", path)
	case err != nil:
		log.Printf("config read: %v (using defaults)", err)
	default:
		// Unmarshal over the defaults: fields missing from the file keep them.
		if errIn := json.Unmarshal(data, &cfg); errIn != nil {
			log.Printf("config parse %s: %v (recreating with defaults)", path, errIn)
			cfg = defaultConfig()
			if errSv := saveConfig(); errSv != nil {
				log.Printf("config recreate: %v", errSv)
				return
			}
			return
		}
		log.Printf("config loaded: %s", path)
	}

	applyProfileVars(cfg.Current)
}

// saveConfig writes cfg to disk, overwriting any existing file.
func saveConfig() error {
	path, err := configPath()
	if err != nil {
		return err
	}
	// this is for creating the file
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// applyProfileVars copies a profile into the spot's package variables.
// The file is hand-editable, so the size is clamped to the slider's range.
func applyProfileVars(p profile) {
	spotColor = p.Color
	spotSize = max(constants.SpotMinSize, min(p.Size, constants.SpotMaxSize))
	spotCircle = p.Circle
}

// saveCurrent snapshots the spot's look into cfg.Current and writes the file.
func saveCurrent() {
	cfg.Current = profile{Color: spotColor, Size: spotSize, Circle: spotCircle}
	if err := saveConfig(); err != nil {
		log.Printf("config save: %v", err)
	}
}
