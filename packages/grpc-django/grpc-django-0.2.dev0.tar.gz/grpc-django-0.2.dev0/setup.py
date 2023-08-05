import codecs
try:
    codecs.lookup('mbcs')
except LookupError:
    ascii = codecs.lookup('ascii')
    func = lambda name, enc=ascii: {True: enc}.get(name=='mbcs')
    codecs.register(func)

from setuptools import find_packages, setup

version = "0.2dev"

setup(
    name="grpc-django",
    version=version,
    packages=find_packages(exclude=["tests*", "manage.py", "test_pb2.py", "test_pb2_grpc.py"]),
    include_package_data=True,
    install_requires=[
        "Django >= 1.9",
        "grpcio",
        "grpcio-tools", 'google', 'six'
    ],
    author="Sohel Tarir",
    author_email="sohel.tarir@gmail.com",
    description="GRPC integration with Django projects"
)