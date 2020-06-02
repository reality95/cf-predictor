import json
import torch
import torch.nn as nn
import torch.nn.functional as F
import doctest
import csv

class ProblemModel(nn.Module):
	"""docstring for ProblemModel"""
	def __init__(self, tags, scores):
		"""
			tags.shape = (1, P, T)
			scores.shape = (1, P, 1)
			self.layer.shape = (1, P, T)
		"""
		_, self.P, self.T = tags.shape
		super(ProblemModel, self).__init__()
		self.tags = tags.clone().detach()
		self.scores = scores.clone().detach()
		self.layer = torch.randn((1,self.P,self.T),dtype = torch.double,requires_grad = True)

	def forward(self, skill):
		"""
			skill.shape = (N, 1, T)
		"""
		score_pred = (self.layer * self.tags * skill).sum(axis = 2,keepdim = True)
		return torch.sigmoid(self.scores - score_pred)

	def backward(self, lr = 0.01):
		print('self.layer.grad=',self.layer.grad)
		with torch.no_grad():
			if self.layer.grad is not None:
				self.layer -= self.layer.grad * lr
				self.layer.grad.zero_()

def TrainProblemModel(tags, scores, skill, iters, result):
	"""
		tags.shape = (1, P, T)
		scores.shape = (1, P, 1)
		skill.shape = (N, 1, T)
		result.shape = (N, P, 1)
	"""
	criterion = torch.nn.MSELoss(reduction='mean')
	model = ProblemModel(tags = tags,scores = scores)
	for k in range(iters):
		loss = criterion(model.forward(skill.detach()),result.detach())
		print('loss = ',loss)
		loss.backward()
		model.backward()
	return model

if __name__ == '__main__':
	doctest.testmod()
	torch.manual_seed(69)
	tags = None
	result = None
	scores = None
	skill = None
	P = None
	T = None
	N = None
	with open('assets/ML/SampleProblems/problems.csv') as csvfile:
		c = csv.reader(csvfile,delimiter = ',')
		List = []
		scores = []
		for rowReader in c:
			row = []
			try:
				scores.append(int(rowReader[4]))
			except:
				_ = _
			for b in rowReader:
				if b in ['true','false']:
					row.append(True if b == 'true' else False)
			if len(row) > 0:
				List.append(row)
		P = len(List)
		T = len(List[0])
		tags = torch.zeros((1,P,T),dtype = torch.double,requires_grad = False)
		tags[torch.tensor(List).view(1,P,T)] = 1
		scores = torch.tensor(scores).view((1,P,1))
	with open('assets/ML/SampleProblems/scores.csv') as csvfile:
		c = csv.reader(csvfile, delimiter = ',')
		List = []
		for rowReader in c:
			try:
				List.append([int(x) for x in rowReader[4:]])
			except:
				continue
		N = len(List)
		result = torch.zeros((N,P,1),dtype = torch.double,requires_grad = False)
		result[torch.tensor(List).view((N,P,1)) > 0] = 1

	with open('assets/ML/SampleProblems/rating.csv') as csvfile:
		c = csv.reader(csvfile, delimiter = ',')
		List = []
		for rowReader in c:
			try:
				List.append(int(rowReader[1]))
			except:
				continue
		skill = torch.ones((N,1,T),dtype = torch.double,requires_grad = False) * torch.tensor(List).view((N,1,1))
	print(skill.shape,result.shape,scores.shape,tags.shape)
	TrainProblemModel(skill = skill,result = result,scores = scores,tags = tags,iters = 5)