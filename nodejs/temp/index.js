import { doBench, newRandomArr, prepareBench } from "./db.mjs";

prepareBench(
  () => {
    const COUNT = 30;
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

    const findIn_INPUT_ARRAY = (id) => {
      for (let i = 0; i < INPUT_ARRAY.length; i++) {
        if (INPUT_ARRAY[i].id === id) {
          return INPUT_ARRAY[i];
        }
      }
    };

    TO_FIND_IDS = newRandomArr(TO_FIND_IDS);
    TO_FIND_IDS = TO_FIND_IDS.slice(0, COUNT / 3);

    const EXPECTED_ARRAY = [];
    for (let i = 0; i < TO_FIND_IDS.length; i++) {
      const toFindId = TO_FIND_IDS[i];
      const v = findIn_INPUT_ARRAY(toFindId);
      EXPECTED_ARRAY.push(`${v.name} + ${v.price}`);
    }

    return {
      input: newRandomArr(INPUT_ARRAY),
      toFind: TO_FIND_IDS,
      expected: EXPECTED_ARRAY,
    };
  },
  {
    iterations: 100000,
  }
);

doBench("Test Array", (iterCount, input, toFind) => {
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
  Array.init(input);
  let results = [];

  for (let iter_index = 0; iter_index < iterCount; iter_index++) {
    results = [];

    for (let tfi = 0; tfi < toFind.length; tfi++) {
      const v = Array.findAndFormat(toFind[tfi]);
      results.push(v);
    }
  }

  return results;
});

doBench("Test FasterArray", (iterCount, input, toFind) => {
  const FasterArray = {
    ids: [],
    formatted: [],

    init(srcArray) {
      const srcCopy = srcArray;
      // const srcCopy = [...srcArray];
      // srcCopy.sort((a, b) => a.id - b.id);
      for (let i = 0; i < srcCopy.length; i++) {
        this.ids.push(srcCopy[i].id);
        this.formatted.push(`${srcCopy[i].id} name + ${srcCopy[i].price}`);
      }
    },

    findAndFormat(idToFind) {
      for (let i = 0; i < this.ids.length; i++) {
        if (this.ids[i] === idToFind) {
          return this.formatted[i];
        }
      }
    },
  };

  FasterArray.init(input);
  let results = [];

  for (let iter_index = 0; iter_index < iterCount; iter_index++) {
    results = [];

    for (let tfi = 0; tfi < toFind.length; tfi++) {
      const v = FasterArray.findAndFormat(toFind[tfi]);
      results.push(v);
    }
  }

  return results;
});

doBench("Test Map", (iterCount, input, toFind) => {
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
  MyMap.init(input);
  let results = [];

  for (let iter_index = 0; iter_index < iterCount; iter_index++) {
    results = [];

    for (let tfi = 0; tfi < toFind.length; tfi++) {
      const v = MyMap.findAndFormat(toFind[tfi]);
      results.push(v);
    }
  }
  return results;
});
