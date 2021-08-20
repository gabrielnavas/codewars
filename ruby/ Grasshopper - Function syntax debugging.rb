# Grasshopper - Function syntax debugging

def main(verb, noun)
  return "#{verb}#{noun}"
end

puts main('take ', 'item') == 'take item'
puts main('use ', 'sword') == 'use sword'