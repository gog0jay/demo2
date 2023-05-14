# 说明

## Step01: 输入输出
本节内容：
    - 读取文件
    - 打印到标准输出
    - 处理多返回值
    - 处理errors
    - 创建一个slice并向里面添加元素
    - slice的range循环
    - Defer
    - Log errors

### 首先，我们有一个maze01.txt, 这是一个迷宫的ASCII表示
```
    - # 表示一堵墙
    - . 表示一个点
    - P 表示玩家
    - G 表示鬼怪（敌人）
    - X 表示强化药丸
```

