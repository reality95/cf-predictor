import json
from os import system
import IOI
import Constants
import CF

if __name__ == "__main__":
	Participants = {}
	Directory = '../../assets/IOI_results.json'
	system('touch ' + Directory + '\n')
	file = open(Directory,'w')
	for Year in range(2005,2020):
		sYear = str(Year)
		ParticipantsIDs = IOI.ExtractParticipantsIDs(Year)
		Rank = 0
		for Id in ParticipantsIDs:
			Rank += 1
			if not(Id in Participants):
				Participants[Id] = IOI.ExtractParticipantInfo(Id)
				Participants[Id]['Results'] = {}
				if Participants[Id]['CF Handle'] != None:
					Participants[Id]['Rating']= {}
			Participants[Id]['Results'][sYear] = Rank
			if Participants[Id]['CF Handle'] != None and 2010 <= Year:
				Participants[Id]['Rating'][sYear] = CF.RatingBefore(Participants[Id]['CF Handle'],Constants.IOI[sYear])
	file.write(json.dumps(Participants))
	file.close()