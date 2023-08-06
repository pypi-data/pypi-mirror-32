from setuptools import setup

setup(
    name='falkonry-cli',
    version='2.1.1',
    author='Falkonry Inc',
    author_email='info@falkonry.com',
    license='MIT',
    url='https://github.com/Falkonry/falkonry-cli',
    download_url = 'https://github.com/Falkonry/falkonry-cli/tarball/2.1.1',
    description='Cli tool to access Condition Prediction APIs',
    long_description='Falkonry cli tool to access Condition Prediction APIs',
    py_modules=['falkonry'],
    install_requires=[
        'cmd2==0.8.0',
        'pprint==0.1',
        'falkonryclient==2.1.1'
    ],
    entry_points='''
        [console_scripts]
        falkonry=falkonry:cli
    ''',
    zip_safe=False,
    include_package_data=True,
    keywords='falkonry falkonryclient falkonrycli'
)