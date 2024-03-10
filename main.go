package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"example.com/go-postgres/db/models"

	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const (
	driverName = "postgres"
	dsn        = "postgres://postgres:postgres@localhost/app?sslmode=disable"
)

func main() {
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()
	ping(ctx, db)

	fmt.Println("\nSelect all books:")
	allBooks, err := models.Books().All(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	for _, book := range allBooks {
		log.Printf("Book: %+v\n", book)
	}

	fmt.Println("\nAdvanced query:")
	specificTitle := "Sample Book 1"
	books, err := models.Books(
		models.BookWhere.Title.EQ(specificTitle),
		// models.BookWhere.AuthorID.EQ(1),
		qm.Limit(10),
	).All(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	for _, book := range books {
		fmt.Println("Book:", book.Title)
	}

	fmt.Println("\nCount:")
	count, err := models.Books(models.BookWhere.Title.EQ(specificTitle)).Count(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Count:", count)

	fmt.Println("\nExists:")
	exists, err := models.Books(models.BookWhere.Title.EQ(specificTitle)).Exists(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Exists:", exists)

	fmt.Println("\nInsert:")
	newBookTitle := "New Book Title"
	newBook := &models.Book{
		Title:         newBookTitle,
		AuthorID:      1,
		PublisherID:   1,
		Isbn:          null.StringFrom("1234567890"),
		YearPublished: null.IntFrom(2023),
	}
	if err := newBook.Insert(ctx, db, boil.Infer()); err != nil {
		log.Fatal(err)
	}
	selectTitle(ctx, db, newBookTitle)

	fmt.Println("\nUpdate:")
	updatedBookTitle := "Updated Book Title"
	newBook.Title = updatedBookTitle
	if _, err := newBook.Update(ctx, db, boil.Infer()); err != nil {
		log.Fatal(err)
	}
	selectTitle(ctx, db, updatedBookTitle)

	fmt.Println("\nUpsert:")
	upsertedBookTitle := "Upserted Book Title"
	upsertBook := &models.Book{
		BookID:        newBook.BookID,
		Title:         upsertedBookTitle,
		AuthorID:      2,
		PublisherID:   2,
		Isbn:          null.StringFrom("0987654321"),
		YearPublished: null.IntFrom(2024),
	}
	if err := upsertBook.Upsert(ctx, db, true, []string{"book_id"}, boil.Infer(), boil.Infer()); err != nil {
		log.Fatal(err)
	}
	selectTitle(ctx, db, upsertedBookTitle)

	fmt.Println("\nDelete:")
	if _, err := newBook.Delete(ctx, db); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted.")

	fmt.Println("\nReload")
	if err := newBook.Reload(ctx, db); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Reload: not found.")
		} else {
			log.Fatal(err)
		}
	}

	// Eager Loading の例
	// ユーザーと関連する書籍を取得
	// user, err := models.FindUser(ctx, db, 1, qm.Load("Books"))
	// if err != nil {
	//     log.Fatal(err)
	// }
	// for _, book := range user.R.Books {
	//     fmt.Println("Book:", book.Title)
	// }

	// デバッグ出力の例
	// boil.DebugMode = true
	// books, _ = models.Books().All(ctx, db)
	// boil.DebugMode = false

	// Raw Query の例
	// _, err = queries.Raw("SELECT * FROM books WHERE title = 'New Book'").QueryAll(ctx, db)
	// if err != nil {
	//     log.Fatal(err)
	// }

	// Hook の例
	// func myBookHook(ctx context.Context, exec boil.ContextExecutor, book *models.Book) error {
	//     fmt.Println("Book Hook Triggered")
	//     return nil
	// }
	// models.AddBookHook(boil.BeforeInsertHook, myBookHook)

	// null パッケージの使用例
	// newBook.Isbn = null.StringFromPtr(nil) // ISBN を null に設定
}

// Ping the database to verify DSN provided by the user is valid and the
// server accessible. If the ping fails exit the program with an error.
func ping(ctx context.Context, db *sql.DB) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}

func selectTitle(ctx context.Context, db *sql.DB, title string) {
	books, err := models.Books(models.BookWhere.Title.EQ(title)).All(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	for _, book := range books {
		fmt.Println("Book:", book.Title)
	}
}
