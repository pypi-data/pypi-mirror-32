import os

from setuptools import setup, find_packages


here = os.path.abspath(os.path.dirname(__file__))
packages = find_packages()

about = {}
with open(os.path.join(here, 'ictmpl', '__version__.py'), 'r') as f:
    exec(f.read(), about)


setup(
    name=about['__title__'],
    version=about['__version__'],
    description=about['__description__'],
    author=about['__author__'],
    author_email=about['__author_email__'],
    url=about['__url__'],
    zip_safe=False,
    packages=find_packages(),
    entry_points="""
    [console_scripts]
    ictmpl=ictmpl.app:run
    """
)