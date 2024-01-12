### shifu-demo

#### 使用步骤
1. 安装docker：https://docs.docker.com/desktop/install/linux-install/
2. 安装shifu
```bash
curl -sfL https://raw.githubusercontent.com/Edgenesis/shifu/main/test/scripts/shifu-demo-install.sh | sudo sh -
```
3. 下载程序代码
```bash
git clone https://github.com/p-mega/shifu-demo
```
4. 打包构建镜像
```bash
docker build -t 1454579997/shifu-demo:1.0 .
```
5. 将镜像上传到docker hub
```bash
# 使用docker login登录，输入docker hub用户名和密码
docker login
# 上传镜像
docker push 1454579997/shifu-demo:1.0
```
6. 构建Pod并运行
```bash
kubectl apply -f shifu-demo.yaml
```
7. 查看运行效果

![](./result.png)
