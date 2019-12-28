# [THIS PROJECT IS NO LONGER MAINTAINED SINCE THE ORIGINAL PROBLEM IS NOW FIXED](https://github.com/travis-ci/dpl/pull/1069)

# 🚢 github-release-cli

Github release CLI is a tool primarily created to circumvent the Travis issue that doesn't allow [complex bodies to be used in releases](https://github.com/travis-ci/dpl/issues/155).

🔧 Usage
----
The github-travis-releaser takes enviroment variables from travis and parses them so that you don't have to.

#### The only variables you need to take into concideration are:
| Name | Optional | Description |
|------|----------|-------------|
| GITHUB_OAUTH_TOKEN | NO | The oath token from github, needed to get the access to create the release |
| BODY | YES | The body of the release, can be markdown. It's recomended to use `$(envsubst < ./CHANGELOG.md)` or `$(cat ./CHANGELOG.md)` for longer bodies |
| FILES | YES | The path to the file(s) that should be uploaded. Wildcards can be used (e.g. `release-files/*`) |
| RELEASE_NAME | YES | The name (title) of the release. If nothing is set, it will be the same as the tag name |

#### Command line flags:
| Name | Description |
|------|-------------|
| -draft | Set the release as a draft |
| -prerelease | Set if the the release is identified as non-production ready |
| -verbose | Print logging statements |

### Example `.travis.yml`
```yaml
before_deploy:
- curl https://github.com/lindell/github-release-cli/releases/download/LATEST_RELEASE/github-releaser-travis -L --output github-releaser && chmod +x github-releaser
- export BODY=$(envsubst < ./CHANGELOG.md)
- export FILES=release-files/*
deploy:
  provider: script
  script: ./github-releaser -draft -verbose
  skip_cleanup: true
  on:
    tags: true
```
