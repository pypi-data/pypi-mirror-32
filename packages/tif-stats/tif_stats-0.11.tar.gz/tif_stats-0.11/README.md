# Tif Stats

Generates basic statistics on a geojson clipped tif image

Example (assumues you have one or more tif and geojson files in same directory)

```python
import tif_stats

tif_stats.generate_stats()
```

## Original Code

Code for developing stats has been adopted from [developmentseed/SEZ-U](https://github.com/developmentseed/SEZ-U/) project.

## Publishing to PyPi

```
rm -rf build dist
python setup.py sdist
python setup.py bdist_wheel --universal
~/Library/Python/3.6/bin/twine upload dist/*
```

# TODO

* File arguments - should be able to specify local or remote tif and geojson files
* Destination arguments - should be able to specify local or remote places to store results
