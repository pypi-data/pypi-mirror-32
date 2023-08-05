from setuptools import setup


setup(
    name='theobserver',
    packages=['theobserver'],
    version='3.0',
    description='A dataset characteristic extractor for machine learning processing.',
    author='Bernardo Trevizan',
    author_email='trevizanbernardo@gmail.com',
    url='https://github.com/btrevizan/theobserver',
    keywords='feature,characteristic extration,machine learning',
    install_requires=['pandas==0.23', 'scikit-learn==0.19.0']
)
