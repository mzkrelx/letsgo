package main

import (
	"log"

	"github.com/mzkrelx/letsgo/internal/server"
)

// サーバーを作成して起動する
func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
