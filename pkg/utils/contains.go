package utils

func Contains[T comparable](s []T, e T) int {
    for i, a := range s {
        if a == e {
            return i
        }
    }
    return -1
}