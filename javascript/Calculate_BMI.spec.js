const bmi = require('./Calculate_BMI')

test('should Normal', () => {
  expect(bmi(80, 1.80)).toBe("Normal")
})