// Loop Array

function loopArr(arr, direction, steps) {
  if(direction === 'left') {
    const leftNums = arr.splice(0, steps)
    arr.push(...leftNums)
  }
  else if(direction === 'right') {
    const rightNums = arr.splice(arr.length-steps, steps)
    arr.unshift(...rightNums)
  }
  return arr
}

console.log(loopArr([1, 5, 87, 45, 8, 8], 'left', 2))

console.log(loopArr([1, 5, 87, 45, 8, 8], 'right', 2))