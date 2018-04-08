package main

import(
    "fmt"
)

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
    fmt.Println(pomatch("ab.c*|", "ab"))
}