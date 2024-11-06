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

const srcArray = prepSrc(99999);

const small = prepare(srcArray, 100);
const medium = prepare(srcArray, 1000);
const big = prepare(srcArray, 10000);

// add tests
Deno.bench("x1 find", function () {
  const { numbers, toFind } = small;
  let indexes = [];
  indexes.push(numbers.findIndex((n) => n === toFind));
  indexes.push(numbers.findIndex((n) => n === toFind));
  indexes.push(numbers.findIndex((n) => n === toFind));
  indexes.push(numbers.findIndex((n) => n === toFind));
});
Deno.bench("x1 for", function () {
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
});
Deno.bench("x10 find", function () {
  const { numbers, toFind } = medium;
  let indexes = [];
  indexes.push(numbers.findIndex((n) => n === toFind));
  indexes.push(numbers.findIndex((n) => n === toFind));
  indexes.push(numbers.findIndex((n) => n === toFind));
  indexes.push(numbers.findIndex((n) => n === toFind));
});
Deno.bench("x10 for", function () {
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
});
Deno.bench("x100 find", function () {
  const { numbers, toFind } = big;
  let indexes = [];
  indexes.push(numbers.findIndex((n) => n === toFind));
  indexes.push(numbers.findIndex((n) => n === toFind));
  indexes.push(numbers.findIndex((n) => n === toFind));
  indexes.push(numbers.findIndex((n) => n === toFind));
});
Deno.bench("x100 for", function () {
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
});
