# Notes   

Runtime   
0ms   
Beats 100.00% of users with Go   

Memory   
2.03MB   
Beats 100.00% of users with Go     

Notes   
x := start ^ goal: The XOR operation (^) is used to find the differing bits between start and goal. The result x will have set bits (1) where the corresponding bits in start and goal are different.   

count := 0: Initialize a variable count to zero. This variable will be used to count the number of set bits in x.   

for ; x != 0; count++ { x &= x - 1 }: This loop continues until x becomes zero. In each iteration, it clears the least significant set bit in x using the expression x &= x - 1. This operation effectively counts the number of set bits in x.   

return count: Finally, the function returns the count, which represents the minimum number of bit flips needed to transform start into goal.   