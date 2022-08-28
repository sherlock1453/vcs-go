package internal

func Find(slice []string, str string) int{
    for i, n := range slice {
       if str == n{
            return i
        }
    }
    return -1
}

func Contains(slice []string, str string) bool {
    for _, n := range slice{
        if str == n {
            return true
        }
    }
    return false
}
