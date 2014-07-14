// command line tool to escape quotes in  JSON string
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var verboseMode bool

func init() {

	flag.BoolVar(&verboseMode, "v", false, "Verbose output")
}

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

// read data from stdin and convert to escape quoted JSON
func main() {

	flag.Parse()

	inputJSON, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Fatalf("Error reading input: %v", err)
		return
	}

	if verboseMode {
		fmt.Println("Input JSON:")
		fmt.Println(string(inputJSON))
	}

	// unmarshal raw byte string to object
	var genericObject interface{}
	err = json.Unmarshal(inputJSON, &genericObject)

	if err != nil {
		log.Fatalf("Error JSON is not valid: %v", err)
	}

	// marshal back to byte string
	outputJSON, err := json.Marshal(genericObject)

	if err != nil {
		log.Fatalf("Error coverting object to JSON: %v", err)
	}

	if verboseMode {
		fmt.Println("Formatted JSON:")
		fmt.Println(string(outputJSON))
	}

	// escape the JSON
	escapedJSON := strconv.Quote(string(outputJSON))

	if verboseMode {
		fmt.Println("Escaped JSON:")
	}
	fmt.Println(escapedJSON)

}
