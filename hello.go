package main

import "fmt"
import "time"

func groet() string {
    const phrase = "! Welkom bij Fonteyn Vakantieparken"
    hour := time.Now().Hour()
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
    fmt.Println(groet())
}