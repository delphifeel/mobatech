// PREPARE
const prepare = (numbersCount) => {
  let numbers = [];
  for (let i = 0; i < numbersCount; i++) {
    numbers.push(i);
  }
  numbers.sort(() => Math.random() - 0.5);
  return { numbers };
};

const small = prepare(100);
const medium = prepare(1000);
const big = prepare(10000);

Deno.bench({ name: "x1 find" }, () => {
  const { numbers } = small;
  let indexes = [];
  indexes.push(numbers.findIndex((n) => n === 1));
  indexes.push(numbers.findIndex((n) => n === 2));
  indexes.push(numbers.findIndex((n) => n === 3));
  indexes.push(numbers.findIndex((n) => n === 4));
});

// add tests
Deno.bench("x1 for", function () {
  const { numbers } = small;
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
});
Deno.bench("x10 find", function () {
  const { numbers } = medium;
  let indexes = [];
  indexes.push(numbers.findIndex((n) => n === 1));
  indexes.push(numbers.findIndex((n) => n === 2));
  indexes.push(numbers.findIndex((n) => n === 3));
  indexes.push(numbers.findIndex((n) => n === 4));
});
Deno.bench("x10 for", function () {
  const { numbers } = medium;
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
});
Deno.bench("x100 find", function () {
  const { numbers } = big;
  let indexes = [];
  indexes.push(numbers.findIndex((n) => n === 1));
  indexes.push(numbers.findIndex((n) => n === 2));
  indexes.push(numbers.findIndex((n) => n === 3));
  indexes.push(numbers.findIndex((n) => n === 4));
});
Deno.bench("x100 for", function () {
  const { numbers } = big;
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
});
