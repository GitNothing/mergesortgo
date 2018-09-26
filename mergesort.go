package main
import(
	"fmt"
)
func main(){
	arr := []int{4, 5, 2, 7, 20, 1, 18, 3, 8, 14, 9, 13, 16, 17, 11}
	mergeSort(arr)
}

func mergeSort(arr []int) {
	totalLen := len(arr)
	current := [][]int{}
	currentNext := [][]int{}
	for _, e := range arr {
		current = append(current, []int{e})
	}
	fmt.Println("Split: ", current)
	pointer := -1
	condition := true
	fmt.Println("MERGING")
	for ok := true; ok; ok = condition {

		pointer += 2
		if pointer < totalLen {
			//fmt.Println("merging ", (pointer - 1), " and ", pointer)
			left := current[pointer-1]
			right := current[pointer]
			currentNext = append(currentNext, merge(left, right))
			fmt.Println(currentNext)
		} else { //if there is a last one for none even division
			currentNext = append(currentNext, current[len(current)-1])
			fmt.Println(currentNext)
		}

		if totalLen < pointer+2 {
			//completed, set currentNext as current and reset currentNext, set new len, and reset pointer
			current = currentNext
			currentNext = [][]int{}
			totalLen = len(current)
			pointer = -1
			fmt.Println("NEXT MERGE")
		}

		//All finish
		if totalLen == 1 {
			fmt.Println("Original: ", arr)
			fmt.Println("Sorted:   ", current[0])
			condition = false
		}
	}
}

func merge(arr1 []int, arr2 []int) []int {
	current := []int{}
	p1 := 0
	p2 := 0
	condition := true
	//fmt.Println("checking ", arr1, arr2)
	for ok := true; ok; ok = condition {
		//if p1 is not finish
		if p1 < len(arr1) {
			if p2 == len(arr2) { //just add if p2 is finish
				current = append(current, arr1[p1])
				p1++
			} else {
				//otherwise add only if p1 value is less than p2 value
				if arr1[p1] < arr2[p2] {
					current = append(current, arr1[p1])
					p1++
				}
			}
		}
		if p2 < len(arr2) {
			if p1 == len(arr1) {
				current = append(current, arr2[p2])
				p2++
			} else {
				if arr1[p1] > arr2[p2] {
					current = append(current, arr2[p2])
					p2++
				}
			}
		}

		if p1 == len(arr1) && p2 == len(arr2) {
			condition = false
		}
	}
	return current
}