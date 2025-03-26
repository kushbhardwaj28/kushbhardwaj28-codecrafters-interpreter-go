package main;

func Contains[T comparable](arr []T, item T) bool {
    for _, arrItem := range arr {
        if arrItem == item {
            return true;
        }
    }
    return false;
}
