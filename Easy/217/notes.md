# Notes   

Runtime   
76ms   
Beats 78.18% of users with Go   

Memory   
8.35MB   
Beats 91.28% of users with Go   

Notes   
Hashsets arent available in Go. Workaround by creating map with empty struct: type Set map[int]struct{}