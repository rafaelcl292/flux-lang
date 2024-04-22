println 1 + 2 // prints 3
1 |> println // prints 1
a = 1 + 1

print_the_n3 -> println 3 // creates a function that prints 3

print_the_n3 // prints 3

print_sum_of_two_numbers -> a, b = println a + b


// multiline functions
does_a_bunch_of_stuff -> a, b, c {
    println a + b + c
    println a * b * c
    println a - b - c
    return "done"
}

// multiline assignment
my_var = 3 + 2
my_var {
    a = 1
    b = 2
    c = 3
    return a + b + c
}

// if statements
if 1 == 1 {
    println "1 is equal to 1"
} else if 1 == 2 {
    println "1 is equal to 2"
} else println "1 is not equal to 1 or 2"

// loops
for i = 0; i < 10; i++ {
    println i
}

// while loop
for {
    println "infinite loop"
    break
}

// types
a = 1
int a = 1
int addition -> int b, int c {
    return b + c
}
