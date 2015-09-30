# ey_berkshelf

[![Build Status](https://travis-ci.org/ukitazume/ey-berks.svg)](https://travis-ci.org/ukitazume/ey-berks)

### Getting Start

```
cd my_cookbook
ey-berks init .
```


### Configration File

use TOML for the configuration like that

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
path = "cookbooks/cutom_nginx"
```

This configuration creates the following cookbook/ directory

```
☺  tree ./cookbooks/
gather/cookbooks/
├── custom_nginx
│   ├── recipes
│   └── templates
├── env_vars
│   ├── README.md
│   ├── attributes
│   ├── recipes
│   └── templates
└── main
    ├── definitions
    └── libraries
```

### Group

#### [library]
 installed for `cookbooks/main/libraries`
 
#### [definition]
 installed for `cookbooks/main/libraries`
 
#### [[cookbook]]
 add your cookbooks

### Attributes

- repo(requires): repository name
- path(requires): the library path in the repository
- srcpath(option. default: path) the recipe path of the repository 
- distpath(option, default: path) the path that is used for creating cookbook directory
- host(option, default: github.com) the remote repository hostname



```
ey-berks install
```


#### apply you own atributes file

...... thinking the specific.


#### search cookbooks

....... consider how to integrate it
