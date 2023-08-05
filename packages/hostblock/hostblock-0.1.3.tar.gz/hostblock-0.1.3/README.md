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

Importing
-------------------------

To import a list of hosts, just pass it as argument to `ab` or `aw` command. Here is example on how to import `http://someonewhocares.org/hosts` list into your blacklist:

```bash
curl http://someonewhocares.org/hosts/zero/hosts | grep '^0\.0\.0\.0\s' | awk '{print $2}' | xargs hostblock ab
```

Merging
-------------------------

To merge your friend's whitelist and blacklist, see the following example:

```bash
hostblock --config friends.hostblock lb | xargs hostblock ab
hostblock --config friends.hostblock lw | xargs hostblock aw
```

Notes
-------------------------

- **Do not use sudo** when running hostblock commands. Hostblock will ask for sudo password when required to modify hosts file. If you use `sudo` then hostblock will think you are root user and will look for settings file in `/root/.hostblock` thinking that's your home dir. If you really need to use sudo, make sure to specify hostblock config with `--config` option.
- When adding duplicate domains to your blacklist or whitelist, hostblock will keep only single entry, so no worries!


