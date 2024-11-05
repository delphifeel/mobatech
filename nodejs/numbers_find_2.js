import { doBench, prepareBench } from "./db.mjs";

// PREPARE
let numbers = [];
for (let i = 0; i < 400; i++) {
  numbers.push(i);
}
numbers.sort(() => Math.random() - 0.5);

import Benchmark from "benchmark";
var suite = new Benchmark.Suite();

// add tests
suite
  .add("#1", function () {
    let indexes = [];
    indexes.push(numbers.findIndex((n) => n === 1));
    indexes.push(numbers.findIndex((n) => n === 2));
    indexes.push(numbers.findIndex((n) => n === 3));
    indexes.push(numbers.findIndex((n) => n === 4));
  })
  .add("#2", function () {
    let indexes = [];
    for (let i = 0; i < numbers.length; i++) {
      if (numbers[i] === 1) {
        indexes.push(i);
        break;
      }
    }
    for (let i = 0; i < numbers.length; i++) {
      if (numbers[i] === 2) {
        indexes.push(i);
        break;
      }
    }
    for (let i = 0; i < numbers.length; i++) {
      if (numbers[i] === 3) {
        indexes.push(i);
        break;
      }
    }
    for (let i = 0; i < numbers.length; i++) {
      if (numbers[i] === 4) {
        indexes.push(i);
        break;
      }
    }
  })
  // add listeners
  .on("cycle", function (event) {
    console.log(String(event.target));
  })
  .on("complete", function () {
    console.log("Fastest is " + this.filter("fastest").map("name"));
  })
  // run async
  .run({ async: true });
