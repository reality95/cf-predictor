import requests
import re
import CF

def ExtractParticipantInfo(ID):
	"""
	Returns a dictionary containing participant's CodeForcer (CF) Handle if
	it exists, otherwise set to None, the full Name of the participant and
	all the scores from every year he or she participated under the format
	dict["Tasks"][Year][Name of task].
	"""

	Info = {"CF Handle" : None, "Name" : None, "Tasks" : {}, "Results" : {}}

	Cf_profile = 'codeforces.com/profile/'
	Pa_name = '\"participantname\"'

	Text = requests.get("https://stats.ioinformatics.org/people/" + ID).text
	Left = Text.find(Cf_profile)
	if Left != -1:
		Left += len(Cf_profile)
		Right = Left + Text[Left:].find('\"')
		Info["CF Handle"] = Text[Left : Right]

	Left = Text.find(Pa_name)
	if Left != -1:
		Left += Text[Left:].find('<div>') + len('<div>')
		Right = Left + Text[Left:].find('</div>')
		Info["Name"] = Text[Left : Right]

	Task_Code = '\"tasks/'
	LEN = len(Task_Code)

	Tasks = [k + LEN for k in range(len(Text)) if Text[k:k + LEN] == Task_Code]

	Years = []

	for task in Tasks:
		Year = None
		try:
			Year = int(Text[task : task + 4])
		except:
			continue

		Name = Text[task + 5 : task + Text[task:].find('\"')]
		Score = None

		try:
			Score = float(Text[task + Text[task:].find('>') + 1 : task + Text[task:].find('<')])
		except:
			continue

		sYear = str(Year)

		Years.append(Year)

		if not(sYear in Info["Tasks"]):
			Info["Tasks"][sYear] = {}

		Info["Tasks"][sYear][Name] = Score

	Years = list(set(Years))
	Years.sort(reverse = True)
	Rank_Code = ['class=\"gold\">','class=\"silver\">','class=\"bronze\">']
	LEN1 = len(Rank_Code[0])
	LEN2 = len(Rank_Code[1])

	#Most Probably no medal strings
	No_Medal = re.findall('>\d\d\d/\d\d\d<',Text)

	Ranks = [k for k in range(len(Text)) if (Text[k:k+LEN1] in Rank_Code or Text[k:k+LEN2] in Rank_Code)]
	_index = 0

	Info['Results'] = {}

	Used = set()

	for rank in Ranks:
		_rank = rank + Text[rank:].find('>') + 1
		if Text[_rank:].find('/') < Text[_rank:].find('<'):
			assert(_index < len(Years))
			Used.add(Text[_rank : _rank + Text[_rank:].find('<')])
			try:
				val = int(Text[_rank:_rank + Text[_rank:].find('/')])
			except:
				continue
			Info['Results'][str(Years[_index])] = val
			_index += 1

	for rank in No_Medal:
		_rank = rank[1:-1]
		if not (_rank in Used):
			assert(_index < len(Years))
			Info['Results'][str(Years[_index])] = int(_rank[0:3])
			_index += 1

	Info['Rating'] = {}


	if Info['CF Handle'] != None:
		for Year in Years:
			Info['Rating'][str(Year)] = CF.BeforeRating(Info['CF Handle'],Year = Year)

	return Info

def ExtractParticipantsIDs(Year):
	"""
	Extracts the IDs of all the participants of IOI in the year 'year'.
	Returns as a list of strings with 4 digits. All the IDs have at most 4
	digits.
	"""
	Text = requests.get("https://stats.ioinformatics.org/results/" + str(Year)).text
	IDs_5 = re.findall('href="people/\d\d\d\d\d\"',Text)
	IDs_4 = re.findall('href="people/\d\d\d\d\"',Text)
	IDs_3 = re.findall('href="people/\d\d\d\"',Text)
	IDs_2 = re.findall('href="people/\d\d\"',Text)
	IDs_1 = re.findall('href="people/\d\"',Text)
	return [ID[-6:-1] for ID in IDs_5] + [ID[-5:-1] for ID in IDs_4] + [ID[-4:-1] for ID in IDs_3] \
				+ [ID[-3:-1] for ID in IDs_2] + [ID[-2:-1] for ID in IDs_1]
