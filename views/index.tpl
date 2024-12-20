<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat Lovers SPA</title>
    <link rel="stylesheet" href="/static/css/main.css">
</head>
<body>
    
    <div class="content-container">
        {{template "nav" .}}

        <!-- Container where the voting template will be loaded -->
        <div id="voting-container"></div>
        <div id="breed-container"></div>
        <div id="fav-container"></div>
    </div>
    
    <footer class="footer">
        <p>&copy; 2024 Cat Lovers. Contact us at {{.Email}}</p>
    </footer>

    <script src="/static/js/index.js"></script> 

    <!-- <script src = "/static/js/voting.js"></script> -->

</body>
</html>
