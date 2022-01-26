package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/nomadcoders/nomadcoin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/pages/home.gohtml"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()} //블록체인 넣어주기
	tmpl.Execute(rw, data)
}
func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
