package main

type Stack []int

func (s Stack) append(i int) {
	s.append(i)
}

func main() {
	s := Stack{}
	var value Stack = []int{1, 2, 3, 4, 5, 6}
	for _, val := range value {
		s.append(val)
	}

}
