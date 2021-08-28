const rot13  = message => {
  let rot13 = ''
  for(const chr of message) {
    if (chr.codePointAt() >= 65 && chr.codePointAt() <= 77 ||
    chr.codePointAt() >= 97 && chr.codePointAt() <= 109) {
      rot13 += String.fromCharCode(chr.codePointAt() + 13)
    } else if (chr.codePointAt() >= 78 && chr.codePointAt() <= 90 ||
      chr.codePointAt() >= 110 && chr.codePointAt() <= 122) {
        rot13 += String.fromCharCode(chr.codePointAt() - 13)
    } else {
      rot13 += chr
    }
  }
  return rot13
}

console.log("grfg" == rot13("test"))
console.log("Grfg" == rot13("Test"))           