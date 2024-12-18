<!DOCTYPE html>
<html>
<head>
    <title>Cat Breeds</title>
</head>
<body>
    <h1>List of Cat Breeds</h1>
    <ul>
        {{range .Breeds}}
            <li>
                <strong>Name:</strong> {{.Name}}<br>
                <strong>Origin:</strong> {{.Origin}}<br>
                <strong>Temperament:</strong> {{.Temperament}}<br>
                <strong>Description:</strong> {{.Description}}<br>
            </li>
        {{end}}
    </ul>
</body>
</html>
