package main

import (
  "fmt"
)

func main() {

    first, last := "birkin", "diana"
    var middle string = "james"
    fmt.Printf( "person, ```%s %s %s```\n", first, middle, last )

    var sum int = add( 2, 3 )
    fmt.Printf( "sum, ```%d```\n", sum )

    var sum_validity bool = false
    var returned_sum int
    sum_validity, returned_sum = add_and_validate( 3, 4 )
    fmt.Printf( "validity, ```%t```; sum, ```%d```\n", sum_validity, returned_sum )

}

func add( a int, b int ) int {
    return a + b
}

func add_and_validate( a int, b int ) ( bool, int ) {
    return true, a + b
}


// func main() {
//   power := getPower()
//   fmt.Printf("It's over %d\n", power)
// }

// func getPower() int {
//   return 9001
// }
