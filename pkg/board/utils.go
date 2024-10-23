package board

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func sign(x int) int {
    if x < 0 {
        return -1
    }
    if x > 0 {
        return 1
    }
    return 0
}

func direction(isBlack bool) int {
    if isBlack {
        return -1
    }
    return 1
}