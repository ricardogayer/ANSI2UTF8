package main

import (
	"fmt"
	"log"
	"os"

	"github.com/djimenez/iconv-go"
)

func main() {

	fmt.Println("Conversão do arquivo de ANSI para UTF-8")

	b, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	output, _ := iconv.ConvertString(str, "windows-1252", "utf-8")

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(output)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")

}
