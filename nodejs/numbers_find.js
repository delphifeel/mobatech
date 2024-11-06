// PREPARE
const prepSrc = (numbersCount) => {
  let numbers = [];
  for (let i = 0; i < numbersCount; i++) {
    numbers.push(i);
  }
  numbers.sort(() => Math.random() - 0.5);
  return { numbers };
};

const prepare = (srcArray, numbersCount) => {
  let numbers = [];
  for (let i = 0; i < numbersCount; i++) {
    numbers.push(srcArray[i]);
  }

  const toFind = [
    numbers[numbersCount / 2],
    numbers[numbersCount / 5],
    numbers[numbersCount / 10],
    numbers[numbersCount / 20],
  ];

  return { numbers, toFind };
};

import Benchmark from "benchmark";
var suite = new Benchmark.Suite();

const srcArray = prepSrc(99999);

const small = prepare(srcArray, 100);
const medium = prepare(srcArray, 1000);
const big = prepare(srcArray, 10000);

// add tests
suite
  .add("x1 find", function () {
    const { numbers, toFind } = small;
    let indexes = [];
    indexes.push(numbers.findIndex((n) => n === toFind));
    indexes.push(numbers.findIndex((n) => n === toFind));
    indexes.push(numbers.findIndex((n) => n === toFind));
    indexes.push(numbers.findIndex((n) => n === toFind));
  })
  .add("x1 for", function () {
    const { numbers, toFind } = small;
    let indexes = [];
    for (let i = 0; i < numbers.length; i++) {
      if (numbers[i] === toFind) {
        indexes.push(i);
        break;
      }
    }
    for (let i = 0; i < numbers.length; i++) {
      if (numbers[i] === toFind) {
        indexes.push(i);
        break;
      }
    }
    for (let i = 0; i < numbers.length; i++) {
      if (numbers[i] === toFind) {
        indexes.push(i);
        break;
      }
    }
    for (let i = 0; i < numbers.length; i++) {
      if (numbers[i] === toFind) {
        indexes.push(i);
        break;
      }
    }
  })
  .add("x10 find", function () {
    const { numbers, toFind } = medium;
    let indexes = [];
    indexes.push(numbers.findIndex((n) => n === toFind));
    indexes.push(numbers.findIndex((n) => n === toFind));
    indexes.push(numbers.findIndex((n) => n === toFind));
    indexes.push(numbers.findIndex((n) => n === toFind));
  })
  .add("x10 for", function () {
    const { numbers, toFind } = medium;
    let indexes = [];
    for (let i = 0; i < numbers.length; i++) {
      if (numbers[i] === toFind) {
        indexes.push(i);
        break;
      }
    }
    for (let i = 0; i < numbers.length; i++) {
      if (numbers[i] === toFind) {
        indexes.push(i);
        break;
      }
    }
    for (let i = 0; i < numbers.length; i++) {
      if (numbers[i] === toFind) {
        indexes.push(i);
        break;
      }
    }
    for (let i = 0; i < numbers.length; i++) {
      if (numbers[i] === toFind) {
        indexes.push(i);
        break;
      }
    }
  })
  .add("x100 find", function () {
    const { numbers, toFind } = big;
    let indexes = [];
    indexes.push(numbers.findIndex((n) => n === toFind));
    indexes.push(numbers.findIndex((n) => n === toFind));
    indexes.push(numbers.findIndex((n) => n === toFind));
    indexes.push(numbers.findIndex((n) => n === toFind));
  })
  .add("x100 for", function () {
    const { numbers, toFind } = big;
    let indexes = [];
    for (let i = 0; i < numbers.length; i++) {
      if (numbers[i] === toFind) {
        indexes.push(i);
        break;
      }
    }
    for (let i = 0; i < numbers.length; i++) {
      if (numbers[i] === toFind) {
        indexes.push(i);
        break;
      }
    }
    for (let i = 0; i < numbers.length; i++) {
      if (numbers[i] === toFind) {
        indexes.push(i);
        break;
      }
    }
    for (let i = 0; i < numbers.length; i++) {
      if (numbers[i] === toFind) {
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
