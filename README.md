## Jenkins 插件下载换国内源

- 清华源仅做镜像，而Jenkins的插件下载地址在Json内配置。
- 因此使用清华源镜像的Json，插件仍在jenkins.io下载。
- Json内配置的检测地址使用http://www.google.com/,国内无法连接，Jenkins以为自己断线，改为http://www.baidu.com/

#### 配置项

| 名称 | 意义 | 默认值 |
| :---: | :---: | :---: |
| p | 端口 | 18081 |
| s | 源Json地址 | https://mirrors.tuna.tsinghua.edu.cn/jenkins/updates/current/update-center.json |
| h | 显示帮助 | false |

#### 使用方法
- 下载二进制文件
- 启动（最简方式：./redirect-jenkins-to-tsinghua-{平台}-amd64）
- 配置Jenkins的pluginManager/advanced，至http://{本服务IP}:{本服务端口}/json
- 完成