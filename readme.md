https://github.com/speps/go-hashids 的方法封装，数字类型的 ID 转换为随机字符串 ID

不再维护该仓库，代码已迁移到 [goutils](http://github.com/axiaoxin-com/goutils/hashids.go)

安装：

    go get -u github.com/axiaoxin/hashids

用法示例[example.go](https://github.com/axiaoxin/hashids/blob/master/example/example.go)：

    package main

    import (
        "log"

        "github.com/axiaoxin/hashids"
    )

    func main() {
        salt := "my-salt:appid:region:uin"
        minLen := 8
        prefix := ""
        h, err := hashids.New(salt, minLen, prefix)
        if err != nil {
            log.Fatal(err)
        }
        var id int64 = 1
        strID, err := h.Encode(id)
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("int64 id %d encode to %s", id, strID)

        int64ID, err := h.Decode(strID)
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("string id %s decode to %d", strID, int64ID)
    }

运行结果：

    go run example.go
    2020/02/26 13:28:11 int64 id 1 encode to 8Gnejq6A
    2020/02/26 13:28:11 string id 8Gnejq6A decode to 1
