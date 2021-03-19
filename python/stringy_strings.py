def stringy(size):
	one = True
	string = ''
	while size > 0:
		string += '1' if one else '0'
		size -= 1
		one = not one
	return string

print(stringy(1))