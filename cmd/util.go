package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func unmarshalFile(filePath string, d interface{}) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, d)
	if err != nil {
		return err
	}

	return err
}

func prettyPrintStruct(s interface{}) {
	o, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Println("Error pretty printing: " + err.Error())
	}
	fmt.Println(string(o))
}
