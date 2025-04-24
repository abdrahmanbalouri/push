package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)



func sa(stack []int) {
	if len(stack) > 1 {
		stack[0], stack[1] = stack[1], stack[0]
	}
}
func ra(stack []int) {
	if len(stack)>1{
   for i := 0; i < len(stack)-1; i++ {
		stack[i], stack[i+1] = stack[i+1], stack[i]
	 }
	}
	
}

func rra(stack []int) {
	if len(stack) > 1 {
		for i := len(stack) - 1; i > 0; i-- {
			stack[i], stack[i-1] = stack[i-1], stack[i]
		}
	}
}

func pb(stackA *[]int, stackB *[]int) {
	if len(*stackA) > 0 {
		*stackB = append([]int{(*stackA)[0]}, *stackB...)
		*stackA = (*stackA)[1:]
	}
}

func pa(stackA *[]int, stackB *[]int) {
	if len(*stackB) > 0 {
		*stackA = append([]int{(*stackB)[0]}, *stackA...)
		*stackB = (*stackB)[1:]
	}
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	stackA := []int{}
	Map := make(map[int]bool)
	for _, num := range strings.Fields(os.Args[1]) {
		n, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal("malkom hagara")
		}
		if exist := Map[n]; exist {
			log.Fatal("malk mrid haka")
		}
		Map[n] = true
		stackA = append(stackA, n)
	}

	//fmt.Println(stackA)

	stackB := []int{}
	var line string
	for {
		_, err := fmt.Scanln(&line)
		if err != nil {
			break
		}

		switch line {
			case "ra":
			 ra(stackA)
		case "sa":
			sa(stackA)
		case "rra":
			rra(stackA)
		case "pb":
			pb(&stackA, &stackB)
		case "pa":
			pa(&stackA, &stackB)
		default:
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
	}

	cop:=[]int{}
	cop = append(cop, stackA...)
	sort.Ints(cop)
	//fmt.Println(cop)

	for i:=0;i<len(cop);i++{
		if cop[i] != stackA[i] {
			fmt.Println("ko")
			return
		}
	}
	if len(stackB) != 0 {
		fmt.Println("ko")
		return
	}
   fmt.Println("ok")

}
