const take = require('./Enumerable_Magic #25 _-_Take_the_First_N_Elements')

test('should work for sample tests', () => {
  expect(take([0, 1, 2, 3, 5, 8, 13], 3)).toEqual([0, 1, 2])
})
