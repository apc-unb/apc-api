package test

import (
	"context"
	"log"
	"testing"

	"github.com/VerasThiago/api/components/news"
)

func TestNewsDB(t *testing.T) {

	//	Get conection with database
	//	Use config mongo function
	db, err := GetMongoDB("localhost", "27017")

	// Close conection in the end
	defer db.Disconnect(context.TODO())

	// Checks if creating conection with mongo db
	// doesn't return any errors
	if err != nil {
		log.Fatal(err)
	}

	// Get test collection of student
	collection := db.Database("apc_database_test").Collection("news_test")

	// Drop all content to start testing
	collection.Drop(context.TODO())

	// Instantiate grades for test
	news1 := news.NewsCreate{
		Title:       "Teste 1",
		Description: "Testando o teste 1",
		Tags:        []string{"Teste 1", "Teste 2"},
	}

	news2 := news.NewsCreate{
		Title:       "Teste 2",
		Description: "Testando o teste 2",
		Tags:        []string{"Teste 2", "Teste 3"},
	}

	news3 := news.NewsCreate{
		Title:       "Teste 3",
		Description: "Testando o teste 3",
		Tags:        []string{"Teste 3", "Teste 4", "Teste 3", "Teste 4"},
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 								 INSERT NEWS DB TEST 							     	 //
	///////////////////////////////////////////////////////////////////////////////////////////
	// Test if news class array can be inserted in test database
	// Checks if err variable is not null

	if err := news.CreateNews(db, []news.NewsCreate{news1, news2, news3}, "apc_database_test", "news_test"); err != nil {
		t.Errorf("Failed to insert news in Database : %s", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							   GET ALL NEWS FROM DB TEST 				      		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	var newsArray []news.News

	if newsArray, err = news.GetNews(db, "apc_database_test", "news_test"); err != nil {
		t.Errorf("Failed to get news from Database : %s", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							UPDATE LIST OF CLASSES FROM DB TEST   		         		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	newsArray[0].Title = "Teste 7"
	newsArray[0].Description = "Teste 7"

	if err := news.UpdateNews(db, []news.News{newsArray[0]}, "apc_database_test", "news_test"); err != nil {
		t.Errorf("Failed to update news in Database : %s", err)
	}

	newsArray = nil

	if newsArray, err = news.GetNews(db, "apc_database_test", "news_test"); err != nil {
		t.Errorf("Failed to get news from Database : %s", err)
	}

	if newsArray[0].Title != "Teste 7" {
		t.Errorf("Invalid newsArray[0] Title, got: %s, want: %s.", newsArray[0].Title, "Teste 7")
	}

	if newsArray[0].Description != "Teste 7" {
		t.Errorf("Invalid  newsArray[0] Description, got: %s, want: %s.", newsArray[0].Description, "Teste 7")
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// 							DELETE LIST OF CLASSES FROM DB TEST   		         		 //
	///////////////////////////////////////////////////////////////////////////////////////////

	if err := news.DeleteNews(db, []news.News{newsArray[0]}, "apc_database_test", "news_test"); err != nil {
		t.Errorf("Failed to delete news in Database : %s", err)
	}

	newsArray = nil

	if newsArray, err = news.GetNews(db, "apc_database_test", "news_test"); err != nil {
		t.Errorf("Failed to get news from Database : %s", err)
	}

	if len(newsArray) != 2 {
		t.Errorf("Invalid news size, got: %d, want: %d.", len(newsArray), 2)
	}

	if len(newsArray[1].Tags) != 4 {
		t.Errorf("Invalid news[1] Tagasd size, got: %d, want: %d.", len(newsArray[1].Tags), 4)
	}

}
