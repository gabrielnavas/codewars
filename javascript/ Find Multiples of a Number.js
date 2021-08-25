// Find Multiples of a Number


function findMultiples(integer, limit) {
  const multiples = []
  for(let i=integer; i <= limit; i+=integer) {
    if (i % integer === 0) {
      multiples.push(i)
    }
  }
  return multiples
}

console.log(findMultiples(5, 25));
console.log(findMultiples(1, 2));
console.log(findMultiples(5, 7));
