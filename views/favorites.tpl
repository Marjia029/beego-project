<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Your Favorites</title>
    <link rel="stylesheet" href="/static/css/voting.css">
    <link rel="stylesheet" href="/static/css/main.css">
</head>
<body>
    <div class="content-container">
        {{template "navbar" .}}

        <h1>Your Favorite Cat Images</h1>

        <div class="favorites-container">
            {{if .Favorites}}
                <ul>
                    {{range .Favorites}}
                        <li><img src="{{.}}" alt="Favorite Cat Image" width="200"></li>
                    {{end}}
                </ul>
            {{else}}
                <p>You have no favorite cat images yet.</p>
            {{end}}
        </div>
    </div>
</body>
</html>
