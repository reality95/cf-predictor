import json
from os import system
import IOI
import Constants
import CF

if __name__ == "__main__":
	Participants = {}
	system('touch ../../assets/IOI_results.json')
	file = open('../../assets/IOI_results.json','w')
	for Year in range(2005,2020):
		sYear = str(Year)
		ParticipantsIDs = IOI.ExtractParticipantsIDs(Year)
		for Id in ParticipantsIDs:
			if not(Id in Participants):
				Participants[Id] = IOI.ExtractParticipantInfo(Id)
		f_cvs = open('../../assets/' + str(Year) + 'result.cvs')
		Res = ''
		for Id in ParticipantsIDs:
			sId = str(Id)
			Rank = str(Participants[Id]['Results'][sYear])
			Rating = str(Participants[Id]['Rating'][sYear])
			Res += sId + ',' + Participants[Id]['CF Handle'] + ',' + Rank + ',' + Rating + '\n'
		f_cvs.write(Res)
		f_cvs.close()

	file.write(json.dumps(Participants))
	file.close()