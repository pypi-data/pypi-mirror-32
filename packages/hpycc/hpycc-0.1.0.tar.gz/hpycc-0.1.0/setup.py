from setuptools import setup


setup(
    name='hpycc',
    version='0.1.0'
            '',
    description='Download THOR files, run ECL scripts and download their results.',
    url='https://github.com/OdinProAgrica/hpycc',
    author='Rob Mansfield',
    author_email='rob.mansfield@proagrica.com',
    classifiers=[
        'Development Status :: 4 - Beta',
        'License :: OSI Approved :: GNU General Public License v3 or later (GPLv3+)',
        'Intended Audience :: Developers',
        'Intended Audience :: Information Technology',
        'Intended Audience :: Science/Research',
        'Natural Language :: English',
        'Programming Language :: Python :: 3',
        'Topic :: System :: Distributed Computing',
        'Topic :: Utilities'
    ],
    keywords='HPCC ECL data access distributed ROXIE THOR',
    license='GNU GPLv3',
    packages=['hpycc'],
    zip_safe=False,
    install_requires=[
        'pandas',
        'requests', 'docker'
    ],
    project_urls={
        'Bug Reports': 'https://github.com/OdinProAgrica/hpycc/issues',
        'Source': 'https://github.com/OdinProAgrica/hpycc',
        'ReadTheDocs': 'http://hpycc.readthedocs.io'
    }
)
