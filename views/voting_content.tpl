<h1>Vote for a Cat!</h1>

<!-- Cat image display -->
<div class="image-container">
    <img src="{{.ImageURL}}" alt="Cat Image">
</div>

<!-- Voting buttons -->
<div class="buttons-container">
    <!-- Add to favorites button -->
    <form method="POST" action="/voting">
        <input type="hidden" name="action" value="favorite">
        <input type="hidden" name="image_url" value="{{.ImageURL}}">
        <button type="submit" class="heart-icon" aria-label="Add to favorites">&#10084;</button>
    </form>

    <!-- Like and Dislike buttons -->
    <div class="like-dislike-container">
        <form method="POST" action="/voting">
            <input type="hidden" name="action" value="like">
            <input type="hidden" name="image_url" value="{{.ImageURL}}">
            <button type="submit" class="like-button" aria-label="Like">👍</button>
        </form>

        <form method="POST" action="/voting">
            <input type="hidden" name="action" value="dislike">
            <input type="hidden" name="image_url" value="{{.ImageURL}}">
            <button type="submit" class="dislike-button" aria-label="Dislike">👎</button>
        </form>
    </div>
</div>

<!-- Favorites section -->
<div class="favorites-container" id="favorites-section" style="display: none;">
    <h2>Your Favorites</h2>
    <ul>
        {{range .Favorites}}
            <li><img src="{{.}}" alt="Favorite Image" width="100"></li>
        {{end}}
    </ul>
</div>
