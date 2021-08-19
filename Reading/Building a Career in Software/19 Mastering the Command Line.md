# Chapter 19: Mastering the Command Line

## Shell

* I/O redirection and pipes: `>`, `>>`, `<`, `|`, `>&`, `<<`.
* `cd`, `cd -`.
* Backgrounding with `&`.
* Killing and suspending with `ctrl-c` and `ctrl-z`; foregrounding with `fg`.
* Setting and suspending variables; quoting and escaping.
* Special variables: `$?`, `!$`, `!!`, `$#`, `%%`, `$!`, `$_`, `$0` through `$9`.
* Configuring your environment:
  * `PATH`, `.bashrc`, `.bash_profile`.
  * Aliases
* for and while loops.
* read into a variable.
* if-then-else, `[[]]`; file and string tests like `-z`, `-n`, and their many friends.
* Helpful shell config settings like `set -x`, `set -e`, `set -u`, `set -o` pipefail.
* Up arrow and `ctrl-r` for browsing history.
* `hash -r` and `hash -d` for managing cached paths to tools.
* Defining functions.
* Environment variables.
* `PIPESTATUS`
* `which`: see where a command comes from.
* `/dev/null`.
* String substituions and default values: `${foo:-bar}`, `${foo:=bar}`, `${foo#bar}`, and many more.

## Information Discovery and Text Manipulation

* `grep`: The original search tool. Emphasis on `-R`, `-1`, `-L`, `-E`, `-v`.
* `find`: Identify files, including complex boolean tests.
* `xargs`: For turning output of one command into arguments to the next.
* `ls`: Get familiar with at least `-R`, `-d`, `-t`, `-l`, `-1`.
* `cat`: Extract file content`.
* `sort` and `uniq`: e.g., `<do something> | sort | uniq -c #`.
* `awk`: at least enough to print columns: `awk '{print $1}'`.
* `wc`: how many words, characters, or lines.
* `less`, `head`, `tail` for quick glances a files.
* `sed` substring replacement.
* Regular expressions: Learn the basics, they pay off again and again.
* `column`: Align text into columns.
* `tr`: fast character mappings.

## Special Section for the Best Tool Ever

* `jq`.

## Networking

* `ping`: Confirm you can reach a host/IP.
* `dig` and `host`: DNS.
* `netstat`: list sockets on your host. `netstat -nap`.
* `curl`.
* `nc`: Ad hoc TCP.
* To example a host's IP networking: `ifconfig` on mac, `ipconfig` on Linux.

## Local and Remote Sessions

* `ssh`: Tunnel with `-L` and `-R`, and `-o` options. Learn to run both interactive sessions and one-shot commands.
* `scp` and `rsync` for copying files across machines.
* `tmux` or `screen`: Multiple shell in a single window and sessions you can rejoin on remote machines.

## Running Processes and Host State

* `ps`, `pgrep`, `top`, `pstree`, and `htop` for looking at running processes.
* `lsof` for looking at open files.
* `pkill`, `kill`, and `killall` for sending signals.
* `iostat`/`iotop` for I/O behavior.
* `df` and `du` for disk usage.
* `journalctl` and `systemctl` for logs and daemon state on Linux.
* `strace` for observing system calls on Linux.
* `/proc` for more process details on Linux.

## Databases

* The command line of whatever database you're using.

## Git

* `git cherry-pick`
* `git rebase`
* `git log`
* `git shortlog`
* `git tag`
* `git show`
* `git reset` and `git reset --hard`
* `git stash`
* `git status`
* `git rev-parse`

## Miscellany

* The ISO8601 standard for dates.
* `md5` and `sha1sum` for when you need a quick checksum.
* `uuidgen -v` for when you need a random UUID.
* If you use a Mac:
  * `pbcopy` and `pbpaste` for using the Mac clipboard in shell pipelines.
  * `open` for opening browser window from the shell.
  * `homebrew` (`brew`) for package management.
* `autojump` for quick look-up for command line navigation.
