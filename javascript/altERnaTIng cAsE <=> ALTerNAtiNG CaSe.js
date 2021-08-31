// altERnaTIng cAsE <=> ALTerNAtiNG CaSe


String.prototype.toAlternatingCase = function () {
  const minLower = 97
  const maxLower = 122
  const minUpper = 65
  const maxUpper = 90
  
  return this.toString().split('').map(chr => {
    if(chr.charCodeAt(0) >= minLower && chr.charCodeAt(0) <= maxLower) {
      return chr.toUpperCase()
    }
    else if(chr.charCodeAt(0) >= minUpper && chr.charCodeAt(0) <= maxUpper) {
      return chr.toLowerCase()
    } 
    else return chr
  }).join('')
}

console.log("hello world".toAlternatingCase(), "HELLO WORLD");
console.log("HELLO WORLD".toAlternatingCase(), "hello world");
console.log("hello WORLD".toAlternatingCase(), "HELLO world");
console.log("HeLLo WoRLD".toAlternatingCase(), "hEllO wOrld");
console.log("12345".toAlternatingCase(), "12345");
console.log("1a2b3c4d5e".toAlternatingCase(), "1A2B3C4D5E");
console.log("String.prototype.toAlternatingCase".toAlternatingCase(), "sTRING.PROTOTYPE.TOaLTERNATINGcASE");
console.log("Hello World".toAlternatingCase().toAlternatingCase(), "Hello World");
