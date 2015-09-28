# ey_berkshelf

[![Build Status](https://travis-ci.org/ukitazume/ey-berks.svg)](https://travis-ci.org/ukitazume/ey-berks)

### Getting Start

```
cd my_cookbook
eyberks init .
```

```
[main]
libraries = "engineyard/ey-cloud-recipes/main/libraries"
definitions = "engineyard/ey-cloud-recipes/main/definitions"

[cookbook]
sidekiq = "engineyard/ey-cloud-recipes/cookbooks/sidekiq"
fluentd = "ukitazume/ey-mini-recipes/cookbooks/fluentd"
```


```
eyberks install
```

