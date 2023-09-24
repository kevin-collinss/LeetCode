Looking across solutions here, I can see a clever solution using a more algorithmic approach:

```go
func runningSum(nums []int) []int {
    for i := 1; i < len(nums); i++ {
        nums[i] += nums[i-1]
    }

    return nums
}

```
*need to learn to use more clever solutions instead of nested for loops*