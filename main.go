package main
import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	//"music_service/service"
	//"net"
	"os"
    //"strings"
    // "github.com/Dolaantino/server_go/apiserver"
	// "github.com/BurntSushi/toml"
    "server_go/http-rest-api/internal/app/apiserver"
    "log"
    // "net/http"
    // "github.com/gorilla/mux"
)

func main() {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
				"gocloud", "gocloud", "playlist", "localhost", "5432")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
        fmt.Println(err)
		//fmt.Errorf("Db error: %s", err.Error())
        os.Exit(1)
	}
    flag.Parse()


	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil{
		log.Fatal(err)
	}
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
    //r := mux.NewRouter()
    // books = append(books, Book{ID: "1", Title: "Война и Мир", Author: &Author{Firstname: "Лев", Lastname: "Толстой"}})
    // books = append(books, Book{ID: "2", Title: "Преступление и наказание", Author: &Author{Firstname: "Фёдор", Lastname: "Достоевский"}})
    // r.HandleFunc("/books", getBooks).Methods("GET")
    // r.HandleFunc("/books/{id}", getBook).Methods("GET")
    // r.HandleFunc("/books", createBook).Methods("POST")
    // r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
    // r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
    // log.Fatal(http.ListenAndServe(":8000", r))


	// rows, err := db.Query("SELECT id, songname FROM playlist")
    // if err != nil {
    //     fmt.Println("Ошибка при выполнении запроса:", err)
    //     return
    // }
    // defer rows.Close()

    // // Обработка результатов
    // for rows.Next() {
    //     var id int
    //     var name string
    //     if err := rows.Scan(&id, &name); err != nil {
    //         fmt.Println("Ошибка при сканировании строки:", err)
    //         return
    //     }
    //     fmt.Printf("ID: %d, Name: %s\n", id, name)
    // }
    // if err := rows.Err(); err != nil {
    //     fmt.Println("Ошибка при обработке результатов:", err)
    // }
	// defer db.Close()
}