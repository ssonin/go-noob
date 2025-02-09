package main

import "fmt"

func main() {
	fmt.Println(WelcomeMessage("Monica"))
	fmt.Println(AddBorder("Whatever", 10))
	oldMsg := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`
	fmt.Println(CleanupMessage(oldMsg))

	fmt.Println(Welcome("Christiane"))
	fmt.Println(HappyBirthday("Frank", 58))
	fmt.Println(AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))
}
