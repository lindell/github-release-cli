🔧 Usage
----

```yaml
before_deploy:
- curl https://github.com/lindell/travis-golang-release-boilerplate/releases/download/$TRAVIS_TAG/github-release -L --output github-releaser && chmod +x github-releaser
- export BODY=$(cat ./CHANGELOG.md)
- export FILES=bin/*
deploy:
  provider: script
  script: ./github-releaser
  skip_cleanup: true
  on:
    tags: true
```

📡 Release Notes
----
* Example
