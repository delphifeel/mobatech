// PREPARE
const prepSrc = (numbersCount) => {
  let numbers = [];
  for (let i = 0; i < numbersCount; i++) {
    numbers.push(i);
  }
  numbers.sort(() => Math.random() - 0.5);
  return numbers;
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
const bigest = prepare(srcArray, 90000);

// add tests
Deno.bench("100 find", function () {
  const { numbers, toFind } = small;
  let indexes = [];
  indexes.push(numbers.findIndex((n) => n === toFind[0]));
  indexes.push(numbers.findIndex((n) => n === toFind[1]));
  indexes.push(numbers.findIndex((n) => n === toFind[2]));
  indexes.push(numbers.findIndex((n) => n === toFind[3]));
});
Deno.bench("100 for", function () {
  const { numbers, toFind } = small;
  let indexes = [];
  for (let i = 0; i < numbers.length; i++) {
    if (numbers[i] === toFind[0]) {
      indexes.push(i);
      break;
    }
  }
  for (let i = 0; i < numbers.length; i++) {
    if (numbers[i] === toFind[1]) {
      indexes.push(i);
      break;
    }
  }
  for (let i = 0; i < numbers.length; i++) {
    if (numbers[i] === toFind[2]) {
      indexes.push(i);
      break;
    }
  }
  for (let i = 0; i < numbers.length; i++) {
    if (numbers[i] === toFind[3]) {
      indexes.push(i);
      break;
    }
  }
});
Deno.bench("1000 find", function () {
  const { numbers, toFind } = medium;
  let indexes = [];
  indexes.push(numbers.findIndex((n) => n === toFind[0]));
  indexes.push(numbers.findIndex((n) => n === toFind[1]));
  indexes.push(numbers.findIndex((n) => n === toFind[2]));
  indexes.push(numbers.findIndex((n) => n === toFind[3]));
});
Deno.bench("1000 for", function () {
  const { numbers, toFind } = medium;
  let indexes = [];
  for (let i = 0; i < numbers.length; i++) {
    if (numbers[i] === toFind[0]) {
      indexes.push(i);
      break;
    }
  }
  for (let i = 0; i < numbers.length; i++) {
    if (numbers[i] === toFind[1]) {
      indexes.push(i);
      break;
    }
  }
  for (let i = 0; i < numbers.length; i++) {
    if (numbers[i] === toFind[2]) {
      indexes.push(i);
      break;
    }
  }
  for (let i = 0; i < numbers.length; i++) {
    if (numbers[i] === toFind[3]) {
      indexes.push(i);
      break;
    }
  }
});
Deno.bench("10_000 find", function () {
  const { numbers, toFind } = big;
  let indexes = [];
  indexes.push(numbers.findIndex((n) => n === toFind[0]));
  indexes.push(numbers.findIndex((n) => n === toFind[1]));
  indexes.push(numbers.findIndex((n) => n === toFind[2]));
  indexes.push(numbers.findIndex((n) => n === toFind[3]));
});
Deno.bench("10_000 for", function () {
  const { numbers, toFind } = big;
  let indexes = [];
  for (let i = 0; i < numbers.length; i++) {
    if (numbers[i] === toFind[0]) {
      indexes.push(i);
      break;
    }
  }
  for (let i = 0; i < numbers.length; i++) {
    if (numbers[i] === toFind[1]) {
      indexes.push(i);
      break;
    }
  }
  for (let i = 0; i < numbers.length; i++) {
    if (numbers[i] === toFind[2]) {
      indexes.push(i);
      break;
    }
  }
  for (let i = 0; i < numbers.length; i++) {
    if (numbers[i] === toFind[3]) {
      indexes.push(i);
      break;
    }
  }
});
Deno.bench("90_000 find", function () {
  const { numbers, toFind } = bigest;
  let indexes = [];
  indexes.push(numbers.findIndex((n) => n === toFind[0]));
  indexes.push(numbers.findIndex((n) => n === toFind[1]));
  indexes.push(numbers.findIndex((n) => n === toFind[2]));
  indexes.push(numbers.findIndex((n) => n === toFind[3]));
});
Deno.bench("90_000 for", function () {
  const { numbers, toFind } = bigest;
  let indexes = [];
  for (let i = 0; i < numbers.length; i++) {
    if (numbers[i] === toFind[0]) {
      indexes.push(i);
      break;
    }
  }
  for (let i = 0; i < numbers.length; i++) {
    if (numbers[i] === toFind[1]) {
      indexes.push(i);
      break;
    }
  }
  for (let i = 0; i < numbers.length; i++) {
    if (numbers[i] === toFind[2]) {
      indexes.push(i);
      break;
    }
  }
  for (let i = 0; i < numbers.length; i++) {
    if (numbers[i] === toFind[3]) {
      indexes.push(i);
      break;
    }
  }
});
