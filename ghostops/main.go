package main

import (
	"fmt"
	"os"
	"ghostops/read"
	"ghostops/humanRead"
	"ghostops/writeJson"
)

func main(){

	inputTerminal := os.Args[1:]
	var inputPath string
	var command_O string
	var outputPath string
	if len(inputTerminal) > 1{
		command_O = inputTerminal[1]
	}
	if command_O == "-o"{
		if len(inputTerminal) <= 2{
			fmt.Fprintln(os.Stderr, "please give the output file path")
		}else{
			outputPath = inputTerminal[2]
		}
	} else if command_O != ""{
		fmt.Fprintln(os.Stderr, "we don't have this command")
	}
	if len(inputTerminal) <= 0{
		fmt.Fprintln(os.Stderr, "please give the input file path")
	}else{
		inputPath = inputTerminal[0]
		stats,err := read.ReadFile(inputPath)
		if err != nil{
			fmt.Fprintln(os.Stderr, "we have error in reading",err)
			return
		}
		// fmt.Println(stats)
		humanread.HumanRead(stats)
		errorWrite := writejson.WriteJson(stats,outputPath)
		if errorWrite != nil{
			fmt.Fprintln(os.Stderr, "we have error in writing in json file",errorWrite)
		}
		// ---------------------------- make readabe print func for terminal
		// ---------------------------- make json func for making the file
	}

	

}