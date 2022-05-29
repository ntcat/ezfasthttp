# ezfasthttp
easy fasthttp wrapper for golang

解析 web json 可以如此简单快速。 
简化fasthttp的使用。
大部分情况下，二行就得到想到的数据：

```GO
	cli := client.NewFastClient(url)
	data, err := cli.GetJsonDo(nil, nil, dataMapPrase)
```

订制化
一级订制化：如果需要订制化，一般通过重写GetJsonDo的三个回调函数即可完成，
二级订制化：如果还不够，你可以看GetJsonDo内容，分步写就行了。
PostJsom、PostForm类似。