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
            <img id="cat-image" src="{{.ImageURL}}" alt="Cat Image">
        </div>

        <!-- Buttons container below the image -->
        <div class="buttons-container">
            <!-- Heart icon (add to favorites) -->
            <button id="favorite-button" class="heart-icon" data-action="favorite" data-image-url="{{.ImageURL}}">&#10084;</button>

            <div class="like-dislike-container">
                <!-- Like and Dislike buttons -->
                <button id="like-button" class="like-button" data-action="like" data-image-url="{{.ImageURL}}">👍</button>
                <button id="dislike-button" class="dislike-button" data-action="dislike" data-image-url="{{.ImageURL}}">👎</button>
            </div>
        </div>

       

    <!-- <script src="/static/js/voting.js"></script> -->

</body>
</html>
