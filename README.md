# ey_berkshelf

[![Build Status](https://travis-ci.org/ukitazume/ey-berks.svg)](https://travis-ci.org/ukitazume/ey-berks)


Engine Yard Custom Chef tool like Berfshef.

![ScreenCast](http://showterm.io/a1a9b260062666a5238a6)

[![ScreenShot](https://raw.github.com/GabLeRoux/WebMole/master/ressources/WebMole_Youtube_Video.png)](http://showterm.io/a1a9b260062666a5238a6)

### Getting Start

```
ey-berks config .
ey-berks compile .
```

### Usage

```
Engine Yard Cloud cookbook tool like Berkshelf

Usage: ey-berks <command> [<path>] [--config=<config>]

ey-berks config <path> [--config=<config>]             : make a sample configuration file
ey-berks compile <path> [--config=<config>]            : update cahce,  write a main/recipes and gather recipe to the cookbooks directory
ey-berks update-cache [--config=<config>]              : update cache of remote repositories cookbooks
ey-berks create-main-recipe <path> [--config=<config>]
ey-berks copy-recipes <path> [--config=<config>]
ey-berks help
ey-berks version
```

### Configration File

use TOML for the configuration like that

```
[library]
repo = "engineyard/ey-cloud-recipes"
path = "cookbooks/main/libraries"
name = "main/libraries"

[definition]
repo = "engineyard/ey-cloud-recipes"
path = "cookbooks/main/definitions"
name = "main/definitions"

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
- name(option) the path that is used for creating cookbook directory
- host(option, default: github.com) the remote repository hostname


#### apply you own atributes file

...... thinking the specific.


#### search cookbooks

....... consider how to integrate it
