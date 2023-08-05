# from setuptools import setup
#
# setup(
#     name = 'pybcoin',
#     packages = ['pybcoin'],
#     version = '0.0.1',
#     description = 'Bitcoin Trend Forecast',
#     author='Pushottam Shivraj',
#     author_email='pshivraj@uw.edu',
#     url='https://github.com/rguptauw/pybcoin',
#     classifiers=[],
# )
from distutils.core import setup
setup(
  name = 'pybcoin',
  packages = ['pybcoin'], # this must be the same as the name above
  version = '0.1',
  description = 'Bitcoin Trend Forecast',
 author='Pushottam Shivraj',
 author_email='pshivraj@uw.edu',
 url='https://github.com/rguptauw/pybcoin',
  keywords = ['testing', 'Forecasting', 'Timeseries'], # arbitrary keywords
  classifiers = [],
)
