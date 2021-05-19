# Configuration file for the Sphinx documentation builder.
#
# This file only contains a selection of the most common options. For a full
# list see the documentation:
# https://www.sphinx-doc.org/en/master/usage/configuration.html

# -- Path setup --------------------------------------------------------------

# If extensions (or modules to document with autodoc) are in another directory,
# add these directories to sys.path here. If the directory is relative to the
# documentation root, use os.path.abspath to make it absolute, like shown here.
#
import os
import sys
sys.path.insert(0, os.path.abspath('.'))
sys.path.insert(0, os.path.abspath('..'))
import nginx_sphinx

# -- Project information -----------------------------------------------------

project = 'NGINX Ingress Controller Docs'
copyright = '2021, NGINX, an F5 company'
author = 'NGINX, an F5 company'


# -- General configuration ---------------------------------------------------

# Add any Sphinx extension module names here, as strings. They can be
# extensions coming with Sphinx (named 'sphinx.ext.*') or your custom
# ones.
extensions = [
  'myst_parser',
  'sphinx_sitemap'
]

# Add any paths that contain templates here, relative to this directory.
templates_path = ['_templates']

# The root document.
root_doc = 'index'

source_suffix = {
            '.rst': 'restructuredtext',
            '.txt': 'markdown',
            '.md': 'markdown',
       }

# List of patterns, relative to source directory, that match files and
# directories to ignore when looking for source files.
# This pattern also affects html_static_path and html_extra_path.
exclude_patterns = ['_build', 'Thumbs.db', '.DS_Store', '_source']


# -- Options for HTML output -------------------------------------------------

# The theme to use for HTML and HTML Help pages.  See the documentation for
# a list of builtin themes.
#
html_theme = 'nginx_sphinx'

# Add any paths that contain custom themes here, relative to this directory.
html_theme_path = nginx_sphinx.get_html_theme_path()

html_baseurl = 'https://docs.nginx.com/'

html_theme_options = {
  #'base_url': html_baseurl    
}

# Add any paths that contain custom static files (such as style sheets) here,
# relative to this directory. They are copied after the builtin static files,
# so a file named "default.css" will overwrite the builtin "default.css".
html_static_path = ['_static']

html_title = "NGINX Ingress Controller Docs"

html_last_updated_fmt = '%b %d, %Y'