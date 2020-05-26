# 1. Colleting API data from CodeForces and IOI Stats

The IOI Stats [website](https://stats.ioinformatics.org/) does not have an API description so we have to analyze the structure of the html file.
CodeForces on the other hand has a good description of its api on the following [link](https://codeforces.com/apiHelp).

## 1.1 **IOI Participants**

The Participant's profile can be found under the link 'https://stats.ioinformatics.org/people/{Id}' where Id is the unique id of the participant.
For now the Id is a number with at most 4 digits and most likely will have 5 digits in the near future.
Under this page can be found all the participant's performances at IOI as well as links to other websites such as CodeForces Profile, Facebook Profile, etc.
The Participants of each year can be found under the link 'https://stats.ioinformatics.org/results/{Year}' where Year is the Year of the contest. 

### 1.1.1 **Participant's name**

Participant's name can be found uniquely in the occurence of '"participantname">< div>{Name}< /div>' in raw html.

### 1.1.2 **CodeForces Profile**

If the participant has a CodeForces Profile with handle `Handle` then it will be uniquely found in the occurence of 'codeforces.com/profile/{Handle}"' in raw html. 
Since CodeForces url links can be found both under 'http' and 'https', I ommited the left part of the url. 

CodeForces allows every year during a specific period users to change their handles. After an user changed their handle, the old one will still redirect to the new link
but the API doesn't not account for this change so it will return an error with user not found. To find this issue we will hit the the old profile which will redirect to
the new profile, then we can extract the profile of the new profile by looking at the redirected url.

### 1.1.3 **Performance**

Every year starts with  For every task present in the
contestant's years will be found under '"tasks/{Year}/{Name}">{Score}<' where `Name` is the name of the task, `Year` is the Year of the task and `Score` is the score
of the contestant on that task. The score is a real number usually between 0 and 100 with rare exceptions. 

### 1.1.4 **List of participants' ids**

The list of participants' ids can be found in the occurence of 'href="people/{Id}' where `Id` is a number with at most 5 digits.


