# Remote executor

Launch arbitrary commands to remote hosts via different transports.
Was designed with concurrency in mind for better throughput.
By now the only transport implemented is SSH.

The input for this command is via stdin, piping the content of a hosts file. 
The content of a host file needs to be:

``` bash
example.com:22
example2.com:2222
```
Note that the port is mandatory.

It will catch your default private key for connections as well the current user.

## Example

```bash
# Executing a command with all params by default
cat hosts.lst | remote-executor uptime
# output
host1 says :
 21:14:59 up 48 days,  1:15,  0 users,  load average: 0,64, 0,47, 0,42
host2 says :
 21:14:28 up 63 days,  3:11,  0 users,  load average: 0.48, 0.10, 0.03

```
## Advanced examples
```bash
## With different private key, current user.
cat hosts.lst | remote-executor -k /keys/custom_private_key uptime
```
```bash
## With different private key and another user than current.
cat hosts.lst | remote-executor -k /keys/custom_private_key -u otheruser uptime
```
```bash
## Changing concurrency levels to 40 concurrent connections, default is 10.
cat hosts.lst | remote-executor -c 40 uptime
```