from setuptools import setup

setup(
    name='nginx-amplify-agent-health-check',
    version='0.1.1',
    description='Static and Dynamic Analysis for nginx-amplify-agent Health Status',
    url='http://github.com/hiradyazdan/nginx-amplify-agent-health-check',
    author='Hirad Yazdanpanah',
    author_email='hirad.y@gmail.com',
    license='MIT',
    packages=['amplifyhealthcheck'],
    install_requires=[
        'psutil',
        'setuptools',
        'ntplib',
        'crossplane',
        'requests'
    ],
    zip_safe=False
)
