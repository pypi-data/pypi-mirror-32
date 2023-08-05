from setuptools import setup

setup(
    name='msk',
    version='0.1.3',  # Also update in msk/__init__.py
    packages=['msk'],
    install_requires=['GitPython', 'typing', 'msm>=0.5.13', 'pygithub'],
    url='https://github.com/MycroftAI/mycroft-skills-kit',
    license='MIT',
    author='Mycroft AI',
    author_email='support@mycroft.ai',
    maintainer='Matthew Scholefield',
    maintainer_email='matthew331199@gmail.com',
    description='Mycroft Skills Kit',
    entry_points={
        'console_scripts': {
            'msk=msk.__main__:main'
        }
    }
)
