const { expect } = require('@jest/globals')
const countPositivesSumNegatives = require('./Count_of_positives__sum_of_negatives')

test('Generate_range_of_integers', () => {
  expect(countPositivesSumNegatives([0, 2, 3, 0, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14])).toEqual([8, -50])
})
    
