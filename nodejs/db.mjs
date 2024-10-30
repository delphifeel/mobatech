import { performance } from "perf_hooks";

let ITERATIONS_COUNT = 1000;

let STATE = null;

export const prepareBench = (func, config = {}) => {
  const iterations = config.iterations;
  if (iterations) {
    ITERATIONS_COUNT = iterations;
  }
  STATE = func(ITERATIONS_COUNT);
  console.log(`ITERATIONS: ${ITERATIONS_COUNT}`);
};

export const doBench = (name, func) => {
  console.log(`[${name}] Started`);

  let start = performance.now();
  const expected = func(ITERATIONS_COUNT, STATE.input, STATE.toFind);
  let end = performance.now();
  const diff = (end - start).toFixed(2);
  console.log(`[${name}]  ${diff} ms`);

  _expectSameArray(expected, STATE.expected);
};

export const arrayOfRandomNumbers = (size, mul) => {
  const result = [];
  for (let i = 0; i < size; i++) {
    result.push(Math.random() * mul);
  }
  return result;
};

export const newRandomArr = (srcArr) => {
  const res = [...srcArr];
  res.sort(() => Math.random() - 0.5);
  return res;
};

const _expectSameArray = (actual, expected, prefix) => {
  if (actual.length !== expected.length) {
    console.error(
      `[${prefix} TEST] arr1 != arr2\n\nArr1: ${actual}\n\nArr2: ${expected}`
    );
    return;
  }

  for (let i = 0; i < actual.length; i++) {
    const actualJSON = JSON.stringify(actual);
    const expectedJSON = JSON.stringify(expected);
    if (actualJSON !== expectedJSON) {
      console.error(
        `[${prefix} TEST] actual != expected\n\nActual:   ${actualJSON}\n\nExpected: ${expectedJSON}`
      );
      return;
    }
  }
};
