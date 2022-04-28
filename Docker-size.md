# super-app:scratch

```bash
┃ ● Layers ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ │ Current Layer Contents ├───────────────────────────────────
Cmp   Size  Command                                           Permission     UID:GID       Size  Filetree
    6.1 MB  FROM 86384c55d151ae9                              -rwxr-xr-x         0:0     6.1 MB  └── app
    200 kB  #(nop) COPY file:6ef523fe0b4d4f1682b1f576b40cde70

│ Layer Details ├────────────────────────────────────────────

Tags:   (unavailable)
Id:     86384c55d151ae91165cf7cd45d3d2515d0621bd97d1089f1fdc5
24cd2374c92
Digest: sha256:b770403da4d8d6ded039637e91997a4bba62a1a4115740
f417742ea23e56337e
Command:
#(nop) COPY file:9c789743e09b99ca4ac2f976b3685ce5e42937c8f17b
2b1a1b4a64c5a125ec7f in /

│ Image Details ├────────────────────────────────────────────


Total Image size: 6.3 MB
Potential wasted space: 0 B
Image efficiency score: 100 %

Count   Total Space  Path

```

# super-app:distroless

```bash
┃ ● Layers ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ │ Current Layer Contents ├───────────────────────────────────
Cmp   Size  Command                                           Permission     UID:GID       Size  Filetree
    6.1 MB  FROM 86384c55d151ae9                              -rwxr-xr-x         0:0     6.1 MB  └── app
    200 kB  #(nop) COPY file:6ef523fe0b4d4f1682b1f576b40cde70

│ Layer Details ├────────────────────────────────────────────

Tags:   (unavailable)
Id:     86384c55d151ae91165cf7cd45d3d2515d0621bd97d1089f1fdc5
24cd2374c92
Digest: sha256:b770403da4d8d6ded039637e91997a4bba62a1a4115740
f417742ea23e56337e
Command:
#(nop) COPY file:9c789743e09b99ca4ac2f976b3685ce5e42937c8f17b
2b1a1b4a64c5a125ec7f in /

│ Image Details ├────────────────────────────────────────────


Total Image size: 6.3 MB
Potential wasted space: 0 B
Image efficiency score: 100 %

Count   Total Space  Path

```


# 
```bash
┃ ● Layers ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ │ Current Layer Contents ├───────────────────────────────────
Cmp   Size  Command                                           Permission     UID:GID       Size  Filetree
    2.4 MB  FROM 2571f63daa6afb6                              -rwxr-xr-x         0:0     6.1 MB  ├── app
    6.1 MB  #(nop) COPY file:9c789743e09b99ca4ac2f976b3685ce5 drwxr-xr-x         0:0        0 B  ├── bin 
                                                              drwxr-xr-x         0:0        0 B  ├── boot
│ Layer Details ├──────────────────────────────────────────── drwxr-xr-x         0:0        0 B  ├── dev
                                                              drwxr-xr-x         0:0     220 kB  ├── etc               
Tags:   (unavailable)                                         -rw-r--r--         0:0        5 B  │   ├── debian_version
Id:     587b3e17271ee942144eb6a96442a926cfa21fa52bd32d6214a4b drwxr-xr-x         0:0        0 B  │   ├── default
ad65ad1b7cc                                                   drwxr-xr-x         0:0       83 B  │   ├── dpkg       
Digest: sha256:b770403da4d8d6ded039637e91997a4bba62a1a4115740 drwxr-xr-x         0:0       83 B  │   │   └── origins   
f417742ea23e56337e                                            -rw-r--r--         0:0       83 B  │   │       └── debian
Command:                                                      -rw-r--r--         0:0     1.8 kB  │   ├── ethertypes
#(nop) COPY file:9c789743e09b99ca4ac2f976b3685ce5e42937c8f17b -rw-r--r--         0:0       64 B  │   ├── group    
2b1a1b4a64c5a125ec7f in /                                     -rw-r--r--         0:0        9 B  │   ├── host.conf
                                                              -rw-r--r--         0:0       27 B  │   ├── issue    
│ Image Details ├──────────────────────────────────────────── -rw-r--r--         0:0       20 B  │   ├── issue.net    
                                                              -rw-r--r--         0:0      497 B  │   ├── nsswitch.conf     
                                                              -rwxrwxrwx         0:0        0 B  │   ├── os-release → ../us
Total Image size: 8.4 MB                                      -rw-r--r--         0:0      149 B  │   ├── passwd   
Potential wasted space: 0 B                                   drwxr-xr-x         0:0        0 B  │   ├── profile.d
Image efficiency score: 100 %                                 -rw-r--r--         0:0     2.9 kB  │   ├── protocols
                                                              -rw-r--r--         0:0      887 B  │   ├── rpc     
Count   Total Space  Path                                     -rw-r--r--         0:0      13 kB  │   ├── services
```