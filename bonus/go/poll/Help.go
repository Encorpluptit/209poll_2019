package poll

import (
	"fmt"
	"os"
)

func PrintHelp(pgrm string, err *Error) {
	if err != &HelpError {
		fmt.Println(err.Message)
	}
	fmt.Println("Usage")
	fmt.Printf("\t%s pSize sSize p\n\n", pgrm)
	fmt.Println("DESCRIPTION:")
	fmt.Println("\tpSize\tsize of the population")
	fmt.Println("\tsSize\tsize of the sample (supposed to be representative)")
	fmt.Println("\tp\tpercentage of voting intentions for a specific candidate")
	os.Exit(err.Code)
}
