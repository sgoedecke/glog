# glog

Glog will SSH onto a log server for you, via a proxy host, and grep some files for a pattern. It will automatically read your SSH auth from ssh-agent and your server names from `.ssh/config`.

## Installation

The checked-in `glog` binary should work on MacOS.

## Usage

`glog logserveralias myapp*/log.log thing_to_grep_for`

## TODO

Make it work when you don't need a proxy host

Better error handling and argument parsing, with a help page

Allow grepping over multiple hosts

Config file to store common arguments
