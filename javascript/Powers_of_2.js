const powersOfTwo = (nonNegative) => {
  const startArray = [1]
  return startArray
    .concat(Array(nonNegative)
    .fill(2))
    .map((two,index) => Math.pow(two, index))
}

console.log(powersOfTwo(0));