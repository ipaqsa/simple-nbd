# Simple NBD server/client wrapper

## Overview
This is simple NBD server and client, wrapper over github.com/pojntfx/go-nbd

In the future, policy handler will be added.

## Use
### Note
Don`t forget to enable nbd kernel module.

### Create disk
You can create disk using qemu-img use following:
```
make new-disk path={PATH} disk={DISK NAME} size={SIZE}
```
### Following disks
- Via createMeta you can manual enter disks
- Via CreateMetaFromDir you can enter dir with disks
### Build

Build server use following:
```
make server-build
```
Build server use following:
```
make client-build
```

```
if you are using server with config, 
don`t forget flag -c with the path to config file
```

### Mount
After connect client to server, you may create filesystem on client.
```
mkfs.ext4 /path/to/nbd
```
Then you can mount that.
```
mount -t ext4 /path/to/nbd /your/path
```
