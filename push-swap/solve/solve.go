package solver

import (
	"fmt"
	"sort"
)



func SortTwo(a []int) {
	//fmt.Println("SortTwo")
	if a[0] > a[1] {
		sa(a)
	}
}



func SortThree(a []int) {
	

	if a[0] > a[1] && a[1] < a[2] && a[0] < a[2] {
		fmt.Println("sa")
		sa(a)
	} else if a[0] > a[1] && a[1] > a[2] {
		fmt.Println("sa")
		sa(a)
		fmt.Println("rra")
		rra(a)
	} else if a[0] > a[1] && a[1] < a[2] && a[0] > a[2] {
		fmt.Println("ra")
		ra(a)
	} else if a[0] < a[1] && a[1] > a[2] && a[0] < a[2] {
		fmt.Println("sa")
		sa(a)
		fmt.Println("ra")
		ra(a)
	} else if a[0] < a[1] && a[1] > a[2] && a[0] > a[2] {
		fmt.Println("rra")
		rra(a)
	}
}

func SortStack(stackA *[]int) {
	a := *stackA
	b := make([]int, 0, 100)
	
	
	l:=false
	for len(a) != 3 {

	sorted := []int{}
	sorted = append(sorted, a...)
	sort.Ints(sorted)
	

		
		min, indexmin ,min2 := min1min2(a)
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
			fmt.Println(a)

			return
			}else{
				fmt.Println("mn wal mara adchi sahal")
				return
			}
			
			

		}
		pb(&a, &b)
	}

	SortThree(a)

	for len(b) != 0 {
		pa(&a, &b)
	}
	fmt.Println("KAMLAT  9RAYA")
}

func min1min2(a []int) (int, int,int) {
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

