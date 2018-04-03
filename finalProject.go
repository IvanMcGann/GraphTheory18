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

// return type infix is returned from function intopost which converts infix to postfix regular expressions
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

//stores the state of possible links to other structs. symbol is a binary value of zero by rune.
type state struct {
    symbol rune
    edge1 *state
    edge2 *state
}

//keeps track of initial and accept state of Non-deterministic finite automata.
type nfa struct {
    initial *state
    accept *state 
}

//string input is postfix, return pointer to nfa structs 
func postfixRegexNFA(postfix string) *nfa {
	//a stack that is annaray of pointers to nfa that is empty
    nfastack := []*nfa{}

	//loop through rune ata time depending on the character(. concat, | union,* kleane star, default)
    for _, r := range postfix {

        switch r {
		case '.':
			//pop off last index of nfas stack frag2
			frag2 := nfastack[len(nfastack)-1]
			//remove up to end of the nfastack frag2
			nfastack = nfastack[:len(nfastack)-1]
			//pop off last index of nfas stack frag1
			frag1 := nfastack[len(nfastack)-1]
			//remove up to end of the nfastack frag1
            nfastack = nfastack[:len(nfastack)-1]

			//join the frag1 to frag2 initial state. 
            frag1.accept.edge1 = frag2.initial

			//and push to nfastack, new fragment is created using & to use address for pointer nfastack  
            nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})
            
		case '|':
			//pop off last index of nfas stack frag2
			frag2 := nfastack[len(nfastack)-1]
			//remove up to end of the nfastack frag2
			nfastack = nfastack[:len(nfastack)-1]
			//pop off last index of nfas stack frag1
            frag1 := nfastack[len(nfastack)-1]
			//remove up to end of the nfastack frag1
            nfastack = nfastack[:len(nfastack)-1]

			
			// new state edge 1 points to initial frag1 above, edge to to frag2 above
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			//new accept state
            accept := state{}
			//frag1 points to edge1 new accept state
			frag1.accept.edge1 = &accept
			//frag2 points to edge1 new accept state
            frag2.accept.edge1 = &accept            

			//push to nfastack, new initial/accept states pointed to nfastack pointer 
            nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
        
		case '*':
			//one fragment popped off stack for kleane star
            frag := nfastack[len(nfastack)-1]
            nfastack := nfastack[:len(nfastack)-1]
			
			//new accept state
			accept := state{}
			//new state edge1 points to initial fragment and edge 2 needs to point to accept state  
			initial := state{edge1: frag.initial, edge2: &accept}
			//join edge1 to initial state
			frag.accept.edge1 = frag.initial
			//join edge2 to accept state
            frag.accept.edge2 = &accept

			//push new fragment to nfastack
            nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		default: 
			//new accept state
			accept := state{}
			//new iniial state, set symbol to r and only edge points to accept state
            initial := state{symbol: r, edge1: &accept}
			//pushes to nfastack
            nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
        
        }//switch
	}// for	

	

		if len(nfastack) != 1 {
			fmt.Println("Uh oh...", len(nfastack), nfastack)
		}

	//returns value of nfastack which is nfa(just 1 item)	
    return nfastack[0]
}

//take a list of pointers to state, single pointer to state and the accept state. 
func addState(l []*state, s *state, a *state) []*state {
	//append to the list, the state that is passed in
	l = append(l, s)
	//if s is not an accept state and the state has its value as a zero, (0 is the zero value of rune not a character) 
	if s != a && s.symbol == 0 {
		//pass list l, another s state edge and accept state a to the list l
        l = addState(l, s.edge1, a)
		//if edge2 is not a empty/null(nil in go) value
		if s.edge2 != nil {
		//pass list l, another s state and accept state a to the list l
            l = addState(l, s.edge2, a)
        }
	}
	//returns a list of pointers to state
    return l
}

//takes arg postfixRegexNFA po and string s; returns a boolean type (isMatch) 
func pomatch(po string, s string) bool {	
	//set boolean false by default; indicates postfix regexp doesn't match string
	isMatch := false
	//new variable from postfixRegexNFA po is called upon
	ponfa := postfixRegexNFA(po)

	//array of pointers to current and next states. next is accessed by current  
	current := []*state{}
	next := []*state{}

	//passes current (by converting to a slice) to addState function so you can add the initial and accept states for postfix 
	current = addState(current[:], ponfa.initial, ponfa.accept)
	
	//loop through s a character at a time
	for _, r := range s {		
		//everytime read character loops through the current array
		for _, c := range current {
			//if c is same as symbol being read
			if c.symbol == r {
				//add c state and any other state to the next array
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		//set/swap current to next and set next to a blank array, to forget old states and read next character
		current, next = next, []*state{}
	}

	//loop through current array
	for _, c := range current {
		//if its equal to accept
		if c == ponfa.accept {
			// set boolean to true/accept the state
			isMatch = true
			break
		}
	}
	//return boolean expression
	return isMatch
}

/////////////////////////////////////////////////////////////////////////////////////////////////////

func main(){
	//header for UI
	fmt.Println("Graph Theory Project 2018: NFA's from a regular expression")
	//ask user to enter the choice
	fmt.Println("Choose Conversion \n 1. Infix Expressions Conversion to NFA \n 2. Postfix expression conversion \n 3. Exit project \n")
	//choice of user in the UI menu
	var option int
	//read option
	fmt.Scanln(&option)

	switch option {
	case 1:
		//
		fmt.Println("Option", option, "was entered. ")		
		//ask user to enter an infix 
		fmt.Print("Enter an infix expression: ")
		//using above method, read in the infix expression
		infixString, err := ReadFromInput()
		
		//error handling
		if err != nil {
			fmt.Println("Error when scanning input:", err.Error()) /*  */
			return
		}
		//display expression
		fmt.Println("infix", infixString)
		//change infix to postfix expression
		newPostFix := intopost(infixString)
		fmt.Println("postfix notation:", newPostFix)
		//asks user to enter a string to test if above matches it
		fmt.Print("Enter a string to test if it matches the nfa: ")
		testExp, err := ReadFromInput()
		//more error handling
		if err != nil {
			fmt.Println("Error when scanning input:", err.Error()) /*  */
			return
		}
		//display results
		fmt.Println("Does the string", testExp, " match ?", pomatch(newPostFix, testExp))
	case 2:
		fmt.Println("Next time, please enter choice 1 to run the program ;) ")
	default:
		fmt.Println("Next time, please enter a choice (1 or 2) if you want to run the program")
	}

}