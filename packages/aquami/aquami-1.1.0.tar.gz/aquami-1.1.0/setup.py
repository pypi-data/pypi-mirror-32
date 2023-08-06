from setuptools import setup, find_packages
from os import listdir
from os.path import isfile, join
ttps = [f for f in listdir('aquami/ttps') if isfile(join('aquami/ttps', f))]
ttps = [''.join(('aquami/',f)) for f in ttps]
micro = [f for f in listdir('aquami/ttps') if isfile(join('aquami/ttps', f))]
micro = [''.join(('aquami/',f)) for f in ttps]
print(ttps)
with open("README.rst", 'r') as f:
    long_description = f.read()

setup(
   name='aquami',
   version='1.1.0',
   description=''.join(('A module to extract quantitative microstructure ',
                        'information from micrographs of morphologically ',
                        'complex microstructures.')),
   license="MIT",
   long_description=long_description,
   author='Joshua Stuckner',
   author_email='stuckner@vt.edu',
   url="https://github.com/JStuckner/aquami",
   packages=['aquami'],
   install_requires=['matplotlib>=1.5.3',
                     'numpy>=1.12.0',
                     'openpyxl>=2.3.2',
                     'pandas>=0.18.1',
                     'scipy>=0.18.1',
                     'scikit_image>=0.12.3',
                     'Pillow>=4.0.0',
                     'setuptools>=34.3.2'],
   python_requires='>=3',
   include_package_data=True,
   package_data={
       '': ['icon.ico', 'Instruction manual.pdf'],
       'tpps': ttps,
	   'micrographs': micro,
    }
)

