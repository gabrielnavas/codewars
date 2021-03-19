def count_languages(lst): 
	langs_counter = {}
	for item in lst:
		try:
			langs_counter[item['language']] += 1
		except:
			langs_counter[item['language']] = 1
	return langs_counter