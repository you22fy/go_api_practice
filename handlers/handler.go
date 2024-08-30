package handlers // そのファイルがあるディレクトリ名と同じにする

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// helloのハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World!\n")
}

// /articleのハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "posting article...\n")
}

// /article/listのハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query() // URLからクエリパタメータを取得する.

	var page int // ページ番号を格納する変数
	// goのmapでは、キーが存在しない場合、値がゼロ値になり、okがfalseになる.
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}
	retString := fmt.Sprintf("Article List (page %d) \n", page)
	io.WriteString(w, retString)
}

// /article/1のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", articleId)
	io.WriteString(w, resString)
}

// /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

// /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
