### `docker` 实验问题

## 1. 正常项目
- IPFS
- Web Wallet( warning 多，但可以正常运行)
- Go-Flo
- OIP ( warning )
- Elasticsearch
- Kibana

## 2. `caddy`问题

- 原因：`github`上包更新导致包不匹配，无法安装
```
# error message
go: creating new go.mod: module caddy
go: finding github.com/mholt/caddy v1.0.5
go: github.com/mholt/caddy@v1.0.5: parsing go.mod: unexpected module path "github.com/caddyserver/caddy"
go: error loading module requirements
```

- 解决方案
1. (最终方案)搜集 `caddy` 资料，改写 `github` 代码。

```
package main

import (
//	"github.com/mholt/caddy/caddy/caddymain"
	“github.com/caddyserver/caddy/tree/master/cmd/caddy”

	// plug in plugins here
//	_ "github.com/captncraig/cors/caddy"
	_ "github.com/captncraig/cors/tree/master/caddy"
)

func main() {
	// optional: disable telemetry
	caddymain.EnableTelemetry = false
	caddymain.Run()
}
```

  1. 下载 2019 年 5 月 1 日 `docker` 配置提交前的两个代码，定位到原始版本。

  - github.com/mholt/caddy/caddy/caddymain
    
    1. 在 `github` 上 `fork` 原始版本 https://github.com/caddyserver/caddy `15fecbc16151308959674a5ce3843efb9e66bd5f`

    2. `github` 上删除多余的 `branch`

    3. 切换到 `tag v1.0.0`。 

    ```
    git clone https://github.com/GuoxiW/caddy
    cd caddy
    git reset --hard 15fecbc16151308959674a5ce3843efb9e66bd5f
    git push -u origin master -f (不建议，可能出问题)
    ```

    4. `VS Code` 改动失效地址并提交更改

       1. `git clone` 下来用 `vscode` 
       2. 把 `mholt/caddy` 替换为 `GuoxiW/caddy`
       3. 提交更改

    5. 生成新的 `tag v1.0.6`
    ```
    git tag -a v1.0.6 -m "1.0.6"
    git push origin v1.0.6
    ```
  - github.com/captncraig/cors/caddy
    1. 在 `github` 上 `fork` 原始版本 https://github.com/captncraig/cors `c5b990a6b7dedbe4040a01d1eed51b1fa58b28ec`

    2. `github` 上删除多余的 `branch`。

    3. 切换到指定 `tag` ( 3 月 19 日版本)。
    ```
    git clone https://github.com/GuoxiW/cors/
    cd cors
    git reset --hard c5b990a6b7dedbe4040a01d1eed51b1fa58b28ec
    git push -u origin master -f (不建议，可能出问题)
    ```

    4. `VS Code` 改动失效地址并提交更改

       1. `git clone` 下来用 `vscode` 
       2. 把 `mholt/caddy` 替换为 `GuoxiW/caddy`
       3. 把 `captncraig/cors` 替换为 `GuoxiW/cors`
       4. 提交更改
    ```
    # 替换记录
    # cors/caddy/parse_test.go
    github.com/mholt/caddy   ->   github.com/GuoxiW/caddy

    # cors/caddy/corsPlugin.go 
    github.com/captncraig/cors   ->   github.com/GuoxiW/cors
    github.com/mholt/caddy  ->   github.com/GuoxiW/caddy
    github.com/mholt/caddy/caddyhttp/httpserver   ->   github.com/GuoxiW/caddy/caddyhttp/httpserver
    ```

2. 修改官方 `docker` 中的 `caddy` 配置。
  
    1. 在 `github` 上 `fork` 原始版本 https://github.com/oipwg/docker

    2. `VS Code` 改动失效地址并提交更改

    ```
    git clone https://github.com/GuoxiW/docker/
    cd docker
    ```

    ```
    # caddy/Dockerfile
    go get github.com/mholt/caddy && \   ->   go get github.com/GuoxiW/caddy && \
    ```

    ```
    # caddy/with_plugins.go
    import (
	    //"github.com/mholt/caddy/caddy/caddymain"
	    "github.com/GuoxiW/caddy/caddy/caddymain"

	    // plug in plugins here
	    //_ "github.com/captncraig/cors/caddy"
	    _ "github.com/GuoxiW/cors/caddy"
    )
    ```


2. (未使用)直接用官方 `caddy` ，在其上修改。
```
docker pull caddy/caddy
```


