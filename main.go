package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main() {
	// fmt.Println("please enter the option")
	notes:=[]string{}
	reader:=bufio.NewReader(os.Stdin)


	for{
		fmt.Println("\nPlease choose an option:")
		fmt.Println("1. Add Note")
		fmt.Println("2. View Notes")
		fmt.Println("3. Delete Note")
		fmt.Println("4. Exit")
	
		input,_:=reader.ReadString('\n')
		input=strings.TrimSpace(input)
	

	switch input{
	case "1":
		fmt.Println("Enter your note")
		note,_:=reader.ReadString('\n')
		note=strings.TrimSpace(note)
		notes=append(notes,note)

	case "2":
		if len(notes)==0{
			fmt.Println("No notes are available")
		}else{
			fmt.Println("Your Notes")
			for i,Note:=range notes{
				fmt.Println(i,Note)
			}
		}

	case "3":
		if len(notes)==0{
			fmt.Println("No notes to be deleted")
		}

		fmt.Println("Enter the index of the note to be deleted")
		indexinput,_:=reader.ReadString('\n')
		indexinput=strings.TrimSpace(indexinput)
		intIndex,err:=strconv.Atoi(indexinput)   

		if err!=nil ||intIndex<0 ||intIndex>=len(notes){
			fmt.Println("Invalid index")
		}else{
			notes=append(notes[:intIndex],notes[:intIndex+1]...)
			fmt.Println("deleted.....")
		}
		

	case "4":
		fmt.Println("Exiting note app...")
		return

	default:
		fmt.Println("Invalid option...")	
	}

	}
}
