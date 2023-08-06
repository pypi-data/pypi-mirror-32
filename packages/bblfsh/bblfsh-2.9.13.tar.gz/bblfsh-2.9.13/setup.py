import os
import sys

from setuptools import setup, find_packages, Extension
from setuptools.command.build_ext import build_ext

VERSION = "2.9.13"
LIBUAST_VERSION = "v1.9.1"
SDK_VERSION = "v1.8.0"
SDK_MAJOR = SDK_VERSION.split('.')[0]
PYTHON = "python3"

# For debugging libuast-client interactions, set to True in production!
GET_LIBUAST = True
if not GET_LIBUAST:
    print('WARNING: not retrieving libuast, using local version')

os.environ["CC"] = "g++"
os.environ["CXX"] = "g++"
libraries = ['xml2']
sources = ['bblfsh/pyuast.cc', 'bblfsh/memtracker.cc']


class CustomBuildExt(build_ext):
    def run(self):
        global libraries
        global sources

        if "--global-uast" in sys.argv:
            libraries.append('uast')
        else:
            sources.append('bblfsh/libuast/uast.cc')
            sources.append('bblfsh/libuast/roles.c')

        getLibuast()
        build_ext.run(self)


def runc(cmd):
    cmd = cmd.format(**globals())
    print(cmd)
    ret = os.system(cmd)
    if ret != 0:
        raise Exception("Command failed with output %d: %s" % (ret, cmd))


def createDirs():
    runc("mkdir -p gopkg.in/bblfsh/sdk.{SDK_MAJOR}/protocol")
    runc("mkdir -p gopkg.in/bblfsh/sdk.{SDK_MAJOR}/uast")
    runc("mkdir -p bblfsh/gopkg/in/bblfsh/sdk/{SDK_MAJOR}/protocol")
    runc("mkdir -p bblfsh/gopkg/in/bblfsh/sdk/{SDK_MAJOR}/uast")
    runc("mkdir -p bblfsh/github/com/gogo/protobuf/gogoproto")


def createInits():
    initFiles = [
            "bblfsh/github/__init__.py",
            "bblfsh/github/com/__init__.py",
            "bblfsh/github/com/gogo/__init__.py",
            "bblfsh/github/com/gogo/protobuf/__init__.py",
            "bblfsh/github/com/gogo/protobuf/gogoproto/__init__.py",
            "bblfsh/gopkg/__init__.py",
            "bblfsh/gopkg/in/__init__.py",
            "bblfsh/gopkg/in/bblfsh/__init__.py",
            "bblfsh/gopkg/in/bblfsh/sdk/__init__.py",
            "bblfsh/gopkg/in/bblfsh/sdk/{}/__init__.py".format(SDK_MAJOR),
            "bblfsh/gopkg/in/bblfsh/sdk/{}/uast/__init__.py".format(SDK_MAJOR),
            "bblfsh/gopkg/in/bblfsh/sdk/{}/protocol/__init__.py".format(SDK_MAJOR)
    ]

    for f in initFiles:
        open(f, 'w').close()


def getLibuast():
    if not GET_LIBUAST:
        return

    runc("curl -SL https://github.com/bblfsh/libuast/archive/{LIBUAST_VERSION}/"
         "{LIBUAST_VERSION}.tar.gz | tar xz")
    runc("mv {} libuast".format('libuast-' + LIBUAST_VERSION.replace('v', '')))
    runc("cp -a libuast/src bblfsh/libuast")
    runc("rm -rf libuast")


def protoDownload():
    runc("curl -SL https://github.com/bblfsh/sdk/archive/{SDK_VERSION}.tar.gz | tar xz")
    dir_ = "sdk-" + SDK_VERSION[1:]
    runc("cp %s/protocol/generated.proto gopkg.in/bblfsh/sdk.{SDK_MAJOR}/protocol/" % dir_)
    runc("cp %s/uast/generated.proto gopkg.in/bblfsh/sdk.{SDK_MAJOR}/uast/" % dir_)
    runc("rm -rf %s" % dir_)


def protoCompile():
    # SDK
    runc("{PYTHON} -m grpc.tools.protoc --python_out=bblfsh/gopkg/in/bblfsh/sdk/{SDK_MAJOR}/protocol "
         "--grpc_python_out=bblfsh/gopkg/in/bblfsh/sdk/{SDK_MAJOR}/protocol "
         "-I gopkg.in/bblfsh/sdk.{SDK_MAJOR}/protocol -I . "
         "gopkg.in/bblfsh/sdk.{SDK_MAJOR}/protocol/generated.proto")

    # UAST
    runc("protoc --python_out bblfsh github.com/gogo/protobuf/gogoproto/gogo.proto")
    runc("protoc --python_out bblfsh gopkg.in/bblfsh/sdk.{SDK_MAJOR}/uast/generated.proto")


def doGetDeps():
    getLibuast()
    createDirs()
    createInits()
    protoDownload()
    protoCompile()


def clean():
    runc("rm -rf gopkg.in")
    runc("rm -rf bblfsh/github")
    runc("rm -rf bblfsh/gopkg")
    if GET_LIBUAST:
        runc("rm -rf bblfsh/libuast")


def main():
    # The --global-uast flag allows to install the python driver using the installed uast library
    if "--getdeps" in sys.argv:
        doGetDeps()
        sys.exit(0)

    if "--clean" in sys.argv:
        clean()
        sys.exit(0)

    libuast_module = Extension(
        'bblfsh.pyuast',
        libraries=libraries,
        library_dirs=['/usr/lib', '/usr/local/lib'],
        extra_compile_args=['-std=c++11'],
        include_dirs=['bblfsh/libuast/', '/usr/local/include', '/usr/local/include/libxml2',
                      '/usr/include', '/usr/include/libxml2'], sources=sources)

    setup(
        cmdclass = {
            "build_ext": CustomBuildExt,
        },
        name="bblfsh",
        description="Fetches Universal Abstract Syntax Trees from Babelfish.",
        version=VERSION,
        license="Apache 2.0",
        author="source{d}",
        author_email="language-analysis@sourced.tech",
        url="https://github.com/bblfsh/client-python",
        download_url='https://github.com/bblfsh/client-python',
        packages=find_packages(),
        exclude=["bblfsh/test.py"],
        keywords=["babelfish", "uast"],
        install_requires=["grpcio==1.10.0", "grpcio-tools==1.10.0", "docker", "protobuf>=3.4.0"],
        package_data={"": ["LICENSE", "README.md"]},
        ext_modules=[libuast_module],
        classifiers=[
            "Development Status :: 3 - Alpha",
            "Environment :: Console",
            "Intended Audience :: Developers",
            "License :: OSI Approved :: Apache Software License",
            "Operating System :: POSIX",
            "Programming Language :: Python :: 3.4",
            "Programming Language :: Python :: 3.5",
            "Programming Language :: Python :: 3.6",
            "Topic :: Software Development :: Libraries"
        ]
    )


if __name__ == '__main__':
    main()
