package main

import (
	"encoding/json"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	// https://pkg.go.dev/github.com/sqweek/dialog#section-readme
	"github.com/sqweek/dialog"
	"os"
	"path/filepath"
)

var gameSavePath = ""

func CheckSavingAndLoading(sandbox *Sandbox, undo *UndoRedoSystem) {
	var ctrlDown = rl.IsKeyDown(rl.KeyLeftControl) || rl.IsKeyDown(rl.KeyLeftControl)
	if ctrlDown && rl.IsKeyPressed(rl.KeyS) {
		var promptForPath = gameSavePath == "" || rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift)
		_, err := Save(promptForPath, sandbox, undo)
		if err != nil {
			dialog.Message("Something went wrong while saving: %s", err.Error()).Error()
			fmt.Println(err)
			return
		}
	} else if ctrlDown && rl.IsKeyPressed(rl.KeyO) {
		// Open
		if undo.dirty {
			var save = dialog.Message("%s", "Do you want to save your current game?").Title("Unsaved game").YesNo()
			if save {
				_, err := Save(gameSavePath == "", sandbox, undo)
				if err != nil {
					dialog.Message("Something went wrong while saving: %s", err.Error()).Error()
					fmt.Println(err)
					return
				}
			}
		}
		_, err := PromptAndLoad(sandbox, undo)
		if err != nil {
			switch e := err.(type) {
			case *UnknownSerialVersionError:
				dialog.Message("Failed to load saved game. Save file uses unknown version (version %d)", e.UnknownVersion).Error()
			case *UnknownIdMarshallError:
				dialog.Message("Failed to load saved game. Save file contains unknown %s id %d", e.Kind, e.Id).Error()
			case *UnknownTypeMarshallError:
				dialog.Message("Failed to load saved game. Save file contains unknown %s type %q", e.Kind, e.Name).Error()
			}
			fmt.Println(err)
			return
		}
	}
}

func AskAboutSaveBeforeExit(sandbox *Sandbox, undo *UndoRedoSystem) bool {
	if undo.dirty {
		var save = dialog.Message("%s", "Do you want to save your current game?").Title("Unsaved game").YesNo()
		if save {
			_, err := Save(gameSavePath == "", sandbox, undo)
			if err != nil {
				dialog.Message("Something went wrong while saving: %s", err.Error()).Error()
				fmt.Println(err)
				return false
			}
		}
	}
	return true
}

func Save(promptForPath bool, sandbox *Sandbox, undo *UndoRedoSystem) (bool, error) {
	if promptForPath {
		var file, err = dialog.File().Filter("Json files", "json").Save()
		if err != nil {
			return false, nil
		}
		if filepath.Ext(file) != ".json" {
			file += ".json"
		}
		gameSavePath = file
	}
	var data, err = json.MarshalIndent(sandbox, "", "    ")
	if err != nil {
		return false, err
	}
	err = os.WriteFile(gameSavePath, data, 0644)
	if err != nil {
		return false, err
	}
	undo.dirty = false
	fmt.Printf("Successfully saved game as %s\n", gameSavePath)
	return true, nil
}

func PromptAndLoad(sandbox *Sandbox, undo *UndoRedoSystem) (bool, error) {
	var file, err = dialog.File().Filter("Json file", "json").Load()
	if err != nil {
		return false, err
	}
	gameSavePath = file
	data, err := os.ReadFile(gameSavePath)
	if err != nil {
		return false, err
	}
	var newcopy = *sandbox // Copy
	err = json.Unmarshal(data, &newcopy)
	if err != nil {
		return false, err
	}
	*sandbox = newcopy // Commit
	*undo = NewUndoRedoSystem()
	fmt.Printf("Successfully loaded game from %s\n", gameSavePath)
	return true, nil
}
