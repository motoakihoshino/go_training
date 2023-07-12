package main

import "fmt"

type patient struct {
	name   string
	height int
	weight int
	chest  int
	upb    int
	hbp    int
	lbp    int
}

func newpatient(name string) patient {
	return patient{name, 0, 0, 0, 0, 0, 0}
}

func (pat patient) bodys(height int, weight int) patient {
	pat.height = height
	pat.weight = weight
	return pat
}

func (pat patient) bloodp(hbp int, lbp int) patient {
	pat.hbp = hbp
	pat.lbp = lbp
	return pat
}

func (pat patient) bmi() string {
	if pat.weight*pat.height <= 0 {
		return "測定エラー"
	}
	bmi := (pat.weight * 10000) / (pat.height * pat.height)

	if bmi < 19 {
		return "もっと食べましょう"
	} else if bmi > 25 {
		return "もっと運動しましょう"
	}
	return "バランスの取れた体格です"
}

func (pat patient) bprange() string {
	if pat.hbp*pat.lbp <= 0 {
		return "測定エラー"
	}
	if pat.hbp < 140 && pat.lbp < 90 {
		return "血圧は正常です"
	}
	return "血圧に注意が必要です"
}

func (pat patient) report() string {
	report := pat.name + "さん：\n"
	report += pat.bmi() + "\n"
	report += pat.bprange()
	return report
}

func main() {
	patients := []patient{
		newpatient("HM").bodys(170, 70).bloodp(120, 80),
		newpatient("SK").bodys(165, 88).bloodp(135, 92),
		newpatient("TN").bodys(156, 39).bloodp(122, 70),
		newpatient("AT").bodys(160, 55).bloodp(98, 66),
	}

	fmt.Println("***健康診断の結果です。***")

	for i := 0; i < len(patients); i++ {
		fmt.Println(patients[i].report())
		fmt.Println()
	}
	fmt.Println("***なお、胸囲・座高測定は省略します。***")
}
