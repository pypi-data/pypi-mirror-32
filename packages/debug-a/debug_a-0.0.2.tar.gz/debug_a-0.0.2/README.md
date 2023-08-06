# Debug_A -- A股盯盘机器人

>初入股市，对很多东西都理解不深，如果你有更好的盯盘策略需要实现，欢迎联系我，我的邮箱是 zeng_bin8888@163.com

## 1 - 缘由

A股变化莫测，光靠一双眼睛、
四台显示器盯盘难免会有疏漏。如果能有个程序，
模拟人的盯盘动作，实时监控盘面变化，
一旦发现“异常变化”还能发送预警消息，
那就再好不过了。本项目的目标就是实现这样一个程序，
即实现A股交易时间自动盯盘，
检测到“异常情况”，实时发送通知预警。

## 2 - Quick Start

### 2.1 - 安装

`pip install debug_a`

### 2.2 - 升级

`pip install debug_a --upgrade`

### 2.3 使用命令行工具

项目中实现了一个命令行工具，方便快速使用一些简单功能，如监控价格走势、查看市场行情等。
安装`debug_a`之后，打开cmd，执行 `da --help` 即可查看帮助文件。


## 感谢以下项目

- [tushare](https://github.com/waditu/tushare)
- [APScheduler](http://apscheduler.readthedocs.io/en/latest/)

