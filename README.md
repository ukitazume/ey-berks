# ey-berks

[![Build Status](https://travis-ci.org/ukitazume/ey-berks.svg)](https://travis-ci.org/ukitazume/ey-berks)

<img src="https://i.gyazo.com/b0a8e251dba3bd35edf4a3c3c26876b2.png" width="600px">


Engine Yard Custom Chef tool like Berfshef.

ScreenCast http://showterm.io/c8df12c47ff6391ca2a2c

### Basic Usage

```
$ey-berks config .
$ey-berks compile .
$ey recipes upload -e env_name
```

### Usage

```
$ ey-berks help
Engine Yard Cloud cookbook tool like Berkshelf

Usage:
	ey-berks config <path>                                               : make a sample configuration file (default path=$PWD, --config=EyBerks)
	ey-berks compile <path> --config=<path to EyBerks>                   : update cahce,  write a main/recipes and gather recipe to the cookbooks directory
	ey-berks update-cache                                                : update cache of remote repositories cookbooks
	ey-berks create-main-recipe <path>                                   : create main recipes from the configration file
	ey-berks copy-recipes <path>                                         : copy recipes from the cache dir to the cookbooks/ directory
	ey-berks clear <path>                                                : remove EyBerksfile and cookbooks directory
	ey-berks gather-attr <path> --from=</path/cookbooks>                 : gather attbiutes from cookbook directory
	ey-berks apply-attr <path> ---from=<attributes directory>            : apply attbiutes for cookbook directory
	ey-berks help                                                        : show this help
	ey-berks version                                                     : show the version
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

```
$ ey-berks gather-attr . --from=./cookbooks
$ ey-berks apply-attr . --from=./attr-meta
```

#### search cookbooks

....... consider how to integrate it
