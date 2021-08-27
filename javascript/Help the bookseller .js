const stockList = (listOfArt, listOfCat) => {
  if (listOfArt.length === 0) {
    return ''
  }
  
  return listOfCat.reduce((strFinal, cat) => {
    const bookValue = { 
      value: 0, 
      code: cat 
    }
    for (const book of listOfArt) {
      const [code, value] = book.split(' ')
      const firstLetterCode = code[0]
      if (firstLetterCode === bookValue.code) {
        bookValue.value += Number(value)
      }
    }
    return (strFinal.length === 0) ?
      `(${bookValue.code} : ${bookValue.value})`
      :  strFinal += ` - (${bookValue.code} : ${bookValue.value})`
  }, '')
}