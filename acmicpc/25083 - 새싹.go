package main

import (
	"bufio"
	"os"
)

func main() {
	wr := bufio.NewWriter(os.Stdout)
	_, _ = wr.WriteString("         ,r'\"7\nr`-_   ,'  ,/\n \\. \". L_r'\n   `~\\/\n      |\n      |")
	_ = wr.Flush()
}
