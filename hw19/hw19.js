#!/usr/bin/node

myarr = ["a", "b", "c"];
//console.log(myarr);
//console.log(myarr[2]);
//console.log(myarr.length);
//for (i = 0; i < myarr.length; i++) {
//  console.log(myarr[i].toUpperCase());
//}

let frog1 = {
  color: "green",
  age: 4,
  athlete: false,
};

//console.log(frog1.color);
delete frog1.athlete;
//console.log(frog1.athlete);
let frog2 = {};
Object.assign(frog2, frog1);
frog2.color = "blue";
//console.log(frog2);

//for (i of myarr) {
//  console.log(i);
//}
//___________________________________________________________
//Exercise 1
function range(start, end, step) {
  tmparr = [];
  for (i = start; i <= end; i++) {
    tmparr.push(i);
  }
  return tmparr;
}

function sum(nums) {
  tmpsum = 0;
  for (i = 0; i < nums.length; i++) {
    tmpsum += nums[i];
  }
  return tmpsum;
}

//Exercise 2
function reverseArray(arr) {
  tmparr = [];
  for (i = arr.length - 1; i >= 0; i--) {
    tmparr.push(arr[i]);
  }
  return tmparr;
}

function reverseArrayInPlace(arr) {
  for (i = 0; i < Math.floor(arr.length / 2); i++) {
    let val = arr[i];
    arr[i] = arr[arr.length - 1 - i];
    arr[arr.length - 1 - i] = val;
  }
  return arr;
}

//Exercise 3
function arrayToList(arr) {
  let list = {};
  for (i = arr.length - 1; i >= 0; i--) {
    list = { value: arr[i], rest: list };
  }
  return list;
}

function listToArray(list) {
  tmparr = [];
  for (let node = list; node; node = node.rest) {
    tmparr.push(node.value);
  }
  return tmparr;
}

function prepend(element, list) {
  let tmplist = { value: element, rest: list };
  return tmplist;
}

function nth(list, num) {
  tmparr = listToArray(list);
  for (i = 0; i < tmparr.length; i++) {
    if (i == num) {
      return tmparr[i];
    }
  }
  return undefined;
}

//Exercise 4
function deepEqual(a, b) {
  if (a === b) {
    return true;
  } else if (
    typeof a == "object" &&
    a != null &&
    typeof b == "object" &&
    b != null
  ) {
    if (Object.keys(a).length == Object.keys(b).length) {
      return true;
    } else {
      return false;
    }
  } else {
    return false;
  }
}
