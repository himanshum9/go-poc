package main

type Stack []int

func (s Stack) append(i int) {
	s.append(i)
}

func main() {
	var value Stack = []int{1, 2, 3, 4, 5, 6}
	for i := range value {
		Stack.append(i)
	}

}
