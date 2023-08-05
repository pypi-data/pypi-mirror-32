from setuptools import setup, find_packages

setup(name='jampy',
      version='5.0.17',
      description='JamPy: Jeans Anisotropic Modelling of Galactic Dynamics',
      long_description=open("README.rst").read(),
      url='http://purl.org/cappellari/software',
      author='Michele Cappellari',
      author_email='michele.cappellari@physics.ox.ac.uk',
      license='Other/Proprietary License',
      packages=find_packages(),
      package_data={'': ['*/*.txt']},
      install_requires=['numpy', 'scipy', 'matplotlib', 'plotbin', 'adamet'],
      classifiers=["Development Status :: 5 - Production/Stable",
                   "Intended Audience :: Developers",
                   "Intended Audience :: Science/Research",
                   "Operating System :: OS Independent",
                   "Programming Language :: Python :: 3"],
      zip_safe=True)
