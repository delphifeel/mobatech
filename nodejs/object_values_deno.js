/* GOAL
 *
 * Create array of random values
 */

// class F32Slice {
//   constructor(cap) {
//     this.array = new Float32Array(cap);
//     this.len = 0;
//     this.cap = cap;
//   }

//   push(value) {
//     if (this.len == this.cap) {
//       this.cap *= 2;
//       const newArray = new Float32Array(this.cap);
//       newArray.set(this.array, 0);
//       this.array = newArray;
//     }

//     this.array[this.len++] = value;
//   }
// }

// Deno.bench("640 native numbers class", function () {
//   let float32Array = new F32Slice(ARRAY_SIZE);
//   for (let i = 0; i < ARRAY_SIZE; i++) {
//     float32Array.push(Math.random());
//   }
// });

const SMALL_ARRAY_SIZE = 640;
const MEDIUM_ARRAY_SIZE = 6400;
const BIG_ARRAY_SIZE = 64_000;

Deno.bench("640 native numbers", function () {
  let float32Array = new Float32Array(SMALL_ARRAY_SIZE);
  for (let i = 0; i < SMALL_ARRAY_SIZE; i++) {
    float32Array[i] = Math.random();
  }
});
Deno.bench("640 numbers", function () {
  let primitiveArray = [];
  for (let i = 0; i < SMALL_ARRAY_SIZE; i++) {
    primitiveArray.push(Math.random());
  }
});
Deno.bench("640 objects", function () {
  let objectArray = [];
  for (let i = 0; i < SMALL_ARRAY_SIZE; i++) {
    objectArray.push({ value: Math.random() });
  }
});

Deno.bench("6400 native numbers", function () {
  let float32Array = new Float32Array(MEDIUM_ARRAY_SIZE);
  for (let i = 0; i < MEDIUM_ARRAY_SIZE; i++) {
    float32Array[i] = Math.random();
  }
});
Deno.bench("6400 numbers", function () {
  let primitiveArray = [];
  for (let i = 0; i < MEDIUM_ARRAY_SIZE; i++) {
    primitiveArray.push(Math.random());
  }
});
Deno.bench("6400 objects", function () {
  let objectArray = [];
  for (let i = 0; i < MEDIUM_ARRAY_SIZE; i++) {
    objectArray.push({ value: Math.random() });
  }
});

Deno.bench("64_000 native numbers", function () {
  let float32Array = new Float32Array(BIG_ARRAY_SIZE);
  for (let i = 0; i < BIG_ARRAY_SIZE; i++) {
    float32Array[i] = Math.random();
  }
});
Deno.bench("64_000 numbers", function () {
  let primitiveArray = [];
  for (let i = 0; i < BIG_ARRAY_SIZE; i++) {
    primitiveArray.push(Math.random());
  }
});
Deno.bench("64_000 objects", function () {
  let objectArray = [];
  for (let i = 0; i < BIG_ARRAY_SIZE; i++) {
    objectArray.push({ value: Math.random() });
  }
});
