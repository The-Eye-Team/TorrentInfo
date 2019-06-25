# TorrentInfo
Get total size of a folder full of .torrent

# Usage

You can find releases for various operating systems in the [releases tab](https://github.com/The-Eye-Team/TorrentInfo/releases).

Download one, then make it executable:

```
chmod +x TorrentInfo
```

Sample usage with a folder called `torrents` with your torrent files inside:

```
./TorrentInfo -i torrents/
```

You can see the options with the `-h` flag:

```
TorrentInfo [-h|--help] -i|--input "<value>" [-j|--concurrency
                   <integer>]

                   Get infos of a folder full of .torrent

Arguments:

  -h  --help         Print help information
  -i  --input        Input directory
  -j  --concurrency  Concurrency. Default: 4
  ```
 
# Build

```
git clone https://github.com/The-Eye-Team/TorrentInfo.git && cd TorrentInfo
```

```
go get ./...
```

```
go build .
```

[![The-Eye.eu](https://the-eye.eu/public/.css/logo3_x300.png)](https://the-eye.eu)
