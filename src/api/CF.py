import requests
import json
import time
import Constants


def LatestHandle(Handle):
	"""
	Returns the lastest version of the user associated previously with handle 
	`Handle`
	"""
	time.sleep(0.2)
	UNQ_STR = 'https://codeforces.com/profile/'
	try:
		Url = requests.get(UNQ_STR + Handle).url
		return Url[len(UNQ_STR) : ]
	except:
		return Handle

def RatingBefore(Handle,Time = None,Year = None):
	"""
	Finds the rating of User with handle `Handle` before time `Time` 
	given in epoch counters. If `Year` is not None then we set `Time`
	to be the starting of IOI on year `Year`
	"""
	if Year != None:
		Time = Constants.IOI[str(Year)]
	Handle = LatestHandle(Handle)
	data = None
	time.sleep(0.2)
	try:
		data = requests.get('https://codeforces.com/api/user.rating?handle=' + Handle).json()
	except:
		return None
	if data['status'] == 'FAILED':
		return None
	Rating = 1500
	for Contest in data['result']:
		Time_ = int(Contest['ratingUpdateTimeSeconds'])
		NewRat = int(Contest['newRating'])
		if Time == None or Time_ <= Time:
			Rating = NewRat
		else:
			break
	return Rating


def GetHandlesInfo(Handles):
	"""
	Returns the information related to the handles in list `Handles`.
	The list should have the len between 1 and 10000. Returns a list
	for every handle a dictionary containing the 'Name', 'Rating',
	'maxRating', 'CF Handle', 'Rank', 'maxRank'.
	"""
	LEN = len(Handles)
	if (LEN == 0 or LEN > 10000):
		return None
	sHandles = ''
	for k in range(LEN):
		sHandles += str(LatestHandle(Handles[k]))
		if k + 1 < len(Handles):
			sHandles += ';'
	time.sleep(0.2)
	data = requests.get('https://codeforces.com/api/user.info?handles=' + sHandles).json()
	if not('status' in data) or data['status'] != 'OK':
		return None
	data = data['result']
	Info = {}
	for k in range(LEN):
		Info[data[k]['handle']] = {
			'Name' : data[k]['firstName'] + ' ' + data[k]['lastName'],
			'Rating' : data[k]['rating'],
			'maxRating' : data[k]['maxRating'],
			'Rank' : data[k]['rank'],
			'maxRank' : data[k]['maxRank']
		}
	return Info

def ContestStandings(contestId,Start,End,Room = None,showUnofficial = False):
	"""
	Returns the Contest details of the contest with Id 'contestId' and the standings
	starting with position 'Start' and ending in 'End', with optional filters such as
	Room, (None if filter if off) and showUnofficial (False if filter is off).
	"""
	time.sleep(0.2)
	data = requests.get('https://codeforces.com/api/contest.standings?contestId=' + str(contestId) \
						+ '&from=' + str(Start) + '&count=' + str(End-Start+1) + ('' if not Room else '&room=' + str(Room)) \
						+ '&showUnofficial=' + str(showUnofficial).lower()).json()
	if data['status'] != 'OK':
		print(data['status'],contestId)
		return None
	data = data['result']
	Info = {}
	Info['Contest'] = {
		'Id' : data['contest']['id'],
		'Name' : data['contest']['name'],
		'Duration' : data['contest']['durationSeconds'],
	}
	Info['Problems'] = []
	Info['Standings'] = {}
	LEN = len(data['problems'])
	for problem in range(LEN):
		Info['Problems'].append({
			'ConstestId' : data['problems'][problem]['contestId'],
			'Index' : data['problems'][problem]['index'],
			'Name' : data['problems'][problem]['name'],
			'Points' : data['problems'][problem]['points'],
			'Rating' : None if 'rating' not in data['problems'][problem] else data['problems'][problem]['rating'],
			'Tags' : data['problems'][problem]['tags']
		})
	for Contestant in range(len(data['rows'])):
		if len(data['rows'][Contestant]['party']['members']) == 1:
			Info['Standings'][data['rows'][Contestant]['party']['members'][0]['handle']] = {
				'Rank' : data['rows'][Contestant]['rank'],
				'Points' : data['rows'][Contestant]['points'],
				'Results' : data['rows'][Contestant]['problemResults']
			}
	return Info

def ContestsList(Gym = False):
	"""
	Returns the list with all CF contests, including Gym contests if `Gym` is true
	The returned dictionary contains for every problem, 'Name', 'Duration' and 'Time'
	"""
	time.sleep(0.2)
	data = requests.get('https://codeforces.com/api/contest.list?gym=' + str(Gym).lower()).json()
	if data['status'] != 'OK':
		return None
	data = data['result']
	Info = {}
	LEN = len(data)
	for Contest in range(LEN):
		if data[Contest]['relativeTimeSeconds'] > 0 and data[Contest]['type'] == 'CF':
			Info[data[Contest]['id']] = {
				'Name' : data[Contest]['name'],
				'Duration' : data[Contest]['durationSeconds'],
				'Time' : data[Contest]['startTimeSeconds']
			}
	return Info

def ProblemsList(Tags = None):
	"""
	Returns the list of problems filtered by list `Tags` (filter is off if `Tags` = None).
	The returned dictionary contains for every problem, 'contestId', 'Index', 'Name',
	'Points', 'Rating', 'Tags', 'solvedCount'
	"""
	sTags = ''
	if Tags != None:
		sTags = '?tags='
		for tag in Tags:
			sTags += tag
			sTags += ';'
		sTags = sTags[:-1]
	time.sleep(0.2)
	data = requests.get('https://codeforces.com/api/problemset.problems' + sTags).json()
	if data['status'] != 'OK':
		return None
	data = data['result']
	Info = {}
	LEN = len(data['problems'])
	for problem in range(LEN):
		if data['problems'][problem]['type'] == 'PROGRAMMING':
			Info[problem] = {
				'contestId' : data['problems'][problem]['contestId'],
				'Index' : data['problems'][problem]['index'],
				'Name' : data['problems'][problem]['name'],
				'Points' : float(data['problems'][problem]['points']) if 'points' in data['problems'][problem] else None,
				'Rating' : int(data['problems'][problem]['rating']) if 'rating' in data['problems'][problem] else None,
				'Tags' : data['problems'][problem]['tags'],
				'solvedCount' : data['problemStatistics'][problem]['solvedCount']
			}
	return Info
