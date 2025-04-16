package solver

import (
	"fmt"
	"sort"
)

func Solve(stackA []int) {
	switch len(stackA) {
	case 2:
		SortTwo(&stackA)
	case 3:
		SortThree(&stackA)
	default:
		SortStack(&stackA)
	}
}

func SortTwo(stackA *[]int) {
	a := *stackA
	if a[0] > a[1] {
		sa(a)
	}
}

func SortThree(stackA *[]int) {
	a := *stackA
	sorted := []int{}
	sorted=append(sorted, a...)
	
	sort.Ints(sorted)

	switch {
	case sorted[1] == a[0] && sorted[0] == a[1] && sorted[2] == a[2]:
		fmt.Println("sa")
	case sorted[2] == a[0] && sorted[1] == a[1] && sorted[0] == a[2]:
		fmt.Println("sa")
		fmt.Println("rra")
	case sorted[2] == a[0] && sorted[0] == a[1] && sorted[1] == a[2]:
		fmt.Println("ra")
	case sorted[0] == a[0] && sorted[2] == a[1] && sorted[1] == a[2]:
		fmt.Println("sa")
		fmt.Println("ra")
	case sorted[1] == a[0] && sorted[2] == a[1] && sorted[0] == a[2]:
		fmt.Println("rra")
	}

	a[0], a[1], a[2] = sorted[0],sorted[1], sorted[2]
}

func SortStack(stackA *[]int) {
	a := *stackA
	b := make([]int, 0, 100)
	
	
	l:=false
	for len(a) != 3 {

	sorted := []int{}
	sorted = append(sorted, *stackA...)
	sort.Ints(sorted)
	

		
		min, indexmin ,min2 := GetMin(a)
		half := len(a) / 2
		if half >= indexmin {
			for min != a[0] {
				if min == a[1] && a[0] == min2 {
					sa(a)
				
					for i := 2; i < len(a); i++ {
						if a[i]==sorted[i]{
							
							if i==len(a)-1 {
								l=true
								
							}
							continue
						}else{
							break
						}
						
					
					
					
					}
					
				} else {
					ra(a)
					for i := 2; i < len(a); i++ {
						if a[i]==sorted[i]{
							
							if i==len(a)-1 {
								l=true
								
							}
							continue
						}else{
							break
						}
						
					
					
					
					}
				}
			}
		} else {
			for min != a[0] {
				rra(a)
			
			}
			for i := 0; i < len(a); i++ {
				if a[i]==sorted[i]{
					
					if i==len(a)-1 {
						l=true
						
					}
					continue
				}else{
					break
				}
				
			
			
			
			}
		}
		if l{
			if len(b)!=0{
				for len(b) != 0 {
				pa(&a, &b)
			}
			fmt.Println("chwiya  tamara mais haniya")
			return
			}else{
				fmt.Println("mn wal mara adchi sahal")
				return
			}
			
			

		}
		pb(&a, &b)
	}

	SortThree(&a)

	for len(b) != 0 {
		pa(&a, &b)
	}
	fmt.Println("KAMLAT  9RAYA")
}

func GetMin(a []int) (int, int,int) {
	min := a[0]
	index := 0
	for i := 0; i < len(a); i++ {
		if min > a[i] {
			min = a[i]
			index = i
		}
	}
	min2:=a[0]
	newSlice := append([]int{}, a[:index]...)
	newSlice = append(newSlice, a[index+1:]...)
	for i := 0; i < len(newSlice); i++ {
		if min2 > newSlice[i] {
			min2= newSlice[i]
			
		}
	}
	return min ,index,min2
}
