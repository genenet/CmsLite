<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    {{range .result}}
    <div><li><a href="http://www.51shjzh.com/info/{{.Id}}" target="blank">{{ .Title }}</a></li></div>
    {{else}}
    <div><strong>no rows</strong></div>
    {{end}}

</body>
</html>