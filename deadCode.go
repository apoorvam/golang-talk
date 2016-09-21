package main

// START OMIT
func Test() bool {
    return false
}

func Expensive() {
    if Test() {
        // unreachable code
    }
}
// STOP OMIT
