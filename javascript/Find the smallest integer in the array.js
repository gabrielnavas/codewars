class SmallestIntegerFinder {
  findSmallestInt = args => args.reduce( 
    (min, n) => (n < min) ? n : min 
  )
}

let sif = new SmallestIntegerFinder();
console.log(sif.findSmallestInt([78,56,232,12,8]),8,'Should return the smallest int 8');
console.log(sif.findSmallestInt([78,56,232,12,18]),12,'Should return the smallest int 12');
console.log(sif.findSmallestInt([78,56,232,412,228]),56,'Should return the smallest int 56');
console.log(sif.findSmallestInt([78,56,232,12,0]),0,'Should return the smallest int 0');
console.log(sif.findSmallestInt([1,56,232,12,8]),1,'Should return the smallest int 1');