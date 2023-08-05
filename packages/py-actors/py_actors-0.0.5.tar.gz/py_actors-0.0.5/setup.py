import os
from setuptools import setup


# Utility function to read the README.md file.
def read(filename):
    return open(os.path.join(os.path.dirname(__file__), filename)).read()


setup(
    name="py_actors",
    version="0.0.5",
    author="Kevin Conley",
    author_email="kmanc@comcast.net",
    description="Actor implementation in Python",
    license='MIT',
    keywords="Python, Actor, Concurrency, Parallelism",
    url="https://github.com/kmanc/py_actors",
    packages=['py_actors'],
    long_description=read('README.md'),
    python_requires='>=3.6',
    classifiers=[
        "Development Status :: 5 - Production/Stable",
        "Topic :: Utilities"
    ],

    setup_requires=[
        'pytest-runner',
    ],
    tests_require=['pytest']
)
