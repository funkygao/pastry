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
