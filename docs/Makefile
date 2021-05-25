SHELL = /bin/bash
DOCS_THEME_REPO = https://gitlab.com/f5/nginx/controller/poc/f5-hugo
THEME_MODULE = gitlab.com/f5/nginx/controller/poc/f5-hugo
THEME_BRANCH = development

clean:
	rm -rf public

hugo build:
	hugo --gc -b $(DEPLOY_URL)

hugo-build-prod: 
	hugo --gc -e production
	
hugo-build-staging:
	hugo --gc -e staging

hugo-build-dev: 
	hugo --gc -e development

hugo-build-deploy-preview:
	hugo --gc -b $(DEPLOY_PRIME_URL)