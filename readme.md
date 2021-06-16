# Reddrift

<!--- mdtoc: toc begin -->

1.	[Synopsis](#synopsis)
2.	[Why?](#why-)
3.	[Usage](#usage)<!--- mdtoc: toc end -->

## Synopsis

Redrift is a [redshift](https://github.com/jonls/redshift.git) replacement. Currently a simple one but it does the job. It started as a fork of [https://github.com/dim13/sct](https://github.com/dim13/sct) but has developed quite a bit away recently.

## Why?

Because it can be build as static binary and eases the pain of building and installing [redshift](https://github.com/jonls/redshift.git). Copy and let run.

## Usage

```shell
# set certain colour temp
$ reddrift 4500

# use a preset
$ reddrift candle

# auto set based on a location
$ reddrift -c moscow

# keep running and continuously auto set
$ reddrift -c tokyo -r

# list available presets or locations
$ reddrift --list-presets
$ reddrift --list-locations

# help
$ reddrift -h
```
