package leranGit

import "fmt"

//本地代码提交到仓库
/*
	1、怎么提交呢，首先要让本地的（当前工程目前）变成一个git仓库，执行 git init即可
	2、git status 查看当前仓库状态，如果之前有文件修改，会有modified标红。
	3、提交，git add . 	点表示全部提交,这个只是放在了缓存区中，没有真正的提交
	4、真正提交(其实是提交到本地)，git commit -m "m的备注：如修改了XX.go" {如果提交失败，注意当前目录的.git文件夹里的配置文件}

	添加到自己的服务器仓库
	1、找到远程仓库的地址如：https://github.com/pengzihuan/test1.git
	2、git remote add origin https://github.com/pengzihuan/test1.git	origin是自己取的别名
		remote是远端的意思，执行远端的操作，方式是：add，通过给的地址将远端与本地绑定
	3、git push -u origin master	将本地的代码提交到远端，远端叫：origin， 本地叫：master(可通过 git branch 查看)
*/

func main()  {
	fmt.Println("learngit")

	fmt.Println("1")
}