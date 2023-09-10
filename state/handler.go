package gamestate

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	c "github.com/AdamMcAdamson/blockeroni/config"
)

func SaveBoardState() {
	if _, err := os.Stat(c.SaveFilePath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(c.SaveFilePath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	timestamp := time.Now().Format("2006-01-02-150405")
	filename := c.SaveFilePath + c.SaveFileNameBase + timestamp + ".json"
	out, _ := json.Marshal(BoardState)
	if err := os.WriteFile(filename, out, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func GetSaveFiles() bool {

	// LoadBoardState("./saves/blokarooni-save-2023-08-05-142833.json")
	// return
	files, err := os.ReadDir(c.SaveFilePath)
	if err != nil {
		log.Fatal(err)
	}

	type option struct {
		index    int
		filename string
	}

	options := []option{}

	for i, file := range files {
		if !file.IsDir() {
			fmt.Printf("%d: %s\n", i, file.Name())
			// @TODO: Remove
			// LoadBoardState(c.SaveFilePath + file.Name())
			// return
			options = append(options, option{index: i, filename: c.SaveFilePath + file.Name()})
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
				LoadBoardState(o.filename)
				return true
			}
		}
		fmt.Println("Invalid selection, please choose again")
	}

}

func LoadBoardState(filename string) {
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

	if err := json.Unmarshal(data, &BoardState); err != nil {
		fmt.Printf("%s\n", err)
	}

	c.DebugPrinted = false // @DebugRemove
}
