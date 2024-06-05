a = 1;

println a;
bom_dia = "oi";

does_a_bunch_of_stuff -> a, b, c {
    println a + b + c;
    println a * b * c;
    println a - b - c;
    return "done";
};

done = does_a_bunch_of_stuff(1, 2, 3);


my_var = 3 + 2;
my_var = {
    a = 1;
    b = 2;
    c = 3;
    if a == 1 {
        println "a is 1";
    } else {
        println "a is not 1";
    };
    return a + b + c;
};

if 1 == 1 {
    println "1 is equal to 1";
} else { if 1 == 2 {
    println "1 is equal to 2";
} else {
    println "1 is not equal to 1 or 2";
};};

inc -> a {
    return a + 1;
};

for i = 0; i < 10; i = inc(i) {
    println i;
};

for ;; {
    println "infinite loop";
    return;
};
