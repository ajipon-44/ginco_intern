# ginco_intern

## 内容
ソーシャルゲームによくあるガチャを引くことができるAPI

## 機能
```
# ユーザー情報の追加
user/create
input body...{name: string}
output {token: string}

# 閲覧
/user/get
input header...token: string
output {name: string}

# 更新
user/update
input {token: string, name: string}

# キャラクターガチャを引く
/gacha/draw
input header...token, body...{times: int, gacha_id: int}
output
{
    "results": [
        {
            "characterID": string,
            "name": string
        },
    ]
}
      

# 自分が持っているキャラクターの閲覧
character/list
input header...token
output
{
    "characters": [
        {
            "userCharacterID": string,
            "characterID": string,
            "name": string
        },
```

参考：https://github.com/CyberAgentHack/techtrain-mission
