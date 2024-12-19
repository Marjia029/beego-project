document.addEventListener('DOMContentLoaded', function() {
    const favsTab = document.querySelector('a[href="#favs-tab"]'); // Your favs tab
    const favoritesSection = document.getElementById('favorites-section');

    // Add event listener to the "Favs" tab
    favsTab.addEventListener('click', function(event) {
        event.preventDefault(); // Prevent the default action (link following)
        
        // Toggle visibility of the favorites section
        if (favoritesSection.style.display === "none") {
            favoritesSection.style.display = "block"; // Show favorites
        } else {
            favoritesSection.style.display = "none"; // Hide favorites
        }
    });
});
