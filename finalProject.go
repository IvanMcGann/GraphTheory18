package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

//https://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa

func ReadFromInput() (string, error) {

	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')

	return strings.TrimSpace(s), err
}

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

type state struct {
    symbol rune
    edge1 *state
    edge2 *state
}

type nfa struct {
    initial *state
    accept *state 
}

func postfixRegexNFA(postfix string) *nfa {
    nfastack := []*nfa{}

    for _, r := range postfix {

        switch r {
        case '.':
            frag2 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]
            frag1 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]

            frag1.accept.edge1 = frag2.initial

            nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})
            
        case '|':
            frag2 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]
            frag1 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]
            
            initial := state{edge1: frag1.initial, edge2: frag2.initial}
            accept := state{}
            frag1.accept.edge1 = &accept
            frag2.accept.edge1 = &accept            

            nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
        
        case '*':
            frag := nfastack[len(nfastack)-1]
            nfastack := nfastack[:len(nfastack)-1]

            accept := state{}
            initial := state{edge1: frag.initial, edge2: &accept}
            frag.accept.edge1 = frag.initial
            frag.accept.edge2 = &accept

            nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

        default: 
            accept := state{}
            initial := state{symbol: r, edge1: &accept}

            nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
        
        }//switch
	}// for	
		if len(nfastack) != 1 {
			fmt.Println("Uh oh...", len(nfastack), nfastack)
		}
    return nfastack[0]
}

func addState(l []*state, s *state, a *state) []*state {
    l = append(l, s)
    if s != a && s.symbol == 0 {
        l = addState(l, s.edge1, a)
        if s.edge2 != nil {
            l = addState(l, s.edge2, a)
        }
    }
    return l
}

func pomatch(po string, s string) bool {	
	isMatch := false
	ponfa := postfixRegexNFA(po)

	current := []*state{}
	next := []*state{}

	current = addState(current[:], ponfa.initial, ponfa.accept)

	for _, r := range s {		
		for _, c := range current {
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		current, next = next, []*state{}
	}

	for _, c := range current {
		if c == ponfa.accept {
			isMatch = true
			break
		}
	}
	return isMatch
}


func main(){
	fmt.Println("Graph Theory Project 2018: NFA's from a regular expression")
	
	fmt.Println("Choose Conversion \n 1. Infix Expressions Conversionto NFA \n 2. Postfix expression conversion \n 3. Exit project \n")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
	case 2:
	default: fmt.Println("please enter a choice (1 or 2) if you want to run the program")
	}

}