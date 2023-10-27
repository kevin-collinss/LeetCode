package easy

func maximumWealth(accounts [][]int) int {
    final := 0
    for i := len(accounts)-1; i >=0; i-- {
        valA := 0
        for j:= len(accounts[i])-1; j >=0; j--{
            valA += accounts[i][j]
        }
        if valA > final {
            final = valA
        }
    }

    return final
}