// for loop
for i = 0; i < 10; i = i + 1 {
    println i;
};

// while loop
a = 0;

for ; a < 10; {
    println a;
    a = a + 1;
};

// infinite loop
for ;; {
    println "infinite loop";
    // exit from a program with return if not in a function
    return;
};
