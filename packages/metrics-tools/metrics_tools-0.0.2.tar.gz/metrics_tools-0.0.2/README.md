# metrics_tools

A simple python package for making API calls to metrics.tools API.
https://metrics.tools/

It builds upon the requests package. The goal is to achiev an easy syntax for the API calls like: `mt.domain('example.com').rankings.get()`

*Note: I am not professional programmer. Install at your own risk. Useful tips are welcome :-).*

## Install via pip
```python
pip install metrics-tools
```

## Basic usage
```python
from metrics_tools import MetricsTools

api_key = 'your api key'
mt = MetricsTools(api_key, verbose=True)

domain = mt.domain('example.com')

visibility = domain.sk.get()
rankings = domain.rankings.get()

keyword_details = mt.keyword('api').details.get()
```

## Response
The response is parsed into a simple object with fields as object attributes. Additionaly it contains the requests response as '._response'.
```python
example_domain.values # if values are present
example_domain.json # parsed json response
example_domain.result # 'success' or 'error'
```

## More Parameters
The API supports some additional parameters, respectifly filters, to narrow down the API answer. These filters can be used through the avalible functions.
```python
#Visibility for specific date
domain.sk.from_('2018-05-07').get()

#Min and max visibility value
domain.sk.minmax().get()

#Ranking position for domain and given keyword
domain.position.for_keyword('jens fauldrath').get()
```

API calls
=========
The minimum interval for calls is 0.2 sec.