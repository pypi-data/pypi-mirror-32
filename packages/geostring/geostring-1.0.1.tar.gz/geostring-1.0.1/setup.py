from setuptools import setup
setup(
  name = 'geostring',
  packages = ['geostring'], # this must be the same as the name above
  version = '1.0.1',
  description = 'geostring',
  author = 'Deen Freelon',
  author_email = 'dfreelon@gmail.com',
  url = 'https://github.com/dfreelon/geostring/', # use the URL to the github repo
  download_url = 'https://github.com/dfreelon/geostring/', 
  install_requires = ['editdistance','unidecode'],
  keywords = ['geographic', 'location', 'places', 'geolocation'], # arbitrary keywords
  classifiers = [],
)