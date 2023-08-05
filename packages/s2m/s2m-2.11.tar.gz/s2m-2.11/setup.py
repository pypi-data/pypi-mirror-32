from setuptools import setup

with open('long_desc.md') as f:
    long_description = f.read()

setup(
    name='s2m',
    version='2.11',
    packages=['s2m'],
    install_requires=['pyserial>=2.7',
                      'psutil'],
    package_data={'s2m': [('scratch_files/extensions/*.s2e'),
                          ('scratch_files/projects/*.sb2'),
                          ('micro_bit_scripts/*.py')]},
    entry_points={
        'console_scripts': [
            's2m = s2m.s2m:main'
        ]
    },
    url='https://github.com/MrYsLab/s2m',
    download_url='https://github.com/MrYsLab/s2m',
    license='GNU Affero General Public License v3 (AGPLv3+)',
    author='Alan Yorinks',
    author_email='MisterYsLab@gmail.com',
    description='A Scratch 2.0 (Offline) Hardware Extension for micro:bit',
    long_description=long_description,
    long_description_content_type='text/markdown',
    keywords=['micro:bit', 'microbit', 'Scratch'],
    classifiers=[
        'Development Status :: 5 - Production/Stable',
        'Environment :: Other Environment',
        'Intended Audience :: Education',
        'License :: OSI Approved :: GNU Affero General Public License v3 or later (AGPLv3+)',
        'Operating System :: OS Independent',
        'Programming Language :: Python :: 2.7',
        'Programming Language :: Python :: 3.4',
        'Programming Language :: Python :: 3.5',
        'Programming Language :: Python :: 3.6',
        'Topic :: Education',
    ],
)

