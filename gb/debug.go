package gb

import(
	"fmt"
	"os"
)

var InstructionCounter = 0

func DebugPrint(msg string, instructionNo int, terminate bool){
	if(instructionNo == InstructionCounter){
		fmt.Println(msg)

		if(terminate){
			os.Exit(-1)
		}
	}
}
