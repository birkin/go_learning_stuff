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

    setup_saiyans()

    // copy/pointer experimentation
    print_as_copy()
    print_as_pointer()

}  // end func main()


/* copy/pointer experimentation */

type Saiyan struct {
    Name string
    Power int
}

func print_as_copy() {
    goku := Saiyan{ "Goku", 9000 }
    SuperCopy( goku )
    fmt.Printf( "copied power, ```%d```\n", goku.Power )
}

func print_as_pointer() {
    goku := &Saiyan{ "Goku", 9000 }
    SuperPointer( goku )
    fmt.Printf( "pointered power, ```%d```\n", goku.Power )
}

func SuperCopy(s Saiyan) {
    s.Power += 10000
}

func SuperPointer(s *Saiyan) {
    s.Power += 10000
    // s = &Saiyan{"Gohan", 1000}
    // fmt.Printf( "s, ```%s```\n", s )
}

/* --- */


/* add() experimentation */

func add( a int, b int ) int {
    return a + b
}

/* --- */


/* add_and_validate() experimentation */

func add_and_validate( a int, b int ) ( bool, int ) {
    return true, a + b
}

/* --- */


/* initial struct experimentation */

func setup_saiyans() {
    // define the struct
    type Saiyan struct {
        Name string
        Power int
    }
    // create an instance
    var goku Saiyan
    goku.Name = "Goku"
    goku.Power = 9000
    // create another instance
    birkin := Saiyan{
        Name: "Birkin",
        Power: 5000,
    }
    fmt.Printf( "goku power, ```%d```\n", goku.Power )
    fmt.Printf( "birkin power, ```%d```\n", birkin.Power )
}

/* --- */
