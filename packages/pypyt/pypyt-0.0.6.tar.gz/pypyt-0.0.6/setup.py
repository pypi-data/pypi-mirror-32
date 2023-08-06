from setuptools import setup


setup(
    name='pypyt',
    version="0.0.6",
    author='Julio Gajardo',
    author_email='j.gajardo@criteo.com',
    description='Simple library to render ppt templates in python',
    license='MIT',
    url='https://gitlab.criteois.com/j.gajardo/pypyt',
    packages=['pypyt'],
    install_requires=['python-pptx'],
    package_data={
        'pypyt': ['html_docs/*', 'html_docs/_images/*', 'html_docs/_sources/*', 'html_docs/_static/*',
                  'html_docs/source/*', 'html_docs/_static/css/*', 'html_docs/_static/fonts/*',
                  'html_docs/_static/js/*']
    }
)

__author__ = 'Julio Gajardo'
