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
  const numbers = [];
  for (let i = 0; i < numbersCount; i++) {
    numbers.push(srcArray[i]);
  }

  let strings = [];
  for (let i = 0; i < numbersCount; i++) {
    strings.push(String(srcArray[i]));
  }

  const toFind = [
    numbers[numbersCount / 2],
    numbers[numbersCount / 5],
    numbers[numbersCount / 10],
    numbers[numbersCount / 20],
  ];
  const toFindStrings = [];
  for (let tf of toFind) {
    toFindStrings.push(String(tf));
  }

  return { numbers, strings, toFind, toFindStrings };
};

const srcArray = prepSrc(999_999);

const small = prepare(srcArray, 100);
const medium = prepare(srcArray, 1000);
const big = prepare(srcArray, 10000);
const bigest = prepare(srcArray, 90000);

// add tests
Deno.bench("100 numbers", function () {
  const { numbers, toFind } = small;
  const indexesNumbers = [];
  indexesNumbers.push(numbers.findIndex((n) => n === toFind[0]));
  indexesNumbers.push(numbers.findIndex((n) => n === toFind[1]));
  indexesNumbers.push(numbers.findIndex((n) => n === toFind[2]));
  indexesNumbers.push(numbers.findIndex((n) => n === toFind[3]));
  // console.log(indexesNumbers);
});

Deno.bench("100 strings", function () {
  const { strings, toFindStrings } = small;
  const indexesStrings = [];
  indexesStrings.push(strings.findIndex((n) => n === toFindStrings[0]));
  indexesStrings.push(strings.findIndex((n) => n === toFindStrings[1]));
  indexesStrings.push(strings.findIndex((n) => n === toFindStrings[2]));
  indexesStrings.push(strings.findIndex((n) => n === toFindStrings[3]));
  // console.log(indexesStrings);
});

Deno.bench("1000 numbers", function () {
  const { numbers, toFind } = medium;
  const indexesNumbers = [];
  indexesNumbers.push(numbers.findIndex((n) => n === toFind[0]));
  indexesNumbers.push(numbers.findIndex((n) => n === toFind[1]));
  indexesNumbers.push(numbers.findIndex((n) => n === toFind[2]));
  indexesNumbers.push(numbers.findIndex((n) => n === toFind[3]));
  // console.log(indexesNumbers);
});

Deno.bench("1000 strings", function () {
  const { strings, toFindStrings } = medium;
  const indexesStrings = [];
  indexesStrings.push(strings.findIndex((n) => n === toFindStrings[0]));
  indexesStrings.push(strings.findIndex((n) => n === toFindStrings[1]));
  indexesStrings.push(strings.findIndex((n) => n === toFindStrings[2]));
  indexesStrings.push(strings.findIndex((n) => n === toFindStrings[3]));
  // console.log(indexesStrings);
});

Deno.bench("10_000 numbers", function () {
  const { numbers, toFind } = big;
  const indexesNumbers = [];
  indexesNumbers.push(numbers.findIndex((n) => n === toFind[0]));
  indexesNumbers.push(numbers.findIndex((n) => n === toFind[1]));
  indexesNumbers.push(numbers.findIndex((n) => n === toFind[2]));
  indexesNumbers.push(numbers.findIndex((n) => n === toFind[3]));
  // console.log(indexesNumbers);
});

Deno.bench("10_000 strings", function () {
  const { strings, toFindStrings } = big;
  const indexesStrings = [];
  indexesStrings.push(strings.findIndex((n) => n === toFindStrings[0]));
  indexesStrings.push(strings.findIndex((n) => n === toFindStrings[1]));
  indexesStrings.push(strings.findIndex((n) => n === toFindStrings[2]));
  indexesStrings.push(strings.findIndex((n) => n === toFindStrings[3]));
  // console.log(indexesStrings);
});

Deno.bench("90_000 numbers", function () {
  const { numbers, toFind } = bigest;
  const indexesNumbers = [];
  indexesNumbers.push(numbers.findIndex((n) => n === toFind[0]));
  indexesNumbers.push(numbers.findIndex((n) => n === toFind[1]));
  indexesNumbers.push(numbers.findIndex((n) => n === toFind[2]));
  indexesNumbers.push(numbers.findIndex((n) => n === toFind[3]));
  // console.log(indexesNumbers);
});

Deno.bench("90_000 strings", function () {
  const { strings, toFindStrings } = bigest;
  const indexesStrings = [];
  indexesStrings.push(strings.findIndex((n) => n === toFindStrings[0]));
  indexesStrings.push(strings.findIndex((n) => n === toFindStrings[1]));
  indexesStrings.push(strings.findIndex((n) => n === toFindStrings[2]));
  indexesStrings.push(strings.findIndex((n) => n === toFindStrings[3]));
  // console.log(indexesStrings);
});
