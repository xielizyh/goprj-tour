# 总结
1. init函数

   有关golang package中init[方法的多处定义及运行顺序问题](https://blog.csdn.net/zhuxinquan61/article/details/73712251)。

   - 一个包中可以包含多个init函数，执行顺序按照文件名执行
   - 一个源文件中也可以包含多个init函数，执行顺序按照代码编写顺序执行

2. vscode启动带参数调试go

   VsCode Go[插件配置最佳实践指南](https://zhuanlan.zhihu.com/p/320343679)，参考的launch.json如下：

   ```json
   {
       // Use IntelliSense to learn about possible attributes.
       // Hover to view descriptions of existing attributes.
       // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
       "version": "0.2.0",
       "configurations": [
           {
               "name": "Launch",
               "type": "go",
               "request": "launch",
               "mode": "auto",
               "program": "${workspaceFolder}/main.go",
               "cwd": "${workspaceFolder}",
               "env": {},
               "args": ["time", "calc", "-c", "2022-10-09 11:14:22", "-d", "5m"],
           }
       ]
   }
   ```

   

