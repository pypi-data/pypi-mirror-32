from setuptools import setup, find_packages

__version__ = '1.0.2'


if __name__ == "__main__":
    setup(
        name='vakt',
        description='Attribute-based access control (ABAC) SDK for Python',
        keywords='ACL RBAC access policy security',
        version=__version__,
        author='Egor Kolotaev',
        author_email='ekolotaev@gmail.com',
        license="Apache 2.0 license",
        url='https://github.com/kolotaev/vakt',
        py_modules=['vakt'],
        install_requires=[],
        extras_require={
            'dev': [
                'pytest',
                'pytest-cov',
                'pylint',
            ]
        },
        packages=find_packages(exclude='tests'),
        classifiers=[
            'Intended Audience :: Developers',
            'License :: OSI Approved :: Apache Software License',
            'Operating System :: OS Independent',
            'Programming Language :: Python',
            'Programming Language :: Python :: 3.3',
            'Programming Language :: Python :: 3.4',
            'Programming Language :: Python :: 3.5',
            'Programming Language :: Python :: 3.6',
            'Programming Language :: Python :: 3.7',
            'Programming Language :: Python :: Implementation :: PyPy',
        ],
    )
