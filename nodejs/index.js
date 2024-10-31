import { doBench, arrayOfRandomNumbers, prepareBench } from "./db.mjs";

const BUILDS_COUNT = 10;
const ABILITIES_COUNT = 13;
const RATES_COUNT = 200;

prepareBench(
  () => {
    let input = [];
    let expected = [];
    for (let i = 0; i < BUILDS_COUNT; i++) {
      const abilityPickRates = [];
      for (let ab_i = 0; ab_i < ABILITIES_COUNT; ab_i++) {
        let randomRates = arrayOfRandomNumbers(RATES_COUNT, 100);
        randomRates = randomRates.map((rate) => Math.floor(rate));
        abilityPickRates.push({
          ability: `#${ab_i + 1} ability`,
          rates: randomRates,
        });
      }

      const build = {
        buildID: `${Math.random() * 10} ID`,
        matches: Math.floor(Math.random() * 100),
        wins: Math.floor(Math.random() * 30),
        earlyGameItems: `${Math.floor(Math.random() * i)} items`,
        abilityPickRates,
      };
      input.push(build);
    }

    for (let i = 0; i < BUILDS_COUNT; i++) {
      const inputBuild = input[i];
      const expectedBuild = [];
      for (let lvl_i = 0; lvl_i < RATES_COUNT; lvl_i++) {
        let max = 0;
        for (let ab_i = 0; ab_i < ABILITIES_COUNT; ab_i++) {
          const rate = inputBuild.abilityPickRates[ab_i].rates[lvl_i];
          max = Math.max(max, rate);
        }
        expectedBuild.push(max);
      }

      expected.push(expectedBuild);
    }

    return {
      input,
      expected,
      toFind: [],
    };
  },
  {
    iterations: 100000,
  }
);

import Benchmark from "benchmark";
var suite = new Benchmark.Suite();

// add tests
suite
  .add("RegExp#test", function () {
    /o/.test("Hello World!");
  })
  .add("String#indexOf", function () {
    "Hello World!".indexOf("o") > -1;
  })
  // add listeners
  .on("cycle", function (event) {
    console.log(String(event.target));
  })
  .on("complete", function () {
    console.log("Fastest is " + this.filter("fastest").map("name"));
  })
  // run async
  .run({ async: true });

// doBench("#1", (iterCount, input) => {
//   let result = [];
//   for (let iter_index = 0; iter_index < iterCount; iter_index++) {
//     result = [];
//     for (let i = 0; i < BUILDS_COUNT; i++) {
//       const inputBuild = input[i];
//       const expectedBuild = [];
//       for (let lvl_i = 0; lvl_i < RATES_COUNT; lvl_i++) {
//         let max = 0;
//         for (let ab_i = 0; ab_i < ABILITIES_COUNT; ab_i++) {
//           const rate = inputBuild.abilityPickRates[ab_i].rates[lvl_i];
//           max = Math.max(max, rate);
//         }
//         expectedBuild.push(max);
//       }

//       result.push(expectedBuild);
//     }
//   }

//   return result;
// });

// doBench("#2", (iterCount, input) => {
//   let result = [];

//   const index = (i, lvl_i, ab_i) => {
//     return i * RATES_COUNT * ABILITIES_COUNT + lvl_i * ABILITIES_COUNT + ab_i;
//   };

//   let buildToRateToAbility = new Uint16Array(
//     BUILDS_COUNT * RATES_COUNT * ABILITIES_COUNT
//   );
//   for (let i = 0; i < BUILDS_COUNT; i++) {
//     const inputBuild = input[i];
//     for (let lvl_i = 0; lvl_i < RATES_COUNT; lvl_i++) {
//       for (let ab_i = 0; ab_i < ABILITIES_COUNT; ab_i++) {
//         buildToRateToAbility[index(i, lvl_i, ab_i)] =
//           inputBuild.abilityPickRates[ab_i].rates[lvl_i];
//       }
//     }
//   }

//   for (let iter_index = 0; iter_index < iterCount; iter_index++) {
//     result = [];
//     for (let i = 0; i < BUILDS_COUNT; i++) {
//       const expectedBuild = [];
//       for (let lvl_i = 0; lvl_i < RATES_COUNT; lvl_i++) {
//         let max = 0;
//         for (let ab_i = 0; ab_i < ABILITIES_COUNT; ab_i++) {
//           const rate = buildToRateToAbility[index(i, lvl_i, ab_i)];
//           max = Math.max(max, rate);
//         }
//         expectedBuild.push(max);
//       }

//       result.push(expectedBuild);
//     }
//   }

//   return result;
// });
