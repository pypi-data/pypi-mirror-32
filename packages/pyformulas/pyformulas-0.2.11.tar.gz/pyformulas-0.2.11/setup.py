from setuptools import setup
setup(
    name = 'pyformulas',
    packages = ['pyformulas'],
    include_package_data = True,
    python_requires='>=3.0',
    install_requires=[
            'pyaudio',
            'opencv-python',
            'numpy',
            'urllib3',
            'matplotlib'
        ],
    version = '0.2.11',
    description = 'A library of ready-to-go Python formulas',
    author = 'pyformulas',
    author_email = 'pyformulas@gmail.com',
    url = 'https://github.com/pyformulas/pyformulas',
    keywords = ['python', 'formulas', 'recipes', 'cookbook', 'lazy', 'easy', 'quick', 'shortcut'],
    classifiers = [],
)