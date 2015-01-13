package main

import (
	"fmt"
	"encoding/json"
	"reflect"
	"os"
	"bufio"
	"io"
)


func main() {
	configFileText, err := readFile("json_map.json")
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("file text: ", configFileText)
	}
	var responseMap map[string]string
	responseMap = make(map[string]string)
	json.Unmarshal([]byte(configFileText), &responseMap)
	fmt.Println(responseMap)

}

func jsonTests() {
	var testMap1 map[string]string
	testMap1 = make(map[string]string)
	testMap1["hi"] = "hey!"
	testMap1["helloworld"] = "welcome!"
	fmt.Print("testMap1 created: ")
	fmt.Println(testMap1)
	var map1Json []uint8
	fmt.Println("creating json from testMap1...")
	map1Json, _ = json.Marshal(testMap1)
	fmt.Println("test type of map1Json: ", reflect.TypeOf(map1Json))
	fmt.Println("string version: ", string(map1Json))
	//fmt.Println("Creating map2Json from map1Json string...")
	fmt.Println("convert map2Json from json to map data structure:")
	var testMap2 map[string]string
	json.Unmarshal([]byte(map1Json), &testMap2)
	//fmt.Println("test type of map2Json: ", reflect.TypeOf(map2Json))
	fmt.Println(testMap2)
	var someBytes = [5]byte{65, 66, 67, 68, 69}
	fmt.Println(string(someBytes[:5]))
}

func readFile(path string) (FileText string, err error) {
    var (
		file *os.File
		// for some reason this gets converted to uint8
		// so I cast it back to a byte array later
		bufferArray [65536]byte
	)

    if file, err = os.Open(path); err != nil {
        return
    }

    defer file.Close()

    reader := bufio.NewReader(file)
	numBytes, err := reader.Read([]byte(bufferArray[:65536]))
    if err == io.EOF {
        err = nil
    }
	// see above comment on bufferArray
	FileText = string(bufferArray[:numBytes])
    return
}
