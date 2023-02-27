package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

// 投稿一覧を表す型
type Posts []string

// String は投稿一覧をHTMLに変換するメソッド
func (p Posts) String() string {
	sb := strings.Builder{}
	sb.WriteString("<ul>\n")

	for _, v := range p {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	return sb.String()
}

// 投稿一覧
var posts = Posts([]string{
	"Hello World",
	"Hello World2",
	"Hello World3",
})

// HTML テンプレート
var template2 = `
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<title>掲示板</title>
</head>
<body>
<h1>掲示板</h1>
<form action="" method="post" accept-charset="utf-8">
	<label class="label" for="name">コメント：</label>
	<input type="text" size="70" name="comment" value="">
	<p><input type="submit" value="投稿"></p>
</form>
コメント一覧:<br>
@@@POSTS@@@
</body>
</html>
`

// handler はHTTPリクエストを処理するハンドラ
func handler(w http.ResponseWriter, r *http.Request) {
	// POST ならフォームの値を取得
	if r.Method == "POST" {
		// フォームを解析
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(os.Stderr, "ParseForm() err: %v", err)
		}
		// フォームの値をログ出力
		for k, v := range r.Form {
			fmt.Fprintf(os.Stderr, "Form[%q] = %q\n", k, v)
		}
		// 投稿内容を投稿一覧へ追加
		if comment, ok := r.Form["comment"]; ok {
			// posts = append(posts, comment[0])
			posts = append(posts, template.HTMLEscapeString(comment[0]))
		}
	}
	// テンプレートのplaceholderを置換してHTMLを生成
	html := strings.Replace(template2, "@@@POSTS@@@", posts.String(), 1)
	// fmt.Println("test", template.HTMLEscapeString(r.Form.Get("comment")))
	fmt.Fprintf(w, html)
}

// main はプログラムのエントリポイント
func main() {
	// ハンドラの登録
	http.HandleFunc("/", handler)
	// 8080 ポートでHTTPのリスニング開始
	http.ListenAndServe(":8080", nil)
}
