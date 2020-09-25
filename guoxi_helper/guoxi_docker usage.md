# docker usage
当前问题：不知道 docker 环境怎么用
解决方案：从OIP协议入手，看看OIP协议的机理。

## flod
- FLO 的 Golang 语言全节点。
- RPC 端口并没有开放。
 
## oip
- oip 监控整个 FLO 区块链，将 oip 信息索引成一个可检索索引。
- oip 守护进程处理 FLO 区块链上所有的区块和交易,从交易中提取 OIP 记录，在1606 端口开放。
- 所有的 docker image
https://github.com/oipwg/oip

## ElasticSearch
- oip 守护进程的后端数据库。
- 为系统提供近乎立刻的复杂查询。
- 9200端口。

## Kibana
- 方便的 UI 界面来查看 oip 的数据库(ElasticSearch)。
- 5601端口。

## webwallet
- 轻量级的浏览器钱包。
