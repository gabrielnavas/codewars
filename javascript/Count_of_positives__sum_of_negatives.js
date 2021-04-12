const countPositivesSumNegatives = input => {
  if (!input || input.length === 0) return []
  const obj = input.reduce((obj, n) => {
    if(n > 0) obj.countPositives++
    else if(n < 0) obj.sumNegatives += n
    return obj
  }, {
    countPositives: 0,
    sumNegatives: 0
  })
  return [obj.countPositives, obj.sumNegatives]
}

module.exports = countPositivesSumNegatives