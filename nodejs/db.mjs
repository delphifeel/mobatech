import { performance } from "perf_hooks";

const ITERATIONS_COUNT = 1000;
console.log(`ITERATIONS: ${ITERATIONS_COUNT}`);

export const doBench = (name, func) => {
  console.log(`[${name}] Started`);

  let start = performance.now();
  func(ITERATIONS_COUNT);
  let end = performance.now();
  const diff = (end - start).toFixed(2);
  console.log(`[${name}]  ${diff} ms`);
};

export const expectSameArray = (arr1, arr2, prefix) => {
  if (arr1.length !== arr2.length) {
    console.error(`[${prefix} TEST] arr1 length != arr2 length`);
    return;
  }

  for (let i = 0; i < arr1.length; i++) {
    if (arr1[i] !== arr2[i]) {
      console.error(
        `[${prefix} TEST] arr1 != arr2\n\nArr1: ${arr1}\n\nArr2: ${arr2}`
      );
      return;
    }
  }
};

export const newRandomArr = (srcArr) => {
  const res = [...srcArr];
  res.sort(() => Math.random() - 0.5);
  return res;
};
