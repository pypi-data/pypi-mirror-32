from setuptools import setup, find_packages

setup(name='wbnn',
      version='0.1.1',
      description='White Box Neural Networks. Neural Network and Deep Learning toolkit developed for educational purposes.',
      url='https://github.com/develask/White-Box-Neural-Networks',
      project_urls={
            "Documentation": "https://github.com/develask/White-Box-Neural-Networks/wiki",
            "Source Code": "https://github.com/develask/White-Box-Neural-Networks",
      },
      author='Mikel de Velasco Vázquez',
      author_email='develascomikel@gmail.com',
      license='MIT',
      packages=['wbnn', 'wbnn.visualLogs'],
      entry_points={
          'console_scripts': [
              'wbnnplots = wbnn.visualLogs.plot:main'
          ]
      },
      install_requires=['numpy', 'matplotlib'],
      zip_safe=False)
