def duplicate_encode(word):
    word = word.lower()
    words = {}
    string = ''
    for ch in word: 
      try:
        words[ch] += 1
      except:
        words[ch] = 1
    for ch in word:
      num = words[ch]
      if num == 1:
        string += '('
      else:
        string += ')'
    return string