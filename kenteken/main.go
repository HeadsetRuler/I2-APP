package main

import (
	"fmt"
	"os"
	"time"
)

func groet(now time.Time) string {
	const phrase = "! Welkom bij Fonteyn Vakantieparken"
	hour := now.Hour()
	switch {
	case hour < 7:
		return "Sorry, de parkeerplaats is ’s nachts gesloten"
	case hour < 12:
		return fmt.Sprintf("Goedemorgen%v", phrase)
	case hour < 18:
		return fmt.Sprintf("Goedemiddag%v", phrase)
	case hour < 23:
		return fmt.Sprintf("Goedenavond%v", phrase)
	default:
		return "Sorry, de parkeerplaats is ’s nachts gesloten"
	}
}

func main() {
	plate := os.Args[1]
	var kentekens = map[string]bool{
		"12-AB-34": true,
		"56-CD-78": true,
		"90-EF-12": true}
	now := time.Now()
	if len(plate) == 0 {
		fmt.Print("Kenteken: ")
		n, err := fmt.Scan(&plate)
		if n == 0 || err != nil {
			fmt.Println("Geen kenteken opgegeven")
			return
		}
	}
	if kentekenChecker(plate, kentekens) {
		fmt.Println(groet(now))
	} else {
		fmt.Println("U heeft helaas geen toegang tot het parkeerterrein")
	}
}

func kentekenChecker(plate string, kentekens map[string]bool) bool {
	return kentekens[plate]
}
