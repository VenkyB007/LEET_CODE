// leet - 169

func majorityElement(nums []int) int {
    numsMap := make(map[int]int)
    for _, i:= range nums{
        numsMap[i]++
        if numsMap[i]>len(nums)/2{
            return i
        }
    }
    return 0
}
