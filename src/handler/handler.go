package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/phoihos/gosim/database"
	"github.com/phoihos/gosim/route"
)

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		io.WriteString(w, "Hello World")
	default:
		http.NotFound(w, r)
	}
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	db := database.GetConnection("example")
	if db == nil {
		io.WriteString(w, "No database connection exists")
		return
	}

	type product struct {
		Code  string
		Price uint
	}

	var results []product
	db.Raw("select * form products").Scan(&results)

	b, _ := json.Marshal(results)
	jsonText := string(b)
	io.WriteString(w, jsonText)
}

func init() {
	route.MapRouteFunc("/", handle).
		MapRouteFunc("/products", handleProducts)
}
