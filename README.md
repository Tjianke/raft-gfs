# raft-gfs

## 规划

1. 学习raft协议
2. 学习etcd-raft接口的设计
3. 学习GFS
4. 将GFS的master改造成基于raft的内存kv集群

## 使用
```bash
go build
./raft-gfs --id=1
./raft-gfs --id=2
./raft-gfs --id=3
```