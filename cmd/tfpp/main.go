package main

import (
	"bufio"
	"fmt"
	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/pasali/terraform-prettyplan/internal/assistant"
	"log"
	"os"
)

var (
	version string
	commit  string
	date    string
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}
	input := ""
	if fi.Mode()&os.ModeNamedPipe == 0 {
		if len(os.Args) == 1 {
			fmt.Printf("Usage: tfpp [file_path]\n%s (build date: %s commit: %s)", version, date, commit)
			os.Exit(1)
		}
		fileName := os.Args[1]
		data, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatalf("error while reading file, err: %v", err)
		}
		input = string(data)
	} else {
		var stdin []byte
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdin = append(stdin, scanner.Bytes()...)
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("error while reading from stdin, err: %v", err)
		}
		input = string(stdin)
	}

	apiKey := os.Getenv("TFPP_OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("TFPP_OPENAI_API_KEY is not set!!")
	}
	service := assistant.New(apiKey)
	resp, err := service.Assist(input)
	if err != nil {
		log.Fatalf("error while assisting, err: %v", err)
	}
	output := bufio.NewWriter(os.Stdout)
	result := markdown.Render(string(resp), 80, 0)
	output.Write(result)
	_ = output.Flush()
}
