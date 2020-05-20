import matplotlib.pyplot as plt
import json
#import cvs
import sys
import numpy as np
import math
from sklearn import linear_model
sys.path.append('../api')
import CF
import Constants

file = open('../../assets/IOI/results.json','r')

data = json.loads(file.read())

def PlotParticipants():
	"""
	Finds the plot of Year vs Participants with CodeForces Handle.
	The saved file goes to directory '../../assets'
	"""
	D = {}
	X = []
	Y = []
	Z = []
	for Id in data.keys():
		for Year in data[Id]['Results']:
			sYear = str(Year)
			if data[Id]['CF Handle']:
				if not(sYear in D):
					D[sYear] = 0
				D[sYear] = D[sYear] + 1

	X = [str((k//10) % 10) + str(k % 10) + '\'' for k in range(2005,2020)]

	for Year in range(2005,2020):
		sYear = str(Year)
		Y.append(D[sYear])
		Z.append(Constants.Contestants[sYear] - D[sYear])

	fig,ax = plt.subplots(figsize = (10,8))

	ax.bar(X,Y,align='center',label = 'IOI Participants with a CF handle')
	ax.bar(X,Z,align='center',bottom = Y,label = 'IOI Participants without a CF handle')
	ax.set_xlabel('Year')
	ax.set_ylabel('Participants')
	ax.legend()
	fig.savefig('../../assets/cf_handles.png')

def RankPlot(Year):
	"""
	Finds the plot Rank vs Rating for the year `Year`. The saved file
	goes to directory '../../assets/IOI'
	"""
	fig,ax = plt.subplots(figsize=(10,8))
	sYear = str(Year)
	file = open('../../assets/IOI' + sYear + '.csv','rt')
	Lines = file.read().split('\n')
	X,Y = [],[]
	for Line in Lines:
		if Line != '':
			Vars = Line.split(',')
			try:
				Rating = int(Vars[-1])
				Rank = int(Vars[-2])
				if Rating != 0 and Rating != 1500:
					X.append(Rank)
					Y.append(Rating)
			except:
				continue
	X = np.array([[x] for x in X])
	Y = np.array([[y] for y in Y])
	regr = linear_model.LinearRegression()
	regr.fit(X,Y)
	Eq = 'f(x) = ' + str(round(regr.coef_[0,0],3)) + 'x + ' + str(round(regr.intercept_[0],3))
	ax.plot(X,Y,'ro')
	ax.plot(X,regr.predict(X),label = Eq)
	ax.set_xlabel('Rank')
	ax.set_ylabel('CodeForces Rating')
	ax.legend()
	fig.savefig('../../assets/IOI' + sYear + '.png')

def TaskPlot(Year,Task):
	"""
	Finds the plot Rating vs Points for the task `Task` from year
	`Year`. The saved file goes to directory '../../assets/IOI'.
	"""
	sYear = str(Year)
	X,Y = [],[]
	for Id in data.keys():
		if sYear in data[Id]['Tasks'] and sYear in data[Id]['Results'] and Task in data[Id]['Tasks'][sYear]:
			Rating = None
			try:
				Rating = data[Id]['Rating'][sYear]
			except:
				pass
			if Rating != None and Rating != 0 and Rating != 1500:
				X.append(Rating)
				Y.append(data[Id]['Tasks'][sYear][Task])
	fig,ax = plt.subplots(figsize = (10,8))
	ax.plot(X,Y,'ro',label = 'Task: ' + Task.capitalize() + ' ' + sYear)
	ax.set_xlabel('Rating')
	ax.set_ylabel('Points')
	ax.legend()
	fig.savefig('../../assets/IOI/' + Task.capitalize() + ' ' + sYear + '.png')

def ParticipantsPlot(LowerBound):
	f = open('../../assets/CF/results.json','r')
	Rounds = {}
	data = json.loads(f.read())
	Contestants_ = 0
	for contestant in data['Contestants'].keys():
		RoundSet = set()
		for Round,_ in data['Contestants'][contestant]:
			RoundSet.add(Round)
		if len(RoundSet) > LowerBound:
			Contestants_ += 1
			for Round in RoundSet:
				if Round not in Rounds:
					Rounds[Round] = 0
				Rounds[Round] += 1
	print('Nr of Contestants: ',Contestants_)
	X = [Round for Round in Rounds.keys()]
	X.sort()
	Y = [Rounds[Round] for Round in X]
	plt.plot(X,Y,'ro')
	plt.show()

if __name__ == "__main__":
	ParticipantsPlot(int(input()))
