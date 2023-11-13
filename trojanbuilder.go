package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const asciiArt = `
MADE BY https://github.com/kikmanONTOP
████████╗    ██████╗      ██████╗          ██╗     █████╗     ███╗   ██╗    ██████╗     ██╗   ██╗    ██╗    ██╗         ██████╗     ███████╗    ██████╗ 
╚══██╔══╝    ██╔══██╗    ██╔═══██╗         ██║    ██╔══██╗    ████╗  ██║    ██╔══██╗    ██║   ██║    ██║    ██║         ██╔══██╗    ██╔════╝    ██╔══██╗
   ██║       ██████╔╝    ██║   ██║         ██║    ███████║    ██╔██╗ ██║    ██████╔╝    ██║   ██║    ██║    ██║         ██║  ██║    █████╗      ██████╔╝
   ██║       ██╔══██╗    ██║   ██║    ██   ██║    ██╔══██║    ██║╚██╗██║    ██╔══██╗    ██║   ██║    ██║    ██║         ██║  ██║    ██╔══╝      ██╔══██╗
   ██║       ██║  ██║    ╚██████╔╝    ╚█████╔╝    ██║  ██║    ██║ ╚████║    ██████╔╝    ╚██████╔╝    ██║    ███████╗    ██████╔╝    ███████╗    ██║  ██║
   ╚═╝       ╚═╝  ╚═╝     ╚═════╝      ╚════╝     ╚═╝  ╚═╝    ╚═╝  ╚═══╝    ╚═════╝      ╚═════╝     ╚═╝    ╚══════╝    ╚═════╝     ╚══════╝    ╚═╝  ╚═╝
                                                                                                                                                                                                                                                                 
																															
																							 
`

func main() {
	fmt.Println(asciiArt)
	fmt.Println("REMEMBER THAT YOU NEED CHANGE THE WEBHOOK URL IN THAT .TXT FILE")

	var inputFileName string
	fmt.Print("file with the code (.txt): ")
	fmt.Scanln(&inputFileName)
	code, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}

	outputFileName := strings.TrimSuffix(inputFileName, ".txt") + ".go"
	err = ioutil.WriteFile(outputFileName, code, 0644)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("go", "build", "-o", strings.TrimSuffix(outputFileName, ".go")+".exe", outputFileName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	os.Remove(outputFileName)
}
