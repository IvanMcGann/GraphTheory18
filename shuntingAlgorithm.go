package main


//Shunting algorithm based on video


import (
	"fmt"
)
    

func intopost(infix string) string {
	specials := map[rune]int{'*': 10, '.': 9,'|':8}
	pofix, s := []rune{}, []rune{}

	for _, r := range infix{
		switch {
		case r == '(':
			s = append(s, r)
		case r == ')':
		for s[len(s)-1] != '(' {
			pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
		}

		s = s[:len(s)-1]
		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s, r)
		default:
			pofix = append(pofix, r)

		}//switch
	}//for


		for len(s) > 0 {
			pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
		}

		return string(pofix)

}


func main()  {
	fmt.Println("Infix:  ", "a.b.c*")
	fmt.Println("Postfix:  ", intopost("a.b.c*"))

	fmt.Println("Infix:  ", "(a.(b|d))*")
	fmt.Println("Postfix:  ", intopost("(a.(b|d))*"))

	fmt.Println("Infix:  ", "a.(b|d).c*")
	fmt.Println("Postfix:  ", intopost("a.(b|d).c*"))

	fmt.Println("Infix:  ", "a.(b.b)+.c")
	fmt.Println("Postfix:  ", intopost("a.(b.b)+.c"))
}