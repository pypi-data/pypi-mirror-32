# coding: utf-8

"""
    Fiddle Options Platform

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: 1.0
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


from setuptools import setup, find_packages  # noqa: H301

NAME = "swagger-client"
VERSION = "1.0.4"
# To install the library, run the following
#
# python setup.py install
#
# prerequisite: setuptools
# http://pypi.python.org/pypi/setuptools

REQUIRES = ["urllib3 >= 1.15", "six >= 1.10", "certifi", "python-dateutil"]

setup(
    name='FiddleOptions',
    version='1.0.4',
    description="Fiddle Options Platform",
    author_email="kovacicek@hotmail.com",
    url="https://github.com/kovacicek/fiddleoptions",
    # download_url='https://github.com/kovacicek/fiddleoptions/archive/1.0.1.tar.gz',
    keywords=["Swagger", "Fiddle Options Platform"],
    install_requires=REQUIRES,
    # packages=find_packages(),
    packages=['swagger_client', 'swagger_client_test'],
    include_package_data=True,
    long_description="""\
    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501
    """
)
