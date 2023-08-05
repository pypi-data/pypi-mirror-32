from setuptools import setup


setup(
    name="bundlewrap-teamvault",
    version="2.1.0",
    description="Access TeamVault secrets from BundleWrap",
    author="Torsten Rehn",
    author_email="torsten@rehn.email",
    license="GPLv3",
    py_modules=['bwtv'],
    keywords=["configuration", "config", "management"],
    classifiers=[
        "Development Status :: 5 - Production/Stable",
        "Intended Audience :: System Administrators",
        "License :: OSI Approved :: GNU General Public License v3 (GPLv3)",
        "Natural Language :: English",
        "Operating System :: POSIX :: Linux",
        "Programming Language :: Python",
        "Programming Language :: Python :: 2.7",
        "Programming Language :: Python :: 3.3",
        "Programming Language :: Python :: 3.4",
        "Topic :: Software Development :: Libraries :: Python Modules",
        "Topic :: System :: Installation/Setup",
        "Topic :: System :: Systems Administration",
    ],
    install_requires=[
        "bundlewrap >= 2.3.1",
        "passlib",
        "requests >= 1.0.0",
    ],
)
