from setuptools import setup, find_packages

setup(
    name='depynd',
    version='0.4.4',
    description='Evaluating dependencies among random variables.',
    author='Yuya Takashina',
    author_email='takashina2051@gmail.com',
    packages=find_packages(),
    install_requires=[
        'numpy', 'scipy', 'scikit-learn',
    ],
    url='https://github.com/y-takashina/depynd',
)
