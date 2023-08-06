# itunesmusicsearch [![PyPI version](https://badge.fury.io/py/itunespy.svg)](http://badge.fury.io/py/itunesmusicsearch)

itunesmusicsearch is made to fetch necessary information about the track using iTunes Store API. itunesmusicsearch is made for Python 3.X.

## Installing
You can install it from *pip*:

    pip install itunesmusicsearch

Or you can simply clone this project anywhere in your computer:

    git clone https://github.com/fonrimss/itunesmusicsearch.git

And then enter the cloned repo and execute:

    python setup.py install
## Dependencies

itunesmusicsearch requires [Requests](https://github.com/kennethreitz/requests) installed.

## Examples and information

Search for a track:

```python
import itunesmusicsearch

track = itunesmusicsearch.search_track('Iter Impius')  # Returns a list
print(track[0].artist_name + ': ' + track[0].track_name + ' | Length: ' + str(track[0].get_track_time_minutes())) # Get info from the first result
```

