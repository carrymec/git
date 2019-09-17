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
17. 远程分支 查看远程分支 git branch -r
18. 将本地分支提交到远程仓库 git push origin v3(分支名)
19. 切换到分支 git checkout v3
20. 删除远程分支 git branch -r -d origin/v3(分支名)
21. 将删除的分支推送到远程，并在远程删除这个分支 git push origin :v3
22. ps 在分支上提交代码不能直接push
23. 分支上提交代码 git push --set-upstream origin v3(先git add . git commit -m "msg log")
24. 合并同上13