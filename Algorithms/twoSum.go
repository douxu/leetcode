package algorithms

func twoSum(nums []int, target int) []int {
	// dict := make(map[int]int)
	// for index, num := range nums {
	//     if index == 0 {
	//         dict[num] = index
	//         continue
	//     }else{
	//         if _,ok := dict[num];!ok {
	//             dict[num] = index
	//         }
	//     }
	//     other := target - num
	//     value, ok := dict[other]
	//     if ok && value != index{
	//         return []int{value,index}
	//     }
	// }
	// return []int{}

	dict := make(map[int]int)
	for index, value := range nums {
		other := target - value
		if _, ok := dict[other]; ok {
			return []int{index, dict[other]}
		} else {
			dict[value] = index
		}
	}
	return []int{}
}
