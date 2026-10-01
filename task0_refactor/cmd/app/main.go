// Command app wires the CLI, service, and repository.
package main

import (
	"log"
	"os"

	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/handler"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/repository"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/service"
)

func main() {
	repo := &repository.ReportRepository{}
	svc := service.New(repo)
	cli := handler.New(svc, os.Stdout)
	if err := cli.Run(os.Args[1:]); err != nil {
		log.New(os.Stderr, "", 0).Print(err)
		os.Exit(1)
	}
}
