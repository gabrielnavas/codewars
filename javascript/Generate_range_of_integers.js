const generateRange = (min, max, step) => {
  const numbers = []
  for (let n = min; n <= max; n+=step) {
    numbers.push(n)
  }
  return numbers
}


module.exports = generateRange