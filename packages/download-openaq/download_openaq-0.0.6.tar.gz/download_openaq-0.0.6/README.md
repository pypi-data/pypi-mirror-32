# Download OpenAQ

Downloads CSVs from OpenAQ S3

Example:

```python
import download_openaq

download_openaq.download()
```

## Publishing to PyPi

```
rm -rf build dist
python setup.py sdist
python setup.py bdist_wheel --universal
~/Library/Python/3.6/bin/twine upload dist/*
```

# TODO

* Landing directory - should be able to specify landing directory
* Dates - should be able to specify dates
* Concatenate - should be able to concatenate files
* Filter - should be able to filter data

