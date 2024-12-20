<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="https://unpkg.com/swiper/swiper-bundle.min.css" />
    <link rel="stylesheet" href="/static/css/voting.css">
    <link rel="stylesheet" href="/static/css/main.css">
</head>
<body>
    <div class="content-container">
        {{template "navbar" .}}
        {{template "content" .}}
    </div>
    <script src="/static/js/favs.js"></script>
    <script src="/static/js/spa-handler.js"></script>
</body>
</html>
