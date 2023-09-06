package kernel
 
import "terminal"
 
func Main() {
	terminal.Init()
	terminal.Color = terminal.MakeColor(terminal.White,terminal.Blue)
	terminal.Clear()
	terminal.Row = 11
	terminal.Column = 30
	terminal.Print("Hello, Kernel World!\n")
}