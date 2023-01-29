package main

import (
	"fmt"
	"github.com/ssleert/sterm"
)

func main() {
	sterm.Color256Bg(27)
	fmt.Print(" ")
	fmt.Print(sterm.Reset)
	fmt.Println()
	sterm.Color256Bg(63)
	fmt.Print("  ")
	fmt.Print(sterm.Reset)
	fmt.Println()
	sterm.Color256Bg(99)
	fmt.Print("   ")
	fmt.Print(sterm.Reset)
	fmt.Println()
	sterm.Color256Bg(135)
	fmt.Print("    ")
	fmt.Print(sterm.Reset)
	fmt.Println()
	sterm.Color256Bg(171)
	fmt.Print("     ")
	fmt.Print(sterm.Reset)
	fmt.Println()
	sterm.Color256Bg(207)
	fmt.Print("      ")
	fmt.Print(sterm.Reset)

	fmt.Println("\n------------")

	sterm.Color256Bg(199)
	fmt.Print("     ")
	fmt.Print(sterm.Reset)
	fmt.Println()
	sterm.Color256Bg(169)
	fmt.Print("     ")
	fmt.Print(sterm.Reset)
	fmt.Println()
	sterm.Color256Bg(139)
	fmt.Print("     ")
	fmt.Print(sterm.Reset)
	fmt.Println()
	sterm.Color256Bg(109)
	fmt.Print("     ")
	fmt.Print(sterm.Reset)
	fmt.Println()
	sterm.Color256Bg(79)
	fmt.Print("     ")
	fmt.Print(sterm.Reset)
	fmt.Println()
	sterm.Color256Bg(49)
	fmt.Print("     ")
	fmt.Print(sterm.Reset)
	fmt.Println()

}