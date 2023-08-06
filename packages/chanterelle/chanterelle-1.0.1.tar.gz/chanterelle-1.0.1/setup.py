import setuptools


def readme(path='README.rst', encoding='utf-8'):
    """Return the content of the README.rst file."""
    with open(path, 'r', encoding=encoding) as readme:
        return readme.read()


setuptools.setup(
    name='chanterelle',
    version='1.0.1',
    description="Utility for uploading static site files to an S3 bucket.",
    long_description=readme(),
    url='https://gitlab.com/unixispower/chanterelle',
    license='MIT',
    author='Blaine Murphy',
    author_email='myself@blaines.world',
    # See: https://pypi.python.org/pypi?:action=list_classifiers
    classifiers=[
        'Development Status :: 4 - Beta',
        'Environment :: Console',
        'Intended Audience :: Developers',
        'License :: OSI Approved :: MIT License',
        'Programming Language :: Python',
        'Programming Language :: Python :: 3 :: Only',
        'Topic :: Internet :: WWW/HTTP :: Site Management'
    ],

    py_modules=['chanterelle'],
    entry_points={
        'console_scripts': ['chanterelle = chanterelle:main']
    },
    install_requires=[
        'boto3',
        'botocore',
        'PyYAML'
    ],
    test_suite='tests'
)
