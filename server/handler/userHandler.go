package handler

import "net/http"

// CreateUser creates user information.
// ユーザの名前情報をリクエストで受け取り、ユーザIDと認証用のトークンを生成しデータベースへ保存します。
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("Only allows POST method."))
		return
	}

	w.Write([]byte("OK"))
}

// GetUser returns user information.
// ユーザの認証と特定の処理はリクエストヘッダのx-tokenを読み取ってデータベースに照会をします。
func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("Only allows POST method."))
		return
	}

	w.Write([]byte("OK"))
}

// UpdateUser updates user information
// 初期実装では名前の更新を行います。
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("Only allows POST method."))
		return
	}

	w.Write([]byte("OK"))
}
