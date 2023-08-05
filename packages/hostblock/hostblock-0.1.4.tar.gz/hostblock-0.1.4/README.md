
hostblock
==================================================
[![Build Status](https://travis-ci.org/nul-one/hostblock.png)](https://travis-ci.org/nul-one/hostblock)
[![PyPI version](https://badge.fury.io/py/hostblock.svg)](https://badge.fury.io/py/hostblock)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![Requirements Status](https://requires.io/github/nul-one/hostblock/requirements.svg?branch=master)](https://requires.io/github/nul-one/hostblock/requirements/?branch=master)

Block domains using hosts file entries.


Installation
-------------------------

### install from pypi (recommend)
`pip3 install hostblock`

### install from github (latest master)
`pip3 install -U git+https://github.com/nul-one/hostblock.git`


Usage
-------------------------

Run `hostblock apply` to add list of undesired hosts to your `/etc/hosts` file resolving to `0.0.0.0`. It will ask for sudo password.

### Local blacklist and whitelist

Hostblock will automatically create a json file `~/.hostblock` where it keeps your local blacklist and whitelist. When using `hostblock apply` your local whitelist will be subtracted from your blacklist and the result will be then put into your `/etc/hosts`.

Here is how you can control your local list:

- `hostblock ab HOST1 HOST2 ...` this will add single or multiple hosts to your blacklist
- `hostblock rb HOST1 HOST2 ...` remove single or multiple hosts from your blacklist
- `hostblock cb` remove (clear) all hosts from your blacklist
- `hostblock lb` list all hosts currently in your blacklist
- `hostblock aw HOST1 HOST2 ...` this will add single or multiple hosts to your whitelist
- `hostblock rw HOST1 HOST2 ...` remove single or multiple hosts from your whitelist
- `hostblock cw` remove (clear) all hosts from your whitelist
- `hostblock lw` list all hosts currently in your whitelist
- `hostblock list` list hosts from blacklist that do not appear in whitelist
- `hostblock count` show counts for blacklist and whitelist

Commands that don't change config files such as `lw`, `lb`, `list` and `count` can be followed by a list of configuration files. In that case default config file or `--config` option will be ignored and unified data will be displayed. Config files will not be touched in either case:

```bash
hostblock count first.config second.config third.config
```

The `apply` option can be followed by list of configs in similar manner. If so, unified config data will be applied to `/etc/hosts` or other selected hostname:


```bash
hostblock apply .hostblock first.config second.config third.config
```

In example above, we listed the default `.hostblock` config file. If we didn't, it wouldn't be used.


Importing
-------------------------

To import a list of hosts, just pass it as argument to `ab` or `aw` command.  
Here is an example bash script that will import lists from `http://someonewhocares.org` and `https://github.com/notracking/hosts-blocklists` lists into blacklist of multiple config files and then apply those including your default one to your hosts file:

```bash
#!/bin/bash

set -e

echo "Updating 'notracking/hosts'..."
curl -s https://raw.githubusercontent.com/notracking/hosts-blocklists/master/domains.txt \
    | fdump -p '^address=/(.*)/.*$' '{0}' \
    | xargs hostblock -c ~/.hostblock.notracking ab

echo "Updating 'notracking/domains'..."
curl -s https://raw.githubusercontent.com/notracking/hosts-blocklists/master/domains.txt \
    | fdump -p '^address=/(.*)/.*$' '{0}' \
    | xargs hostblock -c ~/.hostblock.notracking ab

echo "Updating 'someonewhocares/hosts'..."
curl -s http://someonewhocares.org/hosts/zero/hosts \
    | fdump -d "130.211.230.53" -p '^0\.0\.0\.0 ([a-zA-Z0-9\._-]+)\s.*' '{0}' \
    | xargs hostblock -c .hostblock.someonewhocares ab

hostblock count ~/.hostblock*

echo "Applying to /etc/hosts..."
hostblock apply ~/.hostblock*
```

The above example uses [fdump](https://github.com/nul-one/fdump) tool for easy filtering. You may install the tool with `pip3 install fdump` or use grep/awk alternatives in it's place.  
I use the above script as a daily cron-job. I only add new things manually into my default config file to keep it separate.


Merging
-------------------------

To merge your friend's whitelist and blacklist, see the following example:

```bash
hostblock --config friends.hostblock lb | xargs hostblock ab
hostblock --config friends.hostblock lw | xargs hostblock aw
```


Notes
-------------------------

- **Do not use sudo** when running hostblock commands. Hostblock will ask for sudo password when required to modify hosts file. If you use `sudo` then hostblock will think you are root user and will look for settings file in `/root/.hostblock` thinking that's your home dir. If you really need to use sudo, make sure to specify proper hostblock config with `--config` option.
- When adding duplicate domains to your blacklist or whitelist, hostblock will keep only single entry, so no worries!


