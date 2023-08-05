import setuptools

setuptools.setup(
    name='otcbtc_client',
    version='0.1.0',
    author='Gimo',
    author_email='self@gimo.me',
    description='OTCBTC SDK for Python',
    url='https://github.com/masakichi/otcbtc-client',
    packages=setuptools.find_packages(exclude=['tests']),
    classifiers=('Intended Audience :: Developers',
                 'License :: OSI Approved :: MIT License',
                 'Operating System :: OS Independent',
                 'Programming Language :: Python :: 2',
                 'Programming Language :: Python :: 2.7',
                 'Programming Language :: Python :: 3',
                 'Programming Language :: Python :: 3.4',
                 'Programming Language :: Python :: 3.5',
                 'Programming Language :: Python :: 3.6',
                 'Programming Language :: Python :: Implementation :: CPython',
                 'Programming Language :: Python :: Implementation :: PyPy'),
    install_requires=[
        'requests >=2.11, <3.0a0',
        'future',
    ])
