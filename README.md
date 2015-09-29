# ey_berkshelf

[![Build Status](https://travis-ci.org/ukitazume/ey-berks.svg)](https://travis-ci.org/ukitazume/ey-berks)

### Getting Start

```
cd my_cookbook
ey-berks init .
```

```
[library]
repo = "engineyard/ey-cloud-recipes"
path = "cookbooks/main/libraries"

[definition]
repo = "engineyard/ey-cloud-recipes"
path = "cookbooks/main/definitions"

[[cookbook]]
repo = "engineyard/ey-cloud-recipes"
path = "cookbooks/env_vars"

[[cookbook]]
host = "bitbucket.org"
repo = "engineyard/ey-cloud-recipes"
srcpath = "cookbooks/env_vars"
distpath = "cookbooks/env_vars_2"
```


```
ey-berks install
```
