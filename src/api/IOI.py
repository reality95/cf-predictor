import requests
import re
import CF

def ExtractParticipantInfo(ID):
	"""
	Returns a dictionary containing participant's CodeForcer (CF) Handle if
	it exists, otherwise set to None, the full Name of the participant and
	all the scores from every year he or she participated under the format
	dict["Tasks"][Year][Name of task]. The participant is defined by their
	id `ID` with at most 5 digits.
	"""

	Info = {"CF Handle" : None, "Name" : None, "Tasks" : {}, "Results" : {}, 'Rating' : {}}

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

	Ol_code = 'olympiads/'
	Ol_len = len(Ol_code)

	Task_Code = '\"tasks/'
	Task_len = len(Task_Code)

	Olympiads = [x.span()[0] for x in re.finditer(Ol_code + '\d\d\d\d\"',Text)]
	Olympiads.append(len(Text))

	for ol in range(len(Olympiads) - 1):

		Left,Right = Olympiads[ol],Olympiads[ol+1]

		Year = None

		try:
			Year = int(Text[Left + Ol_len : Left + Ol_len + 4])
		except:
			continue

		sYear = str(Year)

		Info['Tasks'][sYear] = {}

		Tasks = [Left + x.span()[1] for x in re.finditer(Task_Code + sYear + '/',Text[Left : Right])]
		
		for task in Tasks:

			Name = Text[task : task + Text[task:].find('\"')]
			Score = None

			try:
				Score = float(Text[task + Text[task:].find('>') + 1 : task + Text[task:].find('<')])
			except:
				continue

			Info["Tasks"][sYear][Name] = Score

		if not Info['Tasks'][sYear]:
			del Info['Tasks'][sYear]
		#len(Ranks) <= 1
		Ranks = re.findall('>\d\d\d/\d\d\d<',Text[Left:Right]) + re.findall('>\d\d/\d\d\d<',Text[Left:Right]) + re.findall('>\d/\d\d\d<',Text[Left:Right])
		for rank in Ranks:
			Info['Results'][sYear] = int(rank[1:rank.find('/')])

		if Year >= 2005 and Info['CF Handle'] != None:
			Info['Rating'][sYear] = CF.RatingBefore(Handle = Info['CF Handle'],Year = Year)

	return Info

def ExtractParticipantsIDs(Year):
	"""
	Extracts the IDs of all the participants of IOI in the year `Year`.
	Returns as a list of strings with at most 5 digits.
	"""
	Text = requests.get("https://stats.ioinformatics.org/results/" + str(Year)).text
	IDs_5 = re.findall('href="people/\d\d\d\d\d\"',Text)
	IDs_4 = re.findall('href="people/\d\d\d\d\"',Text)
	IDs_3 = re.findall('href="people/\d\d\d\"',Text)
	IDs_2 = re.findall('href="people/\d\d\"',Text)
	IDs_1 = re.findall('href="people/\d\"',Text)
	return [ID[-6:-1] for ID in IDs_5] + [ID[-5:-1] for ID in IDs_4] + [ID[-4:-1] for ID in IDs_3] \
				+ [ID[-3:-1] for ID in IDs_2] + [ID[-2:-1] for ID in IDs_1]
