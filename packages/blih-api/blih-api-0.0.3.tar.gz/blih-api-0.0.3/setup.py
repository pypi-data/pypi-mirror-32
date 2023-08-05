from setuptools import setup

setup(
    name="blih-api",
    packages=["blih_api"],
    entry_points={
        'console_scripts': [
            'blih-api = blih_api.__main__:main',
        ]
    },
    version='0.0.3',
    description="Blih API - Epitech Blih Library",
    long_description="Interact easily with the blih API",
    author="SakiiR SakiiR (@SakiiR)",
    url="https://github.com/SakiiR/blih-api",
    install_requires=[
        'requests',
    ],
)
