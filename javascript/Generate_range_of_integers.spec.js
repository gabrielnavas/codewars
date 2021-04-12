const generateRange = require('./Generate_range_of_integers')

test('Generate_range_of_integers', () => {
  expect(generateRange(2, 10, 2)).toEqual([2,4,6,8,10]) 
  expect(generateRange(1, 10, 3)).toEqual([1,4,7,10]) 
  expect(generateRange(2, 10, 2)).toEqual([2,4,6,8,10])
})
    
