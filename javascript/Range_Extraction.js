// wrong

const solution = list => {
  const sorted = list.sort((a, b) => a - b)
  const str = sorted.reduce((params, n, index) => {
    if(n+1 === sorted[index+1]) {
      if(!params.str[params.str.length-1] === '-') {

      }
    }
    else if (params.counter === 0){
      const newStr = 
        index === sorted.length -1 
          ? params.str 
          : params.str + `${n}, '`
          return {str: newStr, counter: 0}
        }
    else{
      if(!params.str[params.str.length-1] === '-')
      return {str: params.str + , counter: 0}
    }
  }, { str:'', counter: 0 })
  return sorted
}

console.log(solution([-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20]));