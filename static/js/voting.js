// Function to handle button clicks
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

// Add event listeners to buttons
document.getElementById('favorite-button').addEventListener('click', handleButtonClick);
document.getElementById('like-button').addEventListener('click', handleButtonClick);
document.getElementById('dislike-button').addEventListener('click', handleButtonClick);