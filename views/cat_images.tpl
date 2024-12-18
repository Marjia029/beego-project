<!DOCTYPE html>
<html>
<head>
    <title>Cat Images</title>
</head>
<body>
    <h1>Random Cat Images</h1>
    <div>
        {{range .Images}}
            <img src="{{.url}}" alt="Cat Image" style="width:200px; height:auto; margin:10px;">
        {{end}}
    </div>
</body>
</html>
