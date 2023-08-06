# dEF-Pi Python library

Python library for the dEF-Pi platform, to be used to create python services that are able to communicate with other 
services in the environment.

## Build details
```
rm -f dist/*
python3 setup.py sdist
python3 setup.py bdist_wheel
twine upload -r pypi dist/*
```

