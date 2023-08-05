from setuptools import setup, find_packages

setup(name='ppxf',
      version='6.7.8',
      description='pPXF: Full Spectrum Fitting of Galactic and Stellar Spectra',
      long_description=open("README.rst").read(),
      url='http://purl.org/cappellari/software',
      author='Michele Cappellari',
      author_email='michele.cappellari@physics.ox.ac.uk',
      license='Other/Proprietary License',
      packages=find_packages(),
      package_data={'ppxf': ['*/*.txt', '*/*.fits']},
      install_requires=['numpy', 'scipy', 'matplotlib', 'astropy'],
      classifiers=["Development Status :: 5 - Production/Stable",
                   "Intended Audience :: Developers",
                   "Intended Audience :: Science/Research",
                   "Operating System :: OS Independent",
                   "Programming Language :: Python :: 3"],
      zip_safe=True)
