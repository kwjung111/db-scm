package export

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func SaveFile(fileDir string, fileName string, data string) error {
	fileName = fileName + ".txt"

	err := ensureDir(fileDir)
	if err != nil {
		log.Println(err)
		return err
	}

	fullDir := fileDir + fileName

	file, err := os.Create(fullDir)
	if err != nil {
		//fmt.Errorf("creating %s file : %w", fileName, err)
		log.Println(err)
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(data)
	if err != nil {
		//fmt.Errorf("Error %s writing buffer:", fileName, err)
		log.Println(err)
		return err
	}

	err = writer.Flush()
	if err != nil {
		//fmt.Errorf("Error %s flushing buffer: ", fileName, err)
		log.Println(err)
		return err
	}

	log.Printf("file %s writed \n", fileName)
	return nil
}

func ensureDir(path string) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create dir %w", err)
	}
	return nil
}
