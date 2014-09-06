<html>
<head>
<title></title>
</head>
<body>

<form action="http://127.0.0.1:9090/resume" method="post">
简历:<input name="cvfile" style="width:200px" value=""/>

    <input type="submit" id="btnContent" value="获取内容"/>
    </form>
    <form action="http://127.0.0.1:9090/analysis" method="post">

    <p>
    <textarea  name="content" cols="100" rows="40">{{.}}</textarea>
    </p>
    <p>
    <input type="submit" value="提交">
    </p>
</form>
</body>
</html>