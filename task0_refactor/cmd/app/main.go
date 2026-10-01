// Command app — тонкий main.go, що лише "з'єднує" (wiring) пакунки
// calculator та textanalyzer. Уся бізнес-логіка живе в internal/.
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/calculator"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/textanalyzer"
)

func main() {
	sum := calculator.Add(2, 3)
	fmt.Println("2 + 3 =", sum)

	quotient, err := calculator.Divide(10, 0)
	if err != nil {
		if errors.Is(err, calculator.ErrDivisionByZero) {
			fmt.Println("division by zero was correctly detected:", err)
		} else {
			fmt.Fprintln(os.Stderr, "unexpected error:", err)
		}
	} else {
		fmt.Println("10 / 0 =", quotient)
	}

	words, err := textanalyzer.WordCount("the quick brown fox")
	if err != nil {
		fmt.Fprintln(os.Stderr, "unexpected error:", err)
		return
	}
	fmt.Println("word count:", words)
}
