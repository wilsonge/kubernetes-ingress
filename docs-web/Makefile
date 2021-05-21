NETLIFY_ROOT = /opt/build/repo
BUILDDIR = ${NETLIFY_ROOT}/docs-web/build
THEME_REPO = ${THEME_REPO_URL}

.PHONY: all build build-preview clean setup

all: clean setup build

clean:
	rm -rf ${SITE_ROOT}

setup:
	python -m pip install -r requirements.txt
	
preview-setup:
	cp ${NETLIFY_ROOT}/docs-web/_static/dev-home.txt ${NETLIFY_ROOT}/docs-web/index.rst

build:
	sphinx-build -b dirhtml -d ${BUILDDIR}/doctrees ${NETLIFY_ROOT}/docs-web ${BUILDDIR}/dirhtml
	@echo "Build finished. The HTML pages are in ${BUILDDIR}/dirhtml."

production: build

preview: preview-setup build