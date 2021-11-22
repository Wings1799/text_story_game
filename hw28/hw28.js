#!/usr/bin/node
function parseExpr(str) {
    let expr = "start\n  ";
    for (let i = 0; i < str.length; i++) {
        if (str[i] == 'c') {
            expr += ("change_color" + "\n  " + str[i+1] + "\n  "+ str[i+2] + "\n");
        } else if (str[i] == "f" | "l" | "b" | "r") {
            expr += ("movement" + "\n  " + str[i] + "\n" + str[i+1] + "\n");
        } else if (str[i] == "fill") {
            expr += (str[i] + "\n  " + str[i+1] + "\n");
        } else if (str[i] == "repeat") {
            expr += (str[i] + "\n  " + str[i+1] + "\n  " + str[i+2]);
        } else if (str[i] == "let") {
            expr += (str[i] + "\n  " + str[i+2]);
        }
    }
    return expr;
}

//let str = "c green blue fill 123";
//console.log(parseExpr(str));