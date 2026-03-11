package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

// Article представляет структуру данных для статьи
type Article struct {
	Title string
	Link  string
}

func main() {
	// Создаем новый коллектор
	c := colly.NewCollector(
		colly.AllowedDomains("benbyerslaw.com"),
	)

	articles := []Article{}

	// Селектор для поиска заголовков статей (основан на типичной структуре WP)
	// Сайт использует стандартные блоки ссылок для последних записей
	c.OnHTML("h2 a, .entry-title a", func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.Text)
		link := e.Attr("href")

		if title != "" && link != "" {
			articles = append(articles, Article{
				Title: title,
				Link:  link,
			})
		}
	})

	// Логирование процесса
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// Запуск скрапинга
	err := c.Visit("https://benbyerslaw.com/")
	if err != nil {
		log.Fatal(err)
	}

	// Вывод результата
	fmt.Printf("\nFound %d articles:\n", len(articles))
	fmt.Println(strings.Repeat("-", 40))
	for i, art := range articles {
		fmt.Printf("[%d] %s\n    URL: %s\n", i+1, art.Title, art.Link)
	}
}
