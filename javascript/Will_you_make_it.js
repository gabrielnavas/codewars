const zeroFuel = (distanceToPump, mpg, fuelLeft) => 
  ((fuelLeft * mpg) - distanceToPump) >= 0

module.exports = zeroFuel