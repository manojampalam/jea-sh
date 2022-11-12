A simple contrained custom shell to be used with remote ssh sessions via [ForceCommand](https://man.openbsd.org/sshd_config#ForceCommand)

## Usage
Set user's command to be executed via [SSH_ORIGINAL_COMMAND](https://man.openbsd.org/sshd_config#ForceCommand) environment variable. Pass a list of allowed commands that can be run as a base64 encoding of following yaml
```
allowedCommands:
 - whoami
 - ifconfig
```
Example 
```diff
SSH_ORIGINAL_COMMAND=whoami 
jea-sh YWxsb3dlZENvbW1hbmRzOgogLSB3aG9hbWkKIC0gaWZjb25maWc=
>    localuser

SSH_ORIGINAL_COMMAND=date
jea-sh YWxsb3dlZENvbW1hbmRzOgogLSB3aG9hbWkKIC0gaWZjb25maWc= 
-    unauthorized

```


