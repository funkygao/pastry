# DHT
- load balance
- routing hops
- join/leave/die/reborn
- network locality
- moving keys/replica
- overlay robustness

# pastry

Pastry, a generic peer-to-peer object location and routing scheme, based on a 
self-organizing overlay network of nodes connected to the Internet. 

Pastry is completely decentralized, fault-resilient, scalable, and reliable. 
Moreover, Pastry has good route locality properties.

A DHT with geo- or network-locality: LRM

1M nodes, routing max hop: 5, R with 75 entries
1B nodes, routing max hop: 7, R with 105 entries

Routing table size: 32*16+2L(L=16) = 544


如果next hop在R里找不到，就从LRM里找那个节点，它满足:
1. 它与key的common prefix len >= 当前node和key的common prefix len
2. 它与key的差比当前nodeid-key更小，更近

LeafSet L:
- detect failures of their L by heartbeat
- may be used for replica

Neighborhood M:
- used for populating R with close-by nodes

## repair

- L, eager
- R, lazy, upon failure
- M, eager


sendHeartbeats() -> send()
sendStateTables(node, tables, eol) -> send()
sendRaceNotification(node, tables) -> send()
Send(msg) -> send()
send(msg, node) -> SendToIP(msg, hostPort)

# gossip

# P2P 历史

非结构化P2P中，每个节点存储自己的信息，当用户需要在P2P中获取信息时，他们预先不知道这些信息会在哪个节点上存储，因此
需要类似广播地进行查找

结构化P2P中，每个节点存储部分数据和特定信息的索引，搜索信息时他们知道信息可能存在于哪些节点

## napster 第一代P2P，集中目录模式
目录服务集中，文件交换在两个peers间直接进行

高峰期并发6M用户

            I have X!                        Who has X?
    Peer(A) ----------> central directory <---------------- Peer(B)
      |      1                               2                |
      |                                                       |
      +-------------------------------------------------------+
                        3
    
## Gnutella 第二代P2P，洪泛请求模式(非结构化P2P)
每个Peer的请求广播到连接的 Peers，各Peers又广播到各自连接的Peers，直到收到应答或达到5-9 hops
Gnutella采用了等级制的组成结构：节点被分成超级节点和普通节点，普通节点必须依附于超级节点。
每个超级节点作为一个独立的域管理者，负责处理域内的查询操作。在查找的过程中，查询首先在域内进行，失败后才会扩展到超级节点之间。


## BitTorrent
Tracker 
文件分块，chunk digest

## tapstry, pastry, chord, CAN 路由模式
目标是距离感知，减少hop
系统设计规模是10亿用户，10**14个文件

CARP(cache array routing protocol): 就是个mod based url hash，浏览器事先加载所有cache proxy地址，对
要访问的url进行hash并取模，通过1 hop找到目标proxy


Chord就是分布式二分查找

Pastry路由时，如果在LR都找不到路由节点，那么就从LM中找个proximity最小的进行转发。采用了路由器
上的最大掩码匹配算法

# skipnet

node name = com.wd.host1



    X                   A                       Z
    |                   |                       |
    | join              |                       |
    |------------------>|                       |
    |                   |  send(join)           |
    |                   |---------------------->|
    |                   |                       |
    |                   |                 state |
    |<------------------------------------------|
    |                   |                       |
    |             state |                       |
    |<------------------| X收到state后          |
    |                   | LRM里就有了A
    |                   |
    | heartbeat         | to calculate proximity
    |------------------>|
    |                   |
    |                   |
    | announce          | 
    |------------------>|-------------> all nodes in LRM
    |                   |
    |                   |
    |                   |
