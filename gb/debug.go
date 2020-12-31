package gb

import(
	"fmt"
	"os"
)

var InstructionCounter = 0

func DebugPrint(msg string, instructionNo int, terminate bool){
	if(instructionNo == InstructionCounter){
		//fmt.Println(msg)
    fmt.Printf("[%d]: %s\n", instructionNo, msg)

		if(terminate){
      fmt.Fprintf(os.Stderr, "Exiting from Debug print\n")
			os.Exit(-1)
		}
	}
}
