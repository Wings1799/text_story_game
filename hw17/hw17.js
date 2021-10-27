#!/usr/bin/node

let x = 5;
console.log(5 * x);

let mystring = "my string";
console.log(mystring + " yes this is mine " + x * 5);

//let newNum = Number(prompt("choose a number"));
//if (!Number.isNaN(newNum)) {
//    console.log("Your number is " + newNum);
//}

let mybool = false;
if (mybool == false) {
    console.log("Nope!")
} else {
    console.log("Yup!");
}

let mynum = 0;
while (mynum <= 8) {
    console.log(mynum)
    mynum++;
}

for (i = 0; i < 10; i++) {
    if (i % 2 == 1) {
        console.log(i);
    }
}

/*for (x = 0; x < 15; x++) {
    switch (x) {
        case 1:
            console.log("this is 1" + x);
        case 2:
            console.log("this is 2");
        case 3:
            console.log("this is 3");
        default:
            console.log(x);
    }
}*/