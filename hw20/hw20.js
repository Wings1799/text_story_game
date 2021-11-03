#!/usr/bin/node

function strChange(str, fun) {
  tmpstr = "";
  for (i = 0; i < str.length; i++) {
    tmpstr += fun(str[i]);
    console.log(tmpstr);
  }
  return tmpstr;
}

function charChange(char) {
  tmpchar = "";
  if (char.toUpperCase() == "E") {
    tmpchar = "X";
    return tmpchar;
  } else {
    return char;
  }
}

//console.log(strChange("Hello", charChange));
//_____________________________________________________________________________________
//Exercise 1
function flatten(array) {
  let newarr = [];
  for (let i of array) {
    newarr = newarr.concat(i);
  }
  return newarr;
}
//console.log(flatten([[1, 2], 3, 4, 5]));

//Exercise 2
function loop(val, test, update, body) {
  if (test(val) == true) {
    body(val);
    loop(update(val), test, update, body);
  }
}

/*loop(
  6,
  (n) => n > 0,
  (n) => n - 1,
  console.log
);*/

//Exercise 3
function every(array, test) {
  for (let i of array) {
    if (test(i) == false) {
      return false;
    }
  }
  return true;
}
/*console.log(every([1, 3, 5], (n) => n < 10));
console.log(every([2, 4, 16], (n) => n < 10));
console.log(every([], (n) => n < 10));*/

//Exercise 4
function dominantDirection(text) {
  /*IDK*/
}
//_________________________________________________________________________________________
function num_edges(graph) {
  let count = 0;
  for (let node of graph.adj_list) {
    count += Object.keys(node).length;
  }
  return count;
}

function has_edge(graph, v1, v2) {
  for (let key1 of Object.keys(graph.adj_list)) {
    if (key1 == v1) {
      for (let key2 of Object.keys(key1)) {
        if (key == v2) {
          return true;
        }
      }
    }
  }
  return false;
}

function print(graph) {
  console.log("directed" + graph.is_directed + "\n");
  for (node of graph.nodes) {
    console.log(node + " f" + node.value);
  }
  for (node of graph.adj_list) {
    for (nbr of node.value) {
      console.log(node + nbr + nbr.value);
    }
  }
}
