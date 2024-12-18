<!DOCTYPE html>
<html>
<head>
    <title>API Data - Beego</title>
</head>
<body>
    <h1>Posts from the API</h1>
    <ul>
        {{range .Posts}}
            <li>
                <strong>{{.Title}}</strong><br>
                {{.Body}}
            </li>
        {{end}}
    </ul>
</body>
</html>
