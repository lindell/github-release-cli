🔧 Usage
----
The github-travis-releaser takes enviroment variables from travis and parses them so that you don't have to.
The only variables you need to take into concideration are:

| Name | Optional | Description |
|------|----------|-------------|
| GITHUB_OATH_TOKEN | `true` | The oath token from github, needed to get the access to create the release |
| BODY | `false` | The body of the release, can be markdown. It's recomended to use `$(envsubst < ./CHANGELOG.md)` or `$(cat ./CHANGELOG.md)` for longer bodies |
| FILES | `false` | The path to the file(s) that should be uploaded. Wildcards can be used (e.g. `release-files/*`) |
| RELEASE_NAME | `false` | The name (title) of the release. If nothing is set, it will be the same as the tag name |

### Example `.travis.yml`
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
