def duplicate_count(text):
    repeat = {}
    text = text.lower()
    countsRepeat = 0
    for ch in text:
      try:
        times = repeat[ch]
        times += 1
        repeat[ch] = times
      except:
        repeat[ch] = 1
    for count in repeat.values():
      if(count > 1):
        countsRepeat += 1
    return countsRepeat