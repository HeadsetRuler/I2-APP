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
    var kentekens = map[string]bool{
        "12-AB-34": true,
        "56-CD-78": true,
        "90-EF-12": true}
    var plate string
    n, err := fmt.Scan(&plate)
    if n == 0 || err != nil {
        fmt.Println("Geen kenteken opgegeven")
        return
    }
    if kentekens[plate] {
        fmt.Println(groet())
    } else {
        fmt.Println("U heeft helaas geen toegang tot het parkeerterrein")
    }


    fmt.Println()
}