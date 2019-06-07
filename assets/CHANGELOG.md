🔧 Usage
----

```yaml
before_deploy:
- curl https://github.com/lindell/github-release-cli/releases/download/$TRAVIS_TAG/github-releaser-travis -L --output github-releaser && chmod +x github-releaser
- export BODY=$(envsubst < ./CHANGELOG.md)
- export FILES=release-files/*
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
