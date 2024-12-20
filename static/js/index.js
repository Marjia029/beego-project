function handleButtonClick(event) {
    const button = event.target;
    const action = button.dataset.action;
    const imageURL = button.dataset.imageUrl;

    // Create a FormData object to send to the server
    const formData = new FormData();
    formData.append("action", action);
    formData.append("image_url", imageURL);

    // Send the data to the server using Fetch API (AJAX)
    fetch('/voting', {
        method: 'POST',
        body: formData,
    })
    .then(response => response.json())
    .then(data => {
        // Update the image URL
        document.getElementById('cat-image').src = data.image_url;

        // Update favorites section
        const favoritesList = document.getElementById('favorites-list');
        favoritesList.innerHTML = '';
        data.favorites.forEach(url => {
            const listItem = document.createElement('li');
            const img = document.createElement('img');
            img.src = url;
            img.width = 100;
            listItem.appendChild(img);
            favoritesList.appendChild(listItem);
        });

        // Show the favorites section if there are any favorites
        if (data.favorites.length > 0) {
            document.getElementById('favorites-section').style.display = 'block';
        }
    })
    .catch(error => console.error('Error:', error));
}
 
 // This function loads the voting content by default when the page loads
 document.addEventListener('DOMContentLoaded', function() {
    // Send an AJAX request to the /voting route
    fetch('/voting')
        .then(response => response.text()) // Get the response as text
        .then(data => {
            // Inject the entire voting template into the #voting-container
            const votingContainer = document.getElementById('voting-container');
            votingContainer.innerHTML = data;

            // Now attach event listeners to buttons after voting template is loaded
            document.getElementById('favorite-button').addEventListener('click', handleButtonClick);
            document.getElementById('like-button').addEventListener('click', handleButtonClick);
            document.getElementById('dislike-button').addEventListener('click', handleButtonClick);
        })
        .catch(error => {
            console.error('Error loading voting content:', error);
            alert('Failed to load voting content.');
        });
});


document.getElementById('breed-link').addEventListener('click', function(e) {
    e.preventDefault(); // Prevent the default link behavior

    // Send an AJAX request to the /voting route
    fetch('/breed-search')
        .then(response => response.text()) // Get the response as text
        .then(data => {
            // Inject the returned voting template into the #voting-container
            document.getElementById('voting-container').innerHTML = data;
        })
        .catch(error => {
            console.error('Error loading voting content:', error);
            alert('Failed to load voting content.');
        });
});