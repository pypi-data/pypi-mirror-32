# pypkgdemo1

This is a simple example package. You can use 
[Github-flavored Markdown](https://guides.github.com/features/mastering-markdown/)
to write your content.

This page explain about packging python projects and upload to pypi.
https://packaging.python.org/tutorials/packaging-projects/


## build
pip install --upgrade pip setuptools wheel

pytyhon setup.py sdist bdist_wheel

ls dist/



## upload to TestPypi

pip install --upgrade twine

twine upload --repository -url https://test.pypi.org/legacy/ dist/*

## upload to TestPypi

## install from TestPypi

pip install --index-url https://test.pypi.org/simple/ pypkgdemo1

## testing pypkgdemo1

python

import pypkgdemo1

pypkgdemo1.joke()

## uninstall pypkgdemo1

pip uninstall pypkgdemo1


## upload to Pypi

pip install --upgrade twine

twine upload dist/*

## install from Pypi

pip install pypkgdemo1


