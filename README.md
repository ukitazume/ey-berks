# ey_berkshelf

[![Build Status](https://travis-ci.org/ukitazume/ey-berks.svg)](https://travis-ci.org/ukitazume/ey-berks)

### Getting Start

```
cd my_cookbook
ey-berks init .
```


### Configration File

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

#### common attributes
- repo(requires): repository name
- path(requires): the library path in the repository
- srcpath(option. default: path) the recipe path of the repository 
- distpath(option, default: path) the path that is used for creating cookbook directory
- host(option, default: github.com) the remote repository hostname

#### [library]
 installed for `cookbooks/main/libraries`
 
#### [definition]
 installed for `cookbooks/main/libraries`
 
#### [[cookbook]]
 add your cookbook


```
ey-berks install
```


#### apply you own atributes file

...... thinking the specific.


#### search cookbooks

....... consider how to integrate it
