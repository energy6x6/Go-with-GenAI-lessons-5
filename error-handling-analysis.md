# Lesson 05 — Error handling та рефакторинг

## Промпт для рев’ю

> Review this Go code for missing or incomplete error handling. List every location where an error is ignored, mishandled, or under-wrapped, and propose a fix for each using best practices like fmt.Errorf with %w.

Вхідний код наведено нижче як навчальний приклад «до». Це навмисно
створений поганий моноліт. Він зберігається лише в цьому звіті
та не входить до Go-пакетів проєкту.

## Моноліт до рефакторингу

```go
// Command before is an intentionally bad monolith for the homework AI review.
// Use cmd/app for the corrected application.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Report struct {
	Quotient   float64 `json:"quotient"`
	Words      int     `json:"words"`
	Characters int     `json:"characters"`
}

var errEmptyText = errors.New("empty text")

func divide(a, b float64) float64 {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func wordCount(text string) (int, error) {
	if strings.TrimSpace(text) == "" {
		return 0, fmt.Errorf("word count: %v", errEmptyText)
	}
	return len(strings.Fields(text)), nil
}

func save(path string, report Report) error {
	data, err := json.Marshal(report)
	_ = err
	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("save report: %v", err)
	}
	return nil
}

func main() {
	if len(os.Args) != 5 {
		panic("usage: app <a> <b> <text> <output.json>")
	}
	a, _ := strconv.ParseFloat(os.Args[1], 64)
	b, _ := strconv.ParseFloat(os.Args[2], 64)
	words, err := wordCount(os.Args[3])
	_ = err
	report := Report{
		Quotient: divide(a, b), Words: words,
		Characters: utf8.RuneCountInString(strings.TrimSpace(os.Args[3])),
	}
	err = save(os.Args[4], report)
	_ = err
	fmt.Fprintf(os.Stdout, "quotient=%g words=%d characters=%d\n", report.Quotient, report.Words, report.Characters)
}
```

## Відповідь AI та аналіз

| Місце у моноліті | Проблема | Застосоване виправлення |
| --- | --- | --- |
| `main`: перевірка кількості аргументів | `panic` для помилки користувача | `handler.Run` повертає `ErrUsage` |
| `main`: перший `ParseFloat` | Помилка відкинута через `_` | Перевірка та `fmt.Errorf("parse a: %w", err)` |
| `main`: другий `ParseFloat` | Помилка відкинута через `_` | Перевірка та `fmt.Errorf("parse b: %w", err)` |
| `divide`: нульовий дільник | `panic` для очікуваної помилки | `calculator.Divide` повертає wrapped `ErrDivisionByZero` |
| `wordCount`: порожній текст | `%v` втрачає sentinel у ланцюжку | `textanalyzer.WordCount` використовує `%w` |
| `main`: результат `wordCount` | `_ = err` дозволяє продовжити роботу | Service повертає помилку до запису файла |
| `save`: `json.Marshal` | `_ = err` приховує помилку кодування, наприклад для `NaN` | Repository перевіряє помилку перед записом та обгортає через `%w` |
| `save`: `os.WriteFile` | `%v` втрачає тип і причину файлової помилки; немає шляху | Repository додає шлях і `%w` |
| `main`: результат `save` | `_ = err` приховує невдале збереження | Помилка проходить через service і handler до main |
| `main`: `fmt.Fprintf` | Помилка виведення ігнорується | Handler повертає `print saved report: %w` |

Рекомендації застосовані у робочій версії `cmd/app` та `internal/`.
`recover` тут не потрібен: неправильні аргументи та відмови I/O
обробляються як звичайні помилки. Wrapping додає контекст операції,
а main повідомляє про помилку один раз і завершується з кодом 1.

Якщо виведення не вдалося після запису, JSON уже збережений — саме це
відображено в повідомленні `print saved report`. Запис замінює вміст
вказаного файла; атомарне збереження не входить у цей навчальний приклад.

## Приклад до і після

До:

```go
if err := os.WriteFile(path, data, 0600); err != nil {
    return fmt.Errorf("save report: %v", err)
}
```

Після:

```go
if err := os.WriteFile(path, data, 0600); err != nil {
    return fmt.Errorf("write report %q: %w", path, err)
}
```

Навіть після додаткових обгорток service та handler `errors.Is(err,
fs.ErrNotExist)` знаходить причину, а `errors.As` витягує `*os.PathError`.
Це перевіряє тест `TestRunPreservesStorageError`.

## Завдання агенту для рефакторингу

> Перетвори наданий монолітний main.go на Go-проєкт із cmd/ та internal/.
> Розподіли код між models, handler, service та repository. Забезпеч напрям
> handler → service → repository. Залиш у main.go лише створення залежностей,
> запуск і завершення програми. Використай пакети calculator та textanalyzer
> із завдання 0. Збережи результат для коректного вводу та виправ обробку помилок.

## Результат рефакторингу

```text
task0_refactor/
├── cmd/app/main.go            # wiring та запуск
└── internal/
    ├── models/report.go       # структура JSON-звіту
    ├── handler/cli.go         # аргументи CLI та виведення
    ├── service/report.go      # обчислення, аналіз, збереження
    ├── repository/report.go   # кодування JSON та файловий запис
    ├── calculator/            # арифметика
    └── textanalyzer/          # слова та Unicode-символи
```

Напрям залежностей: `handler → service → repository → models`.
Service також використовує `calculator`, `textanalyzer` і `models`.
Нижчі шари не імпортують вищі. Main створює repository, передає його service,
а service та stdout — handler. Бізнес-логіки в main немає.
