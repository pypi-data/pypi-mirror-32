# 情感分析结果标注格式转换（CSV to JSON）

这个工具将情感分析结果标注的格式从CSV转换为JSON。输出格式依照[标注规范](http://aiwiki.yimian.com.cn/labelling/standard.html)

## 安装：

	#把bitbucket_user_id替换为你的bitbucket账号
	pip install git+https://bitbucket_user_id@bitbucket.org/yimian/fmtlabels
	
## 使用格式：

	lblcsv2json <category> [input.csv [output.json]]

如果省略input和output，分别使用stdin和stdout。

例：

	lblcsv2json skincare xxx.csv yyy.json
	lblcsv2json skincare xxx.csv > yyy.json
	cat xxx.csv | lblcsv2json skincare > yyy.json
