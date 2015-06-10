# DHT
- load balance
- routing hops
- join/leave/die/reborn
- network locality
- moving keys/replica
- overlay robustness

# pastry

A DHT with geo- or network-locality: LRM

1M nodes, routing max hop: 5

Routing table size: 32*16+2L(L=16) = 544

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
## napster 第一代P2P，集中目录模式
目录服务集中，文件交换在两个peers间直接进行

高峰期并发6M用户

            I have X!                        Who has X?
    Peer(A) ----------> central directory <---------------- Peer(B)
      |      1                               2                |
      |                                                       |
      +-------------------------------------------------------+
                        3
    
## Gnutella 第二代P2P，洪泛请求模式
每个Peer的请求广播到连接的 Peers，各Peers又广播到各自连接的Peers，直到收到应答或达到5-9 hops

## BitTorrent
Tracker 
文件分块，chunk digest

## tapstry, pastry, chord, CAN 路由模式
目标是距离感知，减少hop
系统设计规模是10亿用户，10**14个文件


# skipnet

node name = com.wd.host1



    A                   N
    |                   |
    | join              |
    |------------------>|
    |                   |
    |             state |
    |<------------------|
    |                   |
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
