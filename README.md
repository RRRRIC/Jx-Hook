# Jx-hook

A webhook that transfer webhook msg to wechat-work (will support other platform in the future).

It's light weight, configurable tool to help resolve different webhook to you own platform.

## Usage:

create a wechat sender

````
PUT "/sender/save"
{
  "id": "sender001", // Empty id means create &  id means modify or specifced id sender
  "name": "jinxin",
  "wechat_robot_key": "123451212312123",
  "template_msg": "Alert from ${alert.host} is ${current_value}"
  "alert_type": "text",
  "enable": true
}
````



create a custom sender

```
PUT "/sender/save"
{
  "id": "sender001", // Empty id means create &  id means modify or specifced id sender
  "name": "jinxin",
  "wechat_robot_key": "123451212312123",
  "template_msg": "{\"host\" : \"${alert.host}\", \"current_value\": \"${current_value}\"}"
  "custom_url" : "https://my-own-platform/alert"
  "alert_type": "custom",
  "enable": true
}
```





And create a alert to send this msg 

```
PUT "/alert/save"
{
  "id": "alert001", // Empty id means create & No-empty id means modify or specifced id alert
  "id" : "xxxx"
  "sender_ids": [ "sender001", "sender002"]
  "enable": true
}
```



You can configure the web hook url like `http://xxxx/alert/do/${alert_id}`

So when you received an alert from platform like `grafna` and its data like blow

```json
{
  "alert": {
    "host": "host001"
  },
  "current_value": 1000
  "xxx": "xxxx"
}
```

It will automaticlly send a msg to wechat and fill the value of template to "Alert from host001 is 1000"



>   Here's the grafana's alert template
>
>   https://grafana.com/docs/grafana-cloud/alerting/manage-notifications/webhook-notifier/

## Install

### Requirement:

go >= 1.19.1

redis



### Run:

```bash
git clone git@github.com:RRRRIC/Jx-Hook.git
cd ./jx-hook
go build

# service default port 8080
# redis default 127.0.0.1:6379 db 1
./jx-hook -c ${YOU_OWN_CONFIG_FILE}
```

