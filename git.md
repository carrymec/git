# git使用教程

1. 初始化 git init
2. 将项目添加到缓存中 git add .
3. 将缓存文件commit到git库 git commit -m "commit msg"
4. 在远程仓库创建库 https://git.oschina.net/liuqiqiang/gitTest.git  
5. 将本地库连接到远程库 git remote add origin https://git.oschina.net/liuqiqiang/gitTest.git  
6. 先更新本地仓库 git pull origin master
7. 提交 git push origin master
8. 分支管理 新建分支 git branch v1
9. 查看分支 git branch
10. 切换分支 git checkout v1
11. 提交分支修改 git add .(或者文件名)，git commit -m “v1 commit msg”
12. 切换到主分支 git checkout master
13. 将新分支提交的改动提交到主分支 git merge v1
14. 如果有冲突 查看冲突 git diff
15. 提交同上7
16. 删除分支 git branch -d v1