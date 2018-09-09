# MTProxy

`Golang`版本的`MTProxy`，仅本人为了能在`Docker`的`Alpine`中运行而修改，源代码为[https://github.com/vkopitsa/MTProxy](https://github.com/vkopitsa/MTProxy)。

此版本自动生成`Secret`，并保存在`/data/secret`文件中，下次再次运行时将首先检查`/data/secret`，若没有则生成，有则直接使用。

此版本适合在`Docker`中使用，可直接使用仓库中的`Dockerfile`一键构建，也可自行编写`Docker-Compose`。

程序固定运行于`8822`端口，可在`Docker`外部进行绑定。