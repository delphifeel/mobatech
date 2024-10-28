import { doBench, expectSameArray, newRandomArr } from "./db.mjs";

// PREPARE

const COUNT = 3000;
let INPUT_ARRAY = [];
let TO_FIND_IDS = [];
for (let i = 0; i < COUNT; i++) {
  const id = i + 1;
  TO_FIND_IDS.push(id);
  INPUT_ARRAY.push({
    id,
    name: `${id} name`,
    price: (Math.random() + 1) * 33,
  });
}

const inputRandomArr = newRandomArr(INPUT_ARRAY);

const findIn_INPUT_ARRAY = (id) => {
  for (let i = 0; i < INPUT_ARRAY.length; i++) {
    if (INPUT_ARRAY[i].id === id) {
      return INPUT_ARRAY[i];
    }
  }
};

const Array = {
  objectsArray: [],

  init(srcArray) {
    this.objectsArray = [...srcArray];
  },

  findAndFormat(idToFind) {
    for (let oa = 0; oa < this.objectsArray.length; oa++) {
      if (this.objectsArray[oa].id === idToFind) {
        const v = this.objectsArray[oa];
        return `${v.name} + ${v.price}`;
      }
    }
  },
};

Array.init(inputRandomArr);

const MyMap = {
  map: {},

  init(srcArray) {
    for (let i = 0; i < srcArray.length; i++) {
      this.map[srcArray[i].id] = srcArray[i];
    }
  },

  findAndFormat(idToFind) {
    const v = this.map[idToFind];
    return `${v.name} + ${v.price}`;
  },
};
MyMap.init(inputRandomArr);

TO_FIND_IDS = newRandomArr(TO_FIND_IDS);
TO_FIND_IDS = TO_FIND_IDS.slice(0, COUNT / 3);
// console.log(TO_FIND_IDS);

const EXPECTED_ARRAY = [];
for (let i = 0; i < TO_FIND_IDS.length; i++) {
  const toFindId = TO_FIND_IDS[i];
  const v = findIn_INPUT_ARRAY(toFindId);
  EXPECTED_ARRAY.push(`${v.name} + ${v.price}`);
}

doBench("Test Array", (iterCount) => {
  let results = [];
  for (let iter_index = 0; iter_index < iterCount; iter_index++) {
    results = [];

    for (let tfi = 0; tfi < TO_FIND_IDS.length; tfi++) {
      const v = Array.findAndFormat(TO_FIND_IDS[tfi]);
      results.push(v);
    }
  }

  expectSameArray(results, EXPECTED_ARRAY);
});

doBench("Test Map", (iterCount) => {
  let results = [];
  for (let iter_index = 0; iter_index < iterCount; iter_index++) {
    results = [];

    for (let tfi = 0; tfi < TO_FIND_IDS.length; tfi++) {
      const v = MyMap.findAndFormat(TO_FIND_IDS[tfi]);
      results.push(v);
    }
  }
  expectSameArray(results, EXPECTED_ARRAY, "Test Map");
});
