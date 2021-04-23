const zeroFuel =require('./Will_you_make_it')

test("Sample Tests", () => {
  expect(zeroFuel(50, 25, 2)).toBe(true)
  expect(zeroFuel(100, 50, 1)).toBe(false)
});