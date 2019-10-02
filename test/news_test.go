package test

import (
	"testing"

	"github.com/apc-unb/apc-api/components/news"
)

func TestNews(t *testing.T) {

	// Instantiate grades for test
	news1 := news.NewsCreate{
		Title:       "Teste 1",
		Description: "Testando o teste 1",
		Tags:        []string{"Teste 1", "Teste 2"},
	}

	if news1.Title != "Teste 1" {
		t.Errorf("Invalid news title, got: %s, want: %s.", news1.Title, "Teste 1")
	}

	if news1.Description != "Testando o teste 1" {
		t.Errorf("Invalid news Description, got: %s, want: %s.", news1.Description, "Testando o teste 1")
	}

	if len(news1.Tags) != 2 {
		t.Errorf("Invalid news tags size, got: %d, want: %d.", len(news1.Tags), 2)
	}

	if news1.Tags[0] != "Teste 1" {
		t.Errorf("Invalid news title[0] tag, got: %s, want: %s.", news1.Title, "Teste 1")
	}

}
