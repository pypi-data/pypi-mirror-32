import setuptools


setuptools.setup(
    name="livedataframe",
    version="0.0.19",
    author="Chris Dimoff",
    author_email="chris@livedataframe.com",
    description="Client for the LiveDataFrame analysis tool",
    install_requires = ['requests', 'pandas', 'websocket-client', 'matplotlib', 'numpy'],
    url="http://www.livedataframe.com",
    packages=setuptools.find_packages(),
    classifiers=(
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ),
)
