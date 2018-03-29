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

// return type infix is returned from fucntion intopost which converts infix to postfix regular expressions
func intopost(infix string) string {

	//maps characters into integer numbers, can keep track of special characters
	specials := map[rune]int{'*': 10, '.': 9,'|':8}

	// pofix,s are an array of runes, a rune is a character as displayed on screen(utf8), s is a stack 
	pofix, s := []rune{}, []rune{}

	//loop over the infix and return index of read character(1, 0 , 3, etc..) _ is ignored and r is the character. Range converts each element to a rune.
	for _, r := range infix{
		switch {
		//
		case r == '(':
			//put at the end of the stack
			s = append(s, r)
		case r == ')':
			//pop off the stack until the end.
			//while last character on stack doesn't equal open bracket append onto the top of the stack. 
		for s[len(s)-1] != '(' {
			pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
		}

		//s[:len(s)-1] is everything up to the bottom of the stack 
		s = s[:len(s)-1]
		
		//if r is in specials array above. 
		case specials[r] > 0:
			//while something is on the stack and less than what is at the end of the stack, take element off top of stack and put it into the end of the stack 
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s, r)
			//r is not a special or bracket character eg a, b, c, x, y, z
		default:
			//takes character and puts at end of output.
			pofix = append(pofix, r)

		}//switch
	}//for

		//if anything is on the top of the stack append it to the output
		for len(s) > 0 {
			pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
		}

		//returns pofix in string format
		return string(pofix)

}
////////////////////////////////////////////////////////////////////////////////////////////

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

/////////////////////////////////////////////////////////////////////////////////////////////////////

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