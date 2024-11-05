import { arrayOfRandomNumbers } from "./db.mjs";

const prepare = (mul) => {
  const BUILDS_COUNT = 4 * mul;
  const ABILITIES_COUNT = 4 * mul;
  const RATES_COUNT = 16 * mul;

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

    const indexFn = (i, lvl_i, ab_i) => {
      return i * RATES_COUNT * ABILITIES_COUNT + lvl_i * ABILITIES_COUNT + ab_i;
    };

    let buildToRateToAbility = new Uint16Array(
      BUILDS_COUNT * RATES_COUNT * ABILITIES_COUNT
    );
    for (let i = 0; i < BUILDS_COUNT; i++) {
      const inputBuild = input[i];
      for (let lvl_i = 0; lvl_i < RATES_COUNT; lvl_i++) {
        for (let ab_i = 0; ab_i < ABILITIES_COUNT; ab_i++) {
          buildToRateToAbility[indexFn(i, lvl_i, ab_i)] =
            inputBuild.abilityPickRates[ab_i].rates[lvl_i];
        }
      }
    }

    return {
      input,
      expected,
      indexFn,
      buildToRateToAbility,
      BUILDS_COUNT,
      ABILITIES_COUNT,
      RATES_COUNT,
    };
  }
};

// console.log(expected);

const small = prepare(1);
const bigger = prepare(10);
const biggest = prepare(100);

import Benchmark from "benchmark";
var suite = new Benchmark.Suite();

// add tests
suite
  .add("x1     ", function () {
    const { input, BUILDS_COUNT, ABILITIES_COUNT, RATES_COUNT } = small;

    const result = [];
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

      result.push(expectedBuild);
    }
  })
  .add("x1 (fast)", function () {
    const {
      indexFn,
      buildToRateToAbility,
      BUILDS_COUNT,
      ABILITIES_COUNT,
      RATES_COUNT,
    } = small;

    let result = [];

    for (let i = 0; i < BUILDS_COUNT; i++) {
      const expectedBuild = [];
      for (let lvl_i = 0; lvl_i < RATES_COUNT; lvl_i++) {
        let max = 0;
        for (let ab_i = 0; ab_i < ABILITIES_COUNT; ab_i++) {
          const rate = buildToRateToAbility[indexFn(i, lvl_i, ab_i)];
          max = Math.max(max, rate);
        }
        expectedBuild.push(max);
      }

      result.push(expectedBuild);
    }
  })
  .add("x10      ", function () {
    const { input, BUILDS_COUNT, ABILITIES_COUNT, RATES_COUNT } = bigger;

    const result = [];
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

      result.push(expectedBuild);
    }
  })
  .add("x10 (fast)", function () {
    const {
      indexFn,
      buildToRateToAbility,
      BUILDS_COUNT,
      ABILITIES_COUNT,
      RATES_COUNT,
    } = bigger;

    let result = [];

    for (let i = 0; i < BUILDS_COUNT; i++) {
      const expectedBuild = [];
      for (let lvl_i = 0; lvl_i < RATES_COUNT; lvl_i++) {
        let max = 0;
        for (let ab_i = 0; ab_i < ABILITIES_COUNT; ab_i++) {
          const rate = buildToRateToAbility[indexFn(i, lvl_i, ab_i)];
          max = Math.max(max, rate);
        }
        expectedBuild.push(max);
      }

      result.push(expectedBuild);
    }
  })
  .add("x100      ", function () {
    const { input, BUILDS_COUNT, ABILITIES_COUNT, RATES_COUNT } = biggest;

    const result = [];
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

      result.push(expectedBuild);
    }
  })
  .add("x100 (fast)", function () {
    const {
      indexFn,
      buildToRateToAbility,
      BUILDS_COUNT,
      ABILITIES_COUNT,
      RATES_COUNT,
    } = biggest;

    let result = [];

    for (let i = 0; i < BUILDS_COUNT; i++) {
      const expectedBuild = [];
      for (let lvl_i = 0; lvl_i < RATES_COUNT; lvl_i++) {
        let max = 0;
        for (let ab_i = 0; ab_i < ABILITIES_COUNT; ab_i++) {
          const rate = buildToRateToAbility[indexFn(i, lvl_i, ab_i)];
          max = Math.max(max, rate);
        }
        expectedBuild.push(max);
      }

      result.push(expectedBuild);
    }
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
