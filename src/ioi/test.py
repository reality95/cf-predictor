import matplotlib.pyplot as plt
import json
import sys
import numpy as np
from sklearn import linear_model
sys.path.append('../api')
import CF

file = open('../../assets/IOI_results.json','r')

data = json.loads(file.read())

Year = 2019

X,Y = [],[]

D = {}

for Id in data.keys():
	if data[Id]['CF Handle']:
		for Year in data[Id]['Results']:
			sYear = str(Year)
			if not(sYear in D):
				D[sYear] = 0
			D[sYear] = D[sYear] + 1

X = ()

for Year in D.keys():
	X += (Year,)
	Y.append(D[Year])


plt.xticks(arrange(len(Y)),Y)

plt.bar(X,Y,align='center')
plt.show()

"""
X = np.array([[x] for x in X])

Y = np.array([[y] for y in Y])

regr = linear_model.LinearRegression()

regr.fit(X,Y)

plt.plot(X,Y,'ro',X,regr.predict(X))
plt.show()
"""