const longestRepetition = s => {
  if(s.length === 0) {
    return ['', 0]
  }
  const returnInfo = ['', 0]

  for (let index=0; index < s.length; index++) {
    const tmp = [s[index], 1]
    while (index < s.length && s[index] === s[index + 1]) {
      tmp[0] += s[index]
      tmp[1]++
      index++
    }
    if (tmp[1] > returnInfo[1]) {
      returnInfo[0] = tmp[0]
      returnInfo[1] = tmp[1]
    }
  }

  return [returnInfo[0][0], returnInfo[1]]
}

console.log(longestRepetition("aaaabb"),      ["a",4])
console.log(longestRepetition("bbbaaabaaaa"), ["a",4])
console.log(longestRepetition("cbdeuuu900"),  ["u",3])
console.log(longestRepetition("abbbbb"),      ["b",5])
console.log(longestRepetition("aabb"),        ["a",2])
console.log(longestRepetition("ba"),          ["b",1])
console.log(longestRepetition(""),            ["",0])
console.log(longestRepetition("aaaa"),          ["a",4])
console.log(longestRepetition("9fdnhdapmj913sv74gvrhhhvxukhrqevmuk1kdvrp6ifzatata4u0qphe4kt85bfp9rn8n74pnmvjt6v7jf9e72agq8tcnuxv0fctgissdxc9dh1qn48ntj5j8omtrltiayowiut0okfhijlxpoxnl5vtu8b53nkvnmn9qh38l55k3rhup6l2exk34uvx385bm5x6e9xm6xfwhudsbgrgppt39kwugcs09fia7r3qf1nfeybsiaq19d4rlfjnb3l7ga6iwt1fjrpjzvf0vmqabe85cnooz0g6mnksil203tcmvt0w8243gnl6i1mdc2lutj5hzbzcw8y0913kariyilt7u5jyswewg37hoophcagudqlar08k05s3p7kjp7hcyhbh0bvcxx2dsa6vc2ocm7q6mud6lz57yiwxhibhl6ljg5pomvi2rj62g8wvdjg0rv1r3go5hizsyh91blf043xqlq3qmhhucb6g3c0oifnrgk2s5la"), ["h",3])