<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat Lovers SPA</title>
    <link rel="stylesheet" href="/static/css/main.css">
    <script src="/static/js/spa-handler.js" defer></script>
</head>
<body>
    <header>
        <nav class="navbar">
            <ul class="nav-tabs">
                <li class="nav-item"><a href="/" class="nav-link">Home</a></li>
                <li class="nav-item"><a href="/breed-search" class="nav-link">Breed Search</a></li>
                <li class="nav-item"><a href="/voting" class="nav-link">Voting</a></li>
                <li class="nav-item"><a href="/favorites" class="nav-link">Favorites</a></li>
            </ul>
        </nav>
    </header>
    <main>
        <div class="content-container">
            <h1>Welcome to Cat Lovers SPA</h1>
            <p>Explore cat breeds, vote for your favorite cats, and save them to your favorites!</p>
        </div>
    </main>
    <footer>
        <p>&copy; 2024 Cat Lovers. Contact us at {{.Email}}</p>
    </footer>
</body>
</html>
