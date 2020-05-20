import json
from os import system
import IOI
import Constants
import CF

def CollectIOIData(file):
	Participants = {}
	system('touch ' + file)
	f = open(file,'w')
	for Year in range(2005,2020):
		f_cvs = open('../../assets/IOI/IOI' + str(Year) + '.csv','w')
		sYear = str(Year)
		ParticipantsIDs = IOI.ExtractParticipantsIDs(Year)
		for Id in ParticipantsIDs:
			if not(Id in Participants):
				Participants[Id] = IOI.ExtractParticipantInfo(Id)
		Res = ''
		for Id in ParticipantsIDs:
			sId = str(Id)
			Rank = 'None'
			try:
				Rank = str(Participants[Id]['Results'][sYear])
			except:
				pass
			Rating = '0'
			if sYear in Participants[Id]['Rating']:
				Rating = str(Participants[Id]['Rating'][sYear])
			Res += sId + ',' + str(Participants[Id]['CF Handle']) + ',' + str(Participants[Id]['Name']) + ',' + Rank + ',' + Rating + '\n'
		f_cvs.write(Res)
		f_cvs.close()

	f.write(json.dumps(Participants))
	f.close()

def CollectCFData(file):
	data = {}
	Contestants = {}
	Problems = {}
	system('touch ' + file)
	f = open(file,'w')
	Contests = CF.ContestsList(Gym = False)
	print(len(Contests))
	cur = 0
	for contest in Contests:
		cur += 1
		Name = Contests[contest]['Name']
		ContestResults = CF.ContestStandings(contestId = contest,Start = 1,End = 10000)
		if not ContestResults:
			continue
		Problems[contest] = ContestResults['Problems']
		for contestant in ContestResults['Standings'].keys():
			if contestant not in Contestants:
				Contestants[contestant] = []
			Contestants[contestant].append((contest,ContestResults['Standings'][contestant]))
	data['Problems'] = Problems
	data['Contestants'] = Contestants
	f.write(json.dumps(data))
	f.close()

if __name__ == "__main__":
	CollectIOIData('../../assets/IOI/results.json')
	CollectCFData('../../assets/CF/results.json')