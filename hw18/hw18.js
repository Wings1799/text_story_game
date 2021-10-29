#!/usr/bin/node

const dubplusone = function (x) {
    return (x * 2) + 1;
}

function dubplustwo(x) {
    return (x * 2) + 2;
}

const dubplusthree = (x) => {
    return (x * 2) + 3;
}

//_____________________________________________________________________________________________

function min(x, y) {
    if (x > y) {
        return y;
    } else if (x < y) {
        return x;
    } else { return "Tie"; }
}

function isEven(x) {
    if (x == 0) { return true }
    else if (x == 1) { return false }
    else { return isEven(x - 2) }
}

function countBs(str) {
    let count = 0;
    console.log(str.length);
    for (i = 0; i < str.length; i++) {
        if (str[i] == "B") { count++; }
    }
    return count;
}

function add3(arr) {
    let tmparr = [];
    for (i = 0; i < arr.length; i++) {
        let val = (arr[i] + 3);
        tmparr.push(val);
    }
    return tmparr;
}

function echo(arr) {
    let tmparr = [];
    for (i = 0; i < arr.length; i++) {
        let val = arr[i];
        tmparr.push(val);
        tmparr.push(val);
    }
    return tmparr;
}

function add4(arr) {
    let tmparr = [];
    for (i = 0; i < arr.length; i++) {
        if (Array.isArray(i) == false) {
            let val = (arr[i] + 4);
            tmparr.push(val);
        } else if (Array.isArray(i)) {
            let tmp_arr = [];
            for (j = 0; j < i.length; j++) {
                let vall = (i[j] + 4);
                tmp_arr.push(vall);
            }
            console.log(tmp_arr);
            tmparr.push(tmp_arr);
        }
    }
    return tmparr;
}

//for some reason JS can't distinguish between an array and a number after looping through.

//console.log(dubplusone(6));
//console.log(dubplustwo(-2));
//console.log(dubplusthree(3));
//console.log(min(6, 3));
//console.log(isEven(2));
//console.log(countBs("helloBBBBBB"));
//console.log(add3([1, 2, 3, 4]));
//console.log(echo([1, 2, 3, 4]));
//console.log(add4([1, [2, 5], 3, 4]));