
from setuptools import setup, find_packages

setup(
    # Needed to silence warnings (and to be a worthwhile package)
    name='distancesensor',
    author='Gishnu',
    author_email='gishnumohanmg@gmail.com',
    # Needed to actually package something
    packages = find_packages(),
    # Needed for dependencies
    install_requires=['RPi.GPIO'],
    # *strongly* suggested for sharing
    version='1.0',
    # The license can be anything you like
    license='MIT',
    description='Ultrasonic sensor values',
    # We will also need a readme eventually (there will be a warning)
    # long_description=open('README.txt').read(),
)


