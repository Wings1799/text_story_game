#!/usr/bin/node

/*let frog = {};
frog.hop = (dir) => {
  console.log(`Frog hopped ${dir}`);
};
frog.hop("North");
console.log(Object.getPrototypeOf(frog));

class Frog {
  constructor(name, age, color) {
    this.name = name;
    this.age = age;
    this.color = color;
  }
  hop(dir) {
    console.log(`${this.name} hopped towards the ${dir}`);
  }
  birthday() {
    this.age++;
  }
  toString() {
    console.log(`${this.age} year old ${this.color} frog named ${this.name}`);
  }
}

let frog1 = new Frog("Jim", 12, "blue");
console.log(frog1);
frog.hop("East");
frog1.birthday();
console.log(frog1);
String(frog1);*/

//_____________________________________________________________
//Exercise 1
class Vec {
  constructor(x, y) {
    this.x = x;
    this.y = y;
  }
  plus(vector) {
    return new Vec(this.x + vector.x, this.y + vector.y);
  }
  minus(vector) {
    return new Vec(this.x - vector.x, this.y - vector.y);
  }
  get length() {
    return Math.sqrt(Math.pow(this.x, 2) + Math.pow(this.y, 2));
  }
}
//console.log(new Vec(1, 2).plus(new Vec(2, 3)));
//console.log(new Vec(1, 2).minus(new Vec(2, 3)));
//console.log(new Vec(3, 4).length);

//Exercise 2
class Group {
  constructor() {
    this.group = [];
  }
  add(elem) {
    this.group.push(elem);
  }
  delete(elem) {
    let tmpgrp = [];
    for (let i of this.group) {
      if (i != elem) {
        tmpgrp.push(i);
      }
    }
    this.group = tmpgrp;
  }
  has(elem) {
    return this.group.includes(elem);
  }
  static from(iter) {
    let g1 = new Group();
    for (let i of iter) {
      g1.add(i);
    }
    return g1;
  }
}
/*let group = Group.from([10, 20]);
console.log(group.has(10));
console.log(group.has(30));
group.add(10);
group.delete(10);
console.log(group.has(10));*/

//Exercise 3
class GroupIterator {
  constructor() {}
}

//Exercise 4
function hasOwnProperty(obj, prop) {
  for (i of Object.keys(obj)) {
    if (i == prop) {
      return true;
    }
  }
  return false;
}
//let map = { one: true, two: true, hasOwnProperty: true };
//console.log(hasOwnProperty(map, "hasOwnProperty"));
