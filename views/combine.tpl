<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat Browser</title>
    <link rel="stylesheet" href="/static/css/main.css">
</head>
<body>
    <div class="content-container">
        {{template "nav" .}}

        <!-- Content Section -->
        <div id="content">
            <!-- Initially load the "Breeds" template -->
            {{template "breed_search.tpl" .}}
        </div>
    </div>

    <!-- JavaScript -->
    <script src="/static/js/navigation.js"></script>
</body>
</html>
