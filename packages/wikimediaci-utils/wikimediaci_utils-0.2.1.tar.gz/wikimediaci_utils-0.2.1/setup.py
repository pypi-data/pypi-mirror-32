from setuptools import setup

setup(
    name='wikimediaci_utils',
    version='0.2.1',
    packages=['wikimediaci_utils'],
    url='https://gerrit.wikimedia.org/g/integration/utils',
    license='GPL-3.0-or-later',
    author='Kunal Mehta',
    author_email='legoktm@member.fsf.org',
    description='Common utility functions for tools related to Wikimedia CI',
    long_description=open('README.rst').read(),
    long_description_content_type='text/x-rst',
    install_requires=[
        'pyyaml',
        'requests',
    ],
)
