// Highest Scoring Word

function high(x){
  const letters = 'abcdefghijklmnopqrstuvwxyz'
  const words = x.split(' ')
  let returnWord = ''
  let pointWord = 0
  for(const word of words) {
    const points = word
      .split('')
      .reduce((total, chr) => letters.indexOf(chr) + 1 + total , 0)
    if (points > pointWord) {
      returnWord = word
      pointWord = points
    }
  }
  return returnWord
}

console.log(high('man i need a taxi up to ubud'), 'taxi');
console.log(high('what time are we climbing up the volcano'), 'volcano'); 
console.log(high('take me to semynak'), 'semynak'); 
console.log(high('aa b'), 'aa');
console.log(high('b aa'), 'b');
console.log(high('bb d'), 'bb');
console.log(high('d bb'), 'd');
console.log(high('aaa b'), 'aaa');