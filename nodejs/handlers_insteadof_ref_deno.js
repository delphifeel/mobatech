/* GOAL
 *
 * Create array of random values
 */

const prepSrc = (numbersCount) => {
  let numbers = [];
  for (let i = 0; i < numbersCount; i++) {
    numbers.push(i);
  }
  numbers.sort(() => Math.random() - 0.5);
  return numbers;
};

const ARRAY_SIZE = 1_000;

const randomNums = prepSrc(ARRAY_SIZE);

const people1 = [];
const items1 = [];
for (let i = 0; i < ARRAY_SIZE; i++) {
  const person = { id: i, name: `Person ${i}` };
  const item = { id: `${i}item`, owner: person, value: randomNums[i] };
  people1.push(person);
  items1.push(item);
}

const PERSON_ID_LESS = Math.floor(Math.random() * 10) + ARRAY_SIZE * 2;

let done1 = false;
Deno.bench("x1 references", function () {
  let results = [];
  for (let i = 0; i < items1.length; i++) {
    const item = items1[i];
    if (item.owner.id < PERSON_ID_LESS) {
      results.push(item.value);
    }
  }

  // if (!done1) {
  //   console.log(sum);
  //   done1 = true;
  // }
});

const people2 = [];
const items2 = [];
for (let i = 0; i < ARRAY_SIZE; i++) {
  const person = { id: i, name: `Person ${i}` };
  const item = { id: `${i}item`, ownerIndex: i, value: randomNums[i] };
  people2.push(person);
  items2.push(item);
}

let done2 = false;
Deno.bench("x1 handlers", function () {
  let results = [];
  for (let i = 0; i < items2.length; i++) {
    const item = items2[i];
    if (people2[item.ownerIndex].id < PERSON_ID_LESS) {
      results.push(item.value);
    }
  }

  // if (!done2) {
  //   console.log(sum);
  //   done2 = true;
  // }
});
