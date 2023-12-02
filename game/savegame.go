package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	c "github.com/AdamMcAdamson/blockeroni/config"
	s "github.com/AdamMcAdamson/blockeroni/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func saveGame() {
	timestamp := time.Now().Format("2006-01-02-150405")
	saveFolder := c.SavePath + c.SaveNameBase + timestamp
	saveDataFile := saveFolder + "/data.json"
	savePreviewImageFile := saveFolder + "/previewImage.png"

	createFolderIfMissing(c.SavePath)
	createFolderIfMissing(saveFolder)

	previewImage := *rl.LoadImageFromScreen() // Use pause image (will be able to save game from menu, aka game is paused. Should be using image to draw game background while paused).

	saveBoardState(saveDataFile)
	rl.ExportImage(previewImage, savePreviewImageFile)
}

func createFolderIfMissing(folderpath string) {
	if _, err := os.Stat(folderpath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(folderpath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func saveBoardState(filename string) {
	out, _ := json.Marshal(s.BoardState)
	if err := os.WriteFile(filename, out, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func getSaveFiles() bool {

	// LoadBoardState("./saves/blokarooni-save-2023-08-05-142833.json")
	// return
	files, err := os.ReadDir(c.SavePath)
	if err != nil {
		log.Fatal(err)
	}

	type option struct {
		index    int
		filename string
	}

	options := []option{}

	for i, dirEntry := range files {
		if dirEntry.IsDir() {
			fmt.Printf("%d: %s\n", i, dirEntry.Name())
			// @TODO: Remove
			// LoadBoardState(c.SaveFilePath + file.Name())
			// return
			options = append(options, option{index: i, filename: c.SavePath + dirEntry.Name()})
		}
	}

	if len(options) == 0 {
		fmt.Println("No game saves found.")
		return false
	}

	// @TODO: Handle with in-game input
	// @INFO: Have not figured out how to debug with this setup. Will work when in-game input is used.
	for {
		var input int
		_, err := fmt.Scanf("%d", &input)
		//fmt.Printf("%d\n", input)
		if err != nil {
			// fmt.Println("Invalid selection (ERRORED), please choose again")
			continue
		}
		for _, o := range options {
			if input == o.index {
				loadBoardState(o.filename + "/data.json")
				return true
			}
		}
		fmt.Println("Invalid selection, please choose again")
	}

}

func loadBoardState(filename string) {
	fmt.Printf("LoadingBoardState: %s\n", filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Could not open file: %s\n", err)
	}
	defer file.Close()

	// Need to get file size so we can unmarshal into 'data' successfully
	fileinfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, fileinfo.Size())

	if _, err := file.Read(data); err != nil {
		fmt.Printf("Could not read file: %s\n", err)
	}

	if err := json.Unmarshal(data, &s.BoardState); err != nil {
		fmt.Printf("%s\n", err)
	}

	c.DebugPrinted = false // @DebugRemove
}
