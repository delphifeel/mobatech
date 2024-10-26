import { doBench, expectSameArray, newRandomArr } from "../db.mjs";

/* DESCRIPTION:
 * There are [number: value] array.
 * Find some number related values.
 * Do it ITER_COUNT times
 */

// PREPARE ARRAY
const NUMBERS_SIZE = 10000;
const numbersKeys = new Uint32Array(NUMBERS_SIZE);
for (let i = 0; i < NUMBERS_SIZE; i++) {
  numbersKeys[i] = i;
}
numbersKeys.sort(() => Math.random() - 0.5);

let numbersValues = new Uint32Array(NUMBERS_SIZE);
for (let i = 0; i < numbersKeys.length; i++) {
  numbersValues[i] = numbersKeys[i] + 666;
}

// console.log(numbersKeys);
// console.log(numbersValues);

// PREPARE MAP
const numbersMap = {};
for (let i = 0; i < numbersKeys.length; i++) {
  numbersMap[numbersKeys[i]] = numbersKeys[i] + 666;
}

let numbersKeysToFind = newRandomArr(numbersKeys);
// numbersKeysToFind = numbersKeysToFind.slice(0, numbersKeys.length / 5);
numbersKeysToFind = numbersKeysToFind.slice(0, 5);

const EXPECT_ARRAY_VALUES = new Uint32Array(numbersKeysToFind.length);
for (let i = 0; i < numbersKeysToFind.length; i++) {
  EXPECT_ARRAY_VALUES[i] = numbersKeysToFind[i] + 666;
}

doBench("Find in array", (iterCount) => {
  let results = [];
  // for (let iter_index = 0; iter_index < iterCount; iter_index++) {
  results = [];
  for (let j = 0; j < numbersKeys.length; j++) {
    for (let i = 0; i < numbersKeysToFind.length; i++) {
      const keyToFind = numbersKeysToFind[i];
      if (numbersKeys[j] == keyToFind) {
        const v = numbersValues[j] + 666;
        results.push(v);
        break;
      }
    }
  }
  // }

  expectSameArray(results, EXPECT_ARRAY_VALUES, "Find In Array");
});

doBench("Find in map", (iterCount) => {
  let results = [];
  for (let iter_index = 0; iter_index < iterCount; iter_index++) {
    results = [];
    for (let i = 0; i < numbersKeysToFind.length; i++) {
      const keyToFind = numbersKeysToFind[i];
      const v = numbersMap[keyToFind];
      results.push(v);
    }
  }
  expectSameArray(results, EXPECT_ARRAY_VALUES, "Find In Map");
});
