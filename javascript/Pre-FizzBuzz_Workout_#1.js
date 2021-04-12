const preFizz = n => 
  Array(n).fill(0).map((_, index) => index+1)

console.log(preFizz(5));