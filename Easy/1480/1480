func runningSum(nums []int) []int {
    length := len(nums)
    output := make([]int, length)

    for i := 0; i < length; i++{
        for j := 0; j <= i; j++{
            output[i] = output[i] + nums[j]
        }
    }

    return output
}