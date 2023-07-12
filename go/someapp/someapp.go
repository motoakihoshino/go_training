package main

import "fmt"

type someapp struct {
	username string
	use      int
	isopen   bool
}

func (app *someapp) open() {
	if app.isopen {
		fmt.Printf("%sさんのアプリはすでに開いています\n\n", app.username)
	} else {
		app.use++
		if app.use > 2 {
			fmt.Printf("%sさんの試用期間は満了です。御購入検討ください\n\n", app.username)
		} else {
			fmt.Printf("こんにちは%sさん、%d回目のご使用になります\n", app.username, app.use)
			app.isopen = true
		}
	}
}

func (app *someapp) close() {
	if app.isopen {
		fmt.Printf("さようなら%sさん、アプリを終了しています。\n\n", app.username)
		app.isopen = false
	}
}

func newapp(username string) someapp {
	fmt.Printf("ようこそ%sさん、初めてのご使用になります\n", username)
	return someapp{username, 1, true}
}

func main() {
	fmt.Println("[ringoさんがアプリを使用開始]")
	rapp := newapp("ringo")
	rapp.close()

	fmt.Println("[ringoさんがアプリを再び開く]")
	rapp.open()
	rapp.close()

	fmt.Println("[mikanさんがアプリを使用開始]")
	mapp := newapp("mikan")
	mapp.close()

	fmt.Println("[mikanさんがアプリを再び開く]")
	mapp.open()

	fmt.Println("[ringoさんがアプリを三回目に開く]")
	rapp.open()

	fmt.Println("[mikanさんがアプリを三回目に開く]")
	mapp.open()

}
