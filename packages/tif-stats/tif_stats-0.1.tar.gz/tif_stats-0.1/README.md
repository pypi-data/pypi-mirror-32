# Tif Stats

Generates basic statistics on a geojson clipped tif image

Example (assumues you have one or more tif and geojson files in same directory)

```python
import tif_stats

tif_stats.generate_stats()
```

## Original Code

Code for developing stats has been adopted from [developmentseed/SEZ-U](https://github.com/developmentseed/SEZ-U/) project.

# TODO

* File arguments - should be able to specify local or remote tif and geojson files
* Destination arguments - should be able to specify local or remote places to store results
