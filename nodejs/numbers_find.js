import { doBench, prepareBench } from "./db.mjs";

prepareBench(() => {
  return {
    input: [],
    toFind: [],
    expected: [],
  };
});

// PREPARE
const numbers = [];
for (let i = 0; i < 10000; i++) {
  numbers.push(i);
}
numbers.sort(() => Math.random() - 0.5);

doBench("Numbers 10000 find", (iterCount) => {
  let indexes = [];
  for (let iter_index = 0; iter_index < iterCount; iter_index++) {
    indexes.push(numbers.findIndex((n) => n === 1));
    indexes.push(numbers.findIndex((n) => n === 2));
    indexes.push(numbers.findIndex((n) => n === 3));
    indexes.push(numbers.findIndex((n) => n === 4));
  }
  return [];
});

doBench("Numbers 10000 for loop", (iterCount) => {
  let indexes = [];
  for (let iter_index = 0; iter_index < iterCount; iter_index++) {
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
  }

  return [];
});
