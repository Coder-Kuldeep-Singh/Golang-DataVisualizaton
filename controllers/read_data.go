package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func OpenFiles() (*os.File, error) {
	openFile, err := os.Open("jsondata.json")
	if err != nil {
		return nil, err
	}
	return openFile, nil
}

func ReadDataFromFile(file *os.File) ([]byte, error) {
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func UnmarshalData(bytes []byte) ([]DecodeData, error) {
	var jsondata []DecodeData
	err := json.Unmarshal(bytes, &jsondata)
	if err != nil {
		return jsondata, err

	}
	return jsondata, nil

}

func Decoded() []DecodeData {
	file, err := OpenFiles()
	if err != nil {
		log.Println("Error to Open File")
		panic(err)
	}
	bytes, err := ReadDataFromFile(file)
	if err != nil {
		log.Println("Error to read File data")
		panic(err)
	}
	unmarshal, err := UnmarshalData(bytes)
	if err != nil {
		log.Println("Error to Unmarshal Data", err)
	}
	// log.Println(unmarshal)
	return unmarshal
}
