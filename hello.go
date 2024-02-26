package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
    subcommands := map[string]*flag.FlagSet{
        "help": flag.NewFlagSet("help", flag.ExitOnError),
        "kenteken": flag.NewFlagSet("kentekenChecker", flag.ExitOnError),
        "dns": flag.NewFlagSet("dns", flag.ExitOnError),
    }
    

    if len(os.Args) < 2 {
        fmt.Println("No command given, listing help")
        help("")
        return
    }

    switch os.Args[1] {
    case "help":
        subcommands["help"].Parse(os.Args[2:])
        help(subcommands["help"].Arg(0))
    case "kenteken":
        kentekenChecker()
    case "dns": 
        dnsReverse := subcommands["dns"].Bool("r", false, "Reverse lookup")
        subcommands["dns"].Parse(os.Args[2:])
        dns(*dnsReverse, subcommands["dns"].Arg(0))
    }
}

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

func help(Command string) {
    switch Command {
    case "kenteken":
        fmt.Println("Usage: hello kenteken")
        fmt.Println("Checks if a license plate is allowed to enter the parking lot")
    case "help":
        fmt.Println("Usage: hello help [command]")
        fmt.Println("Shows help for the given command")
    case "dns":
        fmt.Println("Usage: hello dns [-r] address")
        fmt.Println("Does a DNS lookup for the given address")
        fmt.Println("  -r: Reverse lookup")
    case "":
        fallthrough
    default:
        fmt.Println("Usage: hello [command]")
        fmt.Println("Available commands:")
        fmt.Println("  kenteken: license plate checker")
        fmt.Println("  help: show this help message")
        fmt.Println("  dns: DNS lookup")
    }
}

func kentekenChecker() {
    var kentekens = map[string]bool{
        "12-AB-34": true,
        "56-CD-78": true,
        "90-EF-12": true}
    var plate string
    fmt.Print("Kenteken: ")
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
}

func dns(reverse bool, address string){
    if address == "" {
        fmt.Println("No host/address given")
        return
    }
    var result []string
    var err error
    if reverse {
        result, err = net.LookupAddr(address)
    } else {
        result, err = net.LookupHost(address)
    }
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    for _, r := range result {
        fmt.Println(r)
    }
}