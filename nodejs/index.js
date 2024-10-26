import { doBench, expectSameArray, newRandomArr } from "./db.mjs";

// PREPARE

const COUNT = 1000;
let objectsArray = [];
let TO_FIND_IDS = [];
for (let i = 0; i < COUNT; i++) {
  const id = i + 1;
  TO_FIND_IDS.push(id);
  objectsArray.push({
    id,
    name: `${id} name`,
    price: (Math.random() + 1) * 33,
  });
}
objectsArray = newRandomArr(objectsArray);
// console.log(objectsArray);

TO_FIND_IDS = newRandomArr(TO_FIND_IDS);
TO_FIND_IDS = TO_FIND_IDS.slice(0, COUNT / 3);
// console.log(TO_FIND_IDS);

// const EXPECTED_ARRAY = [];
// for (let i = 0; i < TO_FIND_IDS.length; i++) {
//   const toFindId = TO_FIND_IDS[i];
//   EXPECTED_ARRAY.push(`${objectsArray[toFindId]}`);
// }

doBench("Test Array", (iterCount) => {
  let results = [];
  for (let iter_index = 0; iter_index < iterCount; iter_index++) {
    results = [];

    for (let tfi = 0; tfi < TO_FIND_IDS.length; tfi++) {
      for (let oa = 0; oa < objectsArray.length; oa++) {
        if (objectsArray[oa].id === TO_FIND_IDS[tfi]) {
          results.push(`${objectsArray[oa].name} + ${objectsArray[oa].price}`);
        }
      }
    }
  }

  // console.log(results);
});

doBench("Test Map", (iterCount) => {
  const map = {};
  for (let oa = 0; oa < objectsArray.length; oa++) {
    map[objectsArray[oa].id] = objectsArray[oa];
  }
  let results = [];
  for (let iter_index = 0; iter_index < iterCount; iter_index++) {
    results = [];

    for (let tfi = 0; tfi < TO_FIND_IDS.length; tfi++) {
      const v = map[TO_FIND_IDS[tfi]];
      results.push(`${v.name} + ${v.price}`);
    }
  }

  // console.log(results);
});

// doBench("Numbers 10000 for loop", (iterCount) => {
//   let indexes = [];
//   for (let iter_index = 0; iter_index < iterCount; iter_index++) {
//     for (let i = 0; i < numbers.length; i++) {
//       if (numbers[i] === 1) {
//         indexes.push(i);
//         break;
//       }
//     }
//     for (let i = 0; i < numbers.length; i++) {
//       if (numbers[i] === 2) {
//         indexes.push(i);
//         break;
//       }
//     }
//     for (let i = 0; i < numbers.length; i++) {
//       if (numbers[i] === 3) {
//         indexes.push(i);
//         break;
//       }
//     }
//     for (let i = 0; i < numbers.length; i++) {
//       if (numbers[i] === 4) {
//         indexes.push(i);
//         break;
//       }
//     }
//   }
// });
