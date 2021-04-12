const dutyFree = (normPrice, discount, hol) => 
  Math.floor(hol / (normPrice * discount / 100))

console.log(dutyFree(17,10,500));