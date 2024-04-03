from setuptools import setup, find_packages

setup(
    name="shared",
    packages=find_packages(include=["*"]),
    install_requires=[
        'grpcio==1.38.1',
        'grpcio-tools==1.38.1',
    ],
    package_data={'': ['*.proto']},
    include_package_data=True,
)
