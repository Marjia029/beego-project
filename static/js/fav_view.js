document.addEventListener('DOMContentLoaded', function () {
    const favoritesList = document.getElementById('favorites-list');

    // Load favorites from localStorage
    const savedFavorites = localStorage.getItem('favorites');
    if (savedFavorites) {
        favorites = JSON.parse(savedFavorites);
    }

    // Set default view mode
    favoritesList.classList.add('grid-view');

    // Add event listeners to view buttons
    document.getElementById('grid-view-btn').addEventListener('click', () => {
        favoritesList.classList.remove('column-view');
        favoritesList.classList.add('grid-view');
    });

    document.getElementById('column-view-btn').addEventListener('click', () => {
        favoritesList.classList.remove('grid-view');
        favoritesList.classList.add('column-view');
    });

    // Function to display favorites
    function displayFavorites() {
        if (favorites.length > 0) {
            favoritesList.innerHTML = favorites.map(url =>
                `<li><img src="${url}" alt="Favorite Cat Image" width="200"></li>`
            ).join('');
        } else {
            favoritesList.innerHTML = '<p>You have no favorite cat images yet.</p>';
        }
    }

    // Display favorites initially
    displayFavorites();
});
