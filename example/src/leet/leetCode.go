package main

import "fmt"

func main() {
	var p = person{"小明", 20}
	p.run(12)
	person.run(p, 10)
	s := "abba1abba"
	longs := longestPalindrome(s)
	fmt.Println(longs)

}

func longestPalindrome(s string) string {

	if len(s) < 2 {
		return s
	}
	longest := s[:1]
	for i := 1; i < len(s); i++ {
		for rightStep := 0; rightStep < 2; rightStep++ {
			for p, q := i-1, i+rightStep; p >= 0 && q < len(s) && s[p] == s[q]; {
				if q-p+1 > len(longest) {
					longest = s[p : q+1]
				}
				p--
				q++
			}

		}
	}

	return longest

}

type person struct {
	name string
	age  int
}

func (p person) run(speed int) {
	fmt.Println(p.name, "年龄是", p.age, "运动速度是", speed)
}
