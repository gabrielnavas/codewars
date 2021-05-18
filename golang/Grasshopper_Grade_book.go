package main

func GetGrade(a, b, c int) rune {
	score := (a + b + c) / 3
	if score <= 100 && score >= 90 {
		return 'A'
	}
	if score >= 80 && score < 90 {
		return 'B'
	}
	if score >= 70 && score < 80 {
		return 'C'
	}
	if score >= 60 && score < 70 {
		return 'D'
	}
	return 'F'
}

func main() {

}
