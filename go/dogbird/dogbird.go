package main

import "fmt"

type dog struct {
	name   string
	group  string
	height int
}

type bird struct {
	name   string
	group  string
	length int
}

func (d dog) describe() string {
	dscr := "わたしは" + d.group
	dscr += "、名は" + d.name
	return dscr
}

func (b bird) describe() string {
	dscr := "わたしは" + b.name
	dscr += "、" + b.group + "の仲間"
	return dscr
}

func (d dog) bigger(d2 dog) string {
	dscr := d.describe()
	dscr += "、私は" + d2.name

	diff := d.height - d2.height
	switch {
	case diff > 5:
		dscr += "より大きい"
	case diff < -5:
		dscr += "より小さい"
	default:
		dscr += "と同じくらいである"
	}
	return dscr
}

func main() {
	pome := dog{"ポメ", "ポメラニアン", 25}
	meji := bird{"メジロ", "スズメ", 12}
	fmt.Println(pome.describe())
	fmt.Println(meji.describe())
	maru := dog{"マル", "マルチーズ", 22}
	shiba := dog{"シバ", "芝犬", 40}
	fmt.Println(maru.bigger(shiba))
}
