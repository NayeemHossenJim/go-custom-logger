package main
import "fmt"

const (
	Trace = iota
	Debug
	Info
	Warning
	Error
)

func main() {
	fmt.Println("Trace:", Trace)
	fmt.Println("Debug:", Debug)
	fmt.Println("Info:", Info)
	fmt.Println("Warning:", Warning)
	fmt.Println("Error:", Error)
}
