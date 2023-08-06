
import io
import setuptools

def readme():
  with io.open('README.md', encoding='utf8') as fp:
    return fp.read()

setuptools.setup(
  name = 'nr.types',
  version = '1.0.3',
  author = 'Niklas Rosenstein',
  author_email = 'rosensteinniklas@gmail.com',
  description = 'Better named tuples, sumtypes and useful map types.',
  long_description = readme(),
  long_description_content_type = 'text/markdown',
  url = 'https://github.com/NiklasRosenstein-Python/nr.types',
  license = 'MIT',
  packages = setuptools.find_packages('src'),
  package_dir = {'': 'src'}
)
