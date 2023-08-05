from setuptools import setup
from setuptools import find_packages
def readme():
    with open('README.md') as f:
        return f.read()

setup(
    name='ppipe',
    version='0.2.7',
    packages=find_packages(),
    package_data={'ppipe': ['logconfig.json','aoi.json','wrs_grid.csv']},
    url='https://github.com/samapriya/Planet-GEE-Pipeline-CLI',
    install_requires=['bs4>=0.0.1','clipboard>=0.0.4','earthengine-api>=0.1.138','future>=0.16.0',
                      'google-cloud-storage>=1.5.0','hurry.filesize>=0.9','planet>=1.1.0','pandas>=0.23.0','psutil>=5.4.5',
                      'pycrypto>=2.6.1','pyshp>=1.2.12','pytest>=3.5.1','requests-toolbelt>=0.8.0','retrying>=1.3.3','simplejson>=3.15.0','pypiwin32'],
    license='Apache 2.0',
    long_description=open('README.txt').read(),
    classifiers=(
        'Development Status :: 3 - Alpha',
        'Intended Audience :: Developers',
        'Intended Audience :: Science/Research',
        'Natural Language :: English',
        'License :: OSI Approved :: Apache Software License',
        'Programming Language :: Python',
        'Programming Language :: Python :: 2.7',
        'Operating System :: OS Independent',
        'Topic :: Scientific/Engineering :: GIS',
    ),
    author='Samapriya Roy',
    author_email='samapriya.roy@gmail.com',
    description='Planet API Pipeline & Google Earth Engine Batch Assets Manager with Addons',
    entry_points={
        'console_scripts': [
            'ppipe=ppipe.ppipe:main',
        ],
    },
)
