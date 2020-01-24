package handler

import (
	"fmt"
	"net/http"
)

// リクエストを処理する関数
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World from Go.")
}
