# SupWSD

SupWSD is a Python binding to an HTTP RESTful service that gives you programmatic access to SupWSD, a framework for supervised Word Sense Disambiguation (WSD).

The SupWSD service is available only for English and enforces a default limit of 100 free requests per 1 hour period.

## Installation
```
pip install supwsd
```

## Code Example
```
from supwsd.wsd import SupWSD

text="The human " + SupWSD.SENSE_TAG + "brain" + SupWSD.SENSE_TAG + " is quite proficient at word-sense disambiguation. The fact that natural language is formed "+ SupWSD.SENSE_TAG+"in a way"+ SupWSD.SENSE_TAG+" that requires so much of it is a " + SupWSD.SENSE_TAG + "reflection" + SupWSD.SENSE_TAG + " of that neurologic reality."
    
for sense in SupWSD().senses(text):
	print("Word: {}\tLemma: {}\tPOS: {}\tSense: {}\tSource: {}\tCount: {}\tValid: {}\tMiss: {}".format(sense.word, sense.lemma, sense.pos,sense.key(),sense.source,sense.count(),sense.valid(),sense.miss()))

	for result in sense.results:
		print("Sense {}\tProbability: {}".format(result.key, result.prob))
```

## Links

* <a target="_blank" href="https://supwsd-supwsdweb.1d35.starter-us-east-1.openshiftapps.com/supwsdweb/index.jsp">SupWSD Toolkit</a>
* <a target="_blank" href="https://supwsd-supwsdweb.1d35.starter-us-east-1.openshiftapps.com/supwsdweb/api-python.jsp">SupWSD API Guide</a>
* <a target="_blank" href="https://supwsd-supwsdweb.1d35.starter-us-east-1.openshiftapps.com/supwsdweb/demo.jsp">SupWSD Demo</a>