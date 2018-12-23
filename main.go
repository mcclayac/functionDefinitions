package main

import (
	"fmt"
	"os"
)

//# Style 1
func printer(msg string, msg2 string) {
	fmt.Printf("#1 printer1 : %s\n", msg)
	fmt.Printf("#1 printer1 : %s\n", msg2)
}

//# Style 2
func printer2(msg, msg2 string) { // twp Pararmeters declared at once same type
	fmt.Printf("#2 printer2 : %s\n", msg)
	fmt.Printf("#2 printer2 : %s\n", msg2)
}

//# Style 3
func printer3(msg, msg2 string, repeat int) { // twp Pararmeters declared at once same type
	for repeat > 0 {
		fmt.Printf("#3 printer3 : %s\n", msg)
		fmt.Printf("#3 printer3 : %s\n", msg2)
		repeat--
	}
}

//# Style 4
func printer4(msg string, msg2 string) error {
	_, err := fmt.Printf("#4 printer1 : %s\n", msg)
	return err
}

//# Style 5
func printer5(msg string) (string, error) { // return multiple arguments
	msg += "\n"
	_, err := fmt.Printf("printer5 : %s", msg)
	return msg, err
}

// stype 6 : defer example
func printer6(msg string) error { // return multiple arguments
	defer fmt.Printf("\t#6 defer ##0\n")
	defer fmt.Printf("\t#6 defer ##1\n")
	defer fmt.Printf("\t#6 defer ##2\n")
	defer fmt.Printf("\t#6 defer ##3\n")
	_, err := fmt.Printf("\n\n#6 printer6 : %s\n", msg)
	return err
}

// stype 7 : defer example closing file
func printer7(msg string) error { // return multiple arguments
	f, err := os.Create("helloworld.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(msg)
	defer f.WriteString("#7  Defer 1 \n")
	defer f.WriteString("#7  Defer 2 \n")
	defer f.WriteString("#7  Defer 3 \n")
	defer f.WriteString("#7  Defer 4 \n")

	return err
}

// stype 8 : names return values
func printer8(msg string) (retuenedMessage string, e error) { // return multiple arguments
	_, e = fmt.Printf("#8 print 8 : %s\n", msg)
	retuenedMessage = "I Did it"
	return
}

// stype 9 : multiple names
func printer9(msgs ...string) { // return multiple arguments
	for _, msg := range msgs {
		fmt.Printf("\t\t\t#9 print 9 : %s\n", msg)
	}
}

// stype 10 : multiple names
func printer10(format string, msgs ...string) { // return multiple arguments
	for _, msg := range msgs {
		fmt.Printf(format, msg)
	}
}

func main() {

	printer("Hello World function!", "These Nuts")
	printer2("Tyanne", "Lisa")
	printer3("Susan", "Brenda", 5)
	err := printer4("Hello,", "World")
	err = err

	msg2, err := printer5("These Nutz!")
	if err == nil {
		fmt.Printf("#5 q = friendly %q\n", msg2)    // prints quotes
		fmt.Printf("#5 hex value %x\n", msg2)       // prints hexidecimal
		fmt.Printf("#5 spaced out hex % x\n", msg2) // prints hexidecimal spaced out
	}
	msg2 = msg2

	err2 := printer6("Hello World!")
	err2 = err2

	err3 := printer7("Hello World!")
	err3 = err3

	rtnMsg, err4 := printer8("Hello World!")
	err4 = err4
	rtnMsg = rtnMsg

	printer7("\t\t#7 These Nuts \n\n")

	printer9("These", "are", "The", "Times", "that", "try",
		"men's", "souls")

	printer10("#10 print 20 : %s\n", "These", "are", "The", "Times", "that", "try",
		"men's", "souls")

}
