# aliyun-iot-device-python

阿里云 IOT 套件设备端 Python 开发SDK


## 支持的协议

- [x] MQTT
- [x] HTTP
- [ ] CoAP


## 环境

- Python3


## 安装

`pip3 install yansondga-aliyun-iot-device`


## 使用

### MQTT

```python
from aliyun_iot_device.mqtt import Client as IOT
import time


PRODUCE_KEY = "b1VzFx30hEm"
DEVICE_NAME = "iot_device_01"
DEVICE_SECRET = "3TSqd6sfzjSkSwEmLmcAdZnI0oGlmRZ8"
CLIENT_ID = ""

iot = IOT((PRODUCE_KEY, DEVICE_NAME, DEVICE_SECRET))

iot.connect()

iot.loop_start()
while True:
    iot.publish('success', 1)
    time.sleep(5)

# iot.mqtt 为原始 mqtt 客户端
```

#### 说明

-  MQTT - 域名直连与 HTTPS 认证

    SDK 默认使用域名直连同时启用 TLS 加密。

    如果您不想使用 TLS 加密，可在初始化时传入 `tls=False` 参数；

    如果您想使用 HTTPS 认证，可在初始化时传入 `domain_direct=False` 参数，HTTPS 认证将强制使用 TLS 认证加密

- MQTT - TLS 认证 CA 证书

    SDK 默认使用了阿里云 IOT 根证书，一般情况无需修改。

    如一定要修改，请传入 `ca_certs="/path/to/cert/root.cer"` 

- MQTT - websocket 通道

    SDK 默认使用 TCP 通道。

    如果需要使用 websocket，请传入 `transport="websockets"`。

    当使用 websocket 时，默认启用 TLS，即使用的是 wss 协议，如果不想使用 wss，请同时传入 `tls=False`

### HTTP

```python
from aliyun_iot_device.http import Client as IOT
import time

PRODUCE_KEY = "b1VzFx30hEm"
DEVICE_NAME = "iot_device_01"
DEVICE_SECRET = "3TSqd6sfzjSkSwEmLmcAdZnI0oGlmRZ8"

iot = IOT((PRODUCE_KEY, DEVICE_NAME, DEVICE_SECRET))

while True:
    iot.publish('success')
    time.sleep(5)
```

#### 注意

- 使用 http 协议进行通讯时，需要 token 进行认证，SDK 默认使用内存型缓存（cachetools 方案）进行 token 的保存。

- 如果您需要自行进行其他方案进行保存（file/memcached/redis），可以 `iot.get_token(cache=False)` 获取 token，再 publish 消息时，请 `iot.publish(payload=payload, token=token)`