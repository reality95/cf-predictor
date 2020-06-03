import json
import torch
import torch.nn as nn
import torch.nn.functional as F
import doctest
import csv

SCALE = 1000

P = None
T = None
N = None
tags = None
result = None
scores = None
skill = None

"""
Under the Codeforces Rating system, the ratings are upper bounded by 4000
and can run into negative but that's usually done on purpose by a hand of
users. The skill on every tag should be of the same magnitudes. To handle
overflow and underflow, the ratings and skills will be scaled down by 1000.
"""


class ProblemModel(nn.Module):
	"""docstring for ProblemModel"""
	def __init__(self, tags, scores):
		"""
			>>> tags.shape == (1, P, T)
			True
			>>> scores.shape == (1, P, 1)
			True
		"""
		_, self.P, self.T = tags.shape
		super(ProblemModel, self).__init__()
		self.tags = tags
		self.scores = scores / SCALE
		layer_ = torch.ones((1,self.P,self.T),dtype = torch.double) * tags / tags.sum(axis = 2,keepdim = True)
		print('layer_ = ',layer_)
		self.layer = torch.tensor(layer_,dtype = torch.double,requires_grad = True)
		self.bias = torch.randn((1,self.P,self.T),dtype = torch.double,requires_grad = True)

	def forward(self, skill):
		"""
			>>> skill.shape == (N, 1, T)
			True
		"""
		score_pred = (self.layer * self.tags * skill + self.bias).sum(axis = 2,keepdim = True)
		return score_pred,torch.tanh(score_pred / self.scores)

	def backward(self, lr = 0.5):
		with torch.no_grad():
			if self.layer.grad is not None:
				self.layer -= self.layer.grad * lr
				self.layer.grad.zero_()
			if self.bias is not None:
				self.bias -= self.bias * lr
				self.bias.grad.zero_()

def TrainProblemModel(tags, scores, skill, iters, result):
	"""
		>>> tags.shape == (1, P, T)
		True
		>>> scores.shape == (1, P, 1)
		True
		>>> skill.shape == (N, 1, T)
		True
		>>> result.shape == (N, P, 1)
		True
	"""
	skill = skill / SCALE
	model = ProblemModel(tags = tags,scores = scores)
	for k in range(iters):
		_,result_pred = model.forward(skill)
		loss = torch.mean((result_pred - result).pow(2))
		loss.backward()
		model.backward()
	return model

def getProblemsCSV(fileName):
	global P,T
	tags = None
	scores = None
	with open(fileName) as csvfile:
		c = csv.reader(csvfile,delimiter = ',')
		List = []
		scores = []
		for rowReader in c:
			row = []
			try:
				scores.append(int(rowReader[4]))
			except:
				pass
			for b in rowReader:
				if b in ['true','false']:
					row.append(True if b == 'true' else False)
			if len(row) > 0:
				List.append(row)
		P = len(List)
		T = len(List[0])
		tags = torch.zeros((1,P,T),dtype = torch.double,requires_grad = False)
		tags[torch.tensor(List).view(1,P,T)] = 1
		scores = torch.tensor(scores,dtype = torch.double).view((1,P,1))
	return tags,scores

def getScoresCSV(fileName):
	global N
	result = None
	with open(fileName) as csvfile:
		c = csv.reader(csvfile, delimiter = ',')
		List = []
		for rowReader in c:
			try:
				List.append([int(x) for x in rowReader[4:]])
			except:
				pass
		N = len(List)
		result = torch.zeros((N,P,1),dtype = torch.double,requires_grad = False)
		result[torch.tensor(List).view((N,P,1)) > 0] = 1
	return result

def getRatingCSV(fileName):
	skill = None
	with open(fileName) as csvfile:
		c = csv.reader(csvfile, delimiter = ',')
		List = []
		for rowReader in c:
			try:
				List.append(int(rowReader[1]))
			except:
				pass
		skill = torch.ones((N,1,T),dtype = torch.double,requires_grad = False) * torch.tensor(List).view((N,1,1))
	return skill

if __name__ == '__main__':
	torch.manual_seed(69)
	tags, scores = getProblemsCSV('assets/ML/SampleProblems/problems.csv')
	result = getScoresCSV('assets/ML/SampleProblems/scores.csv')
	skill = getRatingCSV('assets/ML/SampleProblems/rating.csv')

	doctest.testmod()
	model = TrainProblemModel(skill = skill,result = result,scores = scores,tags = tags,iters = 50)
	print(model.layer)
	print(model.bias)
	scores_pred,result_pred = model.forward(torch.ones((1,1,T),dtype = torch.float,requires_grad = False) * 4)
	print(scores_pred)
	print(result_pred)
