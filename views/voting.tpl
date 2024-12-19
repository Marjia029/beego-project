<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Voting</title>
    <link rel="stylesheet" href="https://unpkg.com/swiper/swiper-bundle.min.css" />
    <link rel="stylesheet" href="/static/css/voting.css">
    <link rel="stylesheet" href="/static/css/main.css">
</head>
<body>
    <div class="content-container">

        {{template "navbar" .}}

        <h1>Vote for a Cat!</h1>

        <div class="image-container">
            <img src="{{.ImageURL}}" alt="Cat Image">
        </div>

        <!-- Buttons container below the image -->
        <div class="buttons-container">
            <!-- Heart icon (add to favorites) -->
            <form method="POST" action="/voting">
                <input type="hidden" name="action" value="favorite">
                <input type="hidden" name="image_url" value="{{.ImageURL}}">
                <button type="submit" class="heart-icon">&#10084;</button>
            </form>

            <div class = "like-dislike-container">
                    <!-- Like and Dislike buttons -->
                <form method="POST" action="/voting">
                    <input type="hidden" name="action" value="like">
                    <input type="hidden" name="image_url" value="{{.ImageURL}}">
                    <button type="submit" class="like-button">👍</button>
                </form>

                <form method="POST" action="/voting">
                    <input type="hidden" name="action" value="dislike">
                    <input type="hidden" name="image_url" value="{{.ImageURL}}">
                    <button type="submit" class="dislike-button">👎</button>
                </form>

            </div>
        </div>

        <div class="favorites-container">
            <h2>Your Favorites</h2>
            <ul>
                {{range .Favorites}}
                    <li><img src="{{.}}" alt="Favorite Image" width="100"></li>
                {{end}}
            </ul>
        </div>
    </div>
</body>
</html>
