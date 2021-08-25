function sortThePile(pileOfTowels, weeklyUsedTowels) {
  for(const t of weeklyUsedTowels) {
    const start = pileOfTowels.length - t
    const usedTowels = pileOfTowels.splice(start, t)
    const usedTowelsSorted = usedTowels.sort((t1, t2) => t1.length - t2.length)
    pileOfTowels.push(...usedTowelsSorted)
  }
  return pileOfTowels;
}


// example 1
const pileOfTowels1 = ["blue", "red", "blue", "red", "blue"];
const weeklyUsedTowels1 = [3];
const expected = ["blue", "red", "red", "blue", "blue"];
console.log(sortThePile(pileOfTowels1, weeklyUsedTowels1))

const pileOfTowels2 = ["blue", "red", "blue", "red", "blue"];
const weeklyUsedTowels2 = [2, 1, 4, 2];
console.log(sortThePile(pileOfTowels2, weeklyUsedTowels2))