Runtime
5ms
Beats 88.55% of users with Go

Memory
4.41MB
Beats 17.59% of users with Go

Notes

FOR LOOPS - **{ i := range nums } vs { i := 0; i < len(nums); i++ }**
When you use for i := range nums, Go efficiently iterates over the slice elements without having to calculate the length of the slice. Calculating the length of the slice (len(nums)) in every iteration can be costly, especially for large slices.