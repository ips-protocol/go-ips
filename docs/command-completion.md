Command Completion
==================

Shell command completion is provided by the script at 
[/misc/completion/ipws-completion.bash](../misc/completion/ipws-completion.bash).


Installation
------------
The simplest way to see it working is to run 
`source misc/completion/ipws-completion.bash` straight from your shell. This
is only temporary and to fully enable it, you'll have to follow one of the steps
below.

### Bash on Linux
For bash, completion can be enabled in a couple of ways. One is to copy the 
completion script to the directory `~/.ipws/` and then in the file 
`~/.bash_completion` add
```bash
source ~/.ipws/ipws-completion.bash
```
It will automatically be loaded the next time bash is loaded.
To enable ipws command completion globally on your system you may also 
copy the completion script to `/etc/bash_completion.d/`.


Additional References
---------------------
* https://www.debian-administration.org/article/316/An_introduction_to_bash_completion_part_1
