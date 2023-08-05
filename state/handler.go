package gamestate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

func GetSaveFiles() {

	// LoadBoardState("./saves/blokarooni-save-2023-08-05-142833.json")
	// return
	files, err := ioutil.ReadDir(c.SaveFilePath)
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

	LoadBoardState(options[3].filename)
	return

	/*
		// @TODO: Handle with in-game input
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
					return
				}
			}
			fmt.Println("Invalid selection, please choose again")
		}
	*/
}

func LoadBoardState(filename string) {
	fmt.Printf("LoadingBoardState: %s\n", filename)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Could not open file: %s\n", err)
	}
	fmt.Println("File Opened")
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, fileinfo.Size())
	if _, err := file.Read(data); err != nil {
		fmt.Printf("Could not read file: %s\n", err)
	}
	// fmt.Println("File Read")
	if err := json.Unmarshal(data, &BoardState); err != nil {
		fmt.Printf("%s\n", err)
	}
	// fmt.Println("File Unmarshalled")
	//fmt.Printf("%v", BoardState)
}
