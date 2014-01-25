package util

/***** GENERAL PURPOSE LINKED LIST NODE *****/
type Element struct {
	value interface{}
	next  *Element
}

func (e *Element) GetValue() interface{} {
	return e.value
}

func (e *Element) SetValue(i interface{}) {
	e.value = i
}

func (e *Element) SetNext(n *Element) {
	e.next = n
}

func (e *Element) Next() *Element {
	return e.next
}

/***** STACK *****/
type Stack struct {
	top  *Element
	size int
}

// Return the stack's length
func (s *Stack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

/***** UTILITY FUNCTIONS *****/
func Absi(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func StrW(s string, l int) string {
	if len(s) > l {
		s = s[:l-3]
		s = s + "..."
	} else {
		for len(s) != l {
			s = s + " "
		}
	}

	return s
}
