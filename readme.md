# OIPWG/docker

Docker Compose configurations for basic FLO/OIP related daemons.


## Prerequisites
- Docker Compose - https://docs.docker.com/compose/install/


## Currently supported daemons
- Go-Flo Mainnet
- Go-Flo Testnet
- OIP Web Wallet
- OIP Mainnet
- OIP Testnet
- Elasticsearch
- Kibana
- Caddy
- IPFS

# Common


### Caddy
Caddy provides a reverse proxy to the potential running services herein  
Configuration/Customization may take place by modifying `caddy/Caddyfile` and adding plugins to `caddy/with_plugins.go`

By default Caddy will provide http:// on port 80 listening on all interfaces

```
http://<ip>/                    Web wallet
http://<ip>/mainnet/oip         OIP daemon apis running on mainnet
http://<ip>/mainnet/kibana      Kibana instance running on mainnet
                                user:pass = kibana:mainnet
http://<ip>/testnet/oip         OIP daemon apis running on testnet
http://<ip>/testnet/kibana      Kibana instance running on testnet
                                user:pass = kibana:testnet
http://<ip>/ipfs/               IPFS Gateway
```

- Build image - necessary if Caddyfile or with_plugins.go are modified `sudo docker-compose build caddy`
- Run daemon in background `sudo docker-compose up -d caddy`
- Tail daemon logs `sudo docker-compose logs -f caddy`



### IPFS
A local directory at `./ipfsStaging` is available and mounted as `/export` within the IPFS container

- Build image - necessary if config file is modified `sudo docker-compose build ipfs`
- Run daemon in background `sudo docker-compose up -d ipfs` **-d 后台运行**
- Tail daemon logs `sudo docker-compose logs -f ipfs`


### Web Wallet
Web wallet provides a multi currency web accessible wallet

- Run daemon in background `sudo docker-compose up -d webwallet`
- Tail daemon logs `sudo docker-compose logs -f webwallet`



# Mainnet


### Go-Flo
Modify `go-flo/mainnet.conf`

- Build image - necessary if config file is modified `sudo docker-compose build goflomainnet`
- Run daemon in background `sudo docker-compose up -d goflomainnet`
- Tail daemon logs `sudo docker-compose logs -f goflomainnet`


### OIP
Modify `oip/config.mainnet.yml`

- Build image - necessary if config file is modified `sudo docker-compose build oipmainnet`
- Run daemon in background `sudo docker-compose up -d oipmainnet`
- Tail daemon logs sudo `sudo docker-compose up -d webwalletdocker-compose logs -f oipmainnet`


### Elasticsearch
Elasticsearch mainnet defaults to 4GB Heap size, adjust `ES_JAVA_OPTS` as appropriate

- Run daemon in background `sudo docker-compose up -d esmainnet`
- Tail daemon logs `sudo docker-compose logs -f esmainnet`


### Kibana

- Run daemon in background `sudo docker-compose up -d kibanamainnet`
- Tail daemon logs `sudo docker-compose logs -f kibanamainnet`



# Testnet


### Go-Flo
Modify `go-flo/testnet.conf`

- Build image - necessary if config file is modified `sudo docker-compose build goflotestnet`
- Run daemon in background `sudo docker-compose up -d goflotestnet`
- Tail daemon logs `sudo docker-compose logs -f goflotestnet`


### Elasticsearch
Elasticsearch testnet defaults to 2GB Heap size, adjust `ES_JAVA_OPTS` as appropriate
```
# to solve 
ERROR: [1] bootstrap checks failed
[1]: max virtual memory areas vm.max_map_count [65530] is too low, increase to at least [262144]
```

```
sudo vim /etc/sysctl.conf　　
```
```
vm.max_map_count=655360
```
```
sudo sysctl -p
```

- Run daemon in background `sudo docker-compose up -d estestnet`
- Tail daemon logs `sudo docker-compose logs -f estestnet`


### OIP
Modify `oip/config.testnet.yml`

- Build image - necessary if config file is modified `sudo docker-compose build oiptestnet`
- Run daemon in background `sudo docker-compose up -d oiptestnet`
- Tail daemon logs `sudo docker-compose logs -f oiptestnet`


### Kibana

- Run daemon in background `sudo docker-compose up -d kibanatestnet`
- Tail daemon logs `sudo docker-compose logs -f kibanatestnet`




|        Name          |        Command       |   State   |       Ports|
|:----:|:----:|:----:|:----:|
|docker_caddy_1       |  /sbin/tini -- caddy -agree ... |  Up   |  0.0.0.0:443->443/tcp, 0.0.0.0:80->80/tcp |
|docker_estestnet_1   |  /usr/local/bin/docker-entr ... | Up    | 9200/tcp, 9300/tcp                       |
|docker_goflotestnet_1|  /flod/bin/flod --configfil ... |    Up |    0.0.0.0:17316->17316/tcp, 127.0.0.1:17317->17317/tcp |
|docker_ipfs_1        |  /sbin/tini --/usr/local/b ...  |    Up |    0.0.0.0:4001->4001/tcp, 4001/udp, 127.0.0.1:5001->5001/tcp, 127.00.|1:8080->8080/tcp, 8081/tcp |
|docker_kibanatestnet_1 |  /usr/local/bin/kibana-docker | Up    | 127.0.0.1:15601->5601/tcp                |                             
|docker_oiptestnet_1    | /oip/bin/oipd --appdir=/oip   |   Up  |   127.0.0.1:11606->11606/tcp            |                         
|docker_webwallet_1     | docker-entrypoint.sh npm start|   Up  |   127.0.0.1:7000->7000/tcp|
