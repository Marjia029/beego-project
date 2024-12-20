// spa-handler.js

document.addEventListener('DOMContentLoaded', function() {
    // Store the current page state
    let currentState = {
        page: window.location.pathname,
        content: document.querySelector('.content-container').innerHTML
    };

    // Initial page load - store in history
    window.history.replaceState(currentState, '', currentState.page);

    // Handle navigation clicks
    document.querySelector('.nav-tabs').addEventListener('click', async function(e) {
        // Find closest anchor tag if clicking on child elements
        const link = e.target.closest('a');
        if (!link) return;

        e.preventDefault();
        const url = link.getAttribute('href');
        await loadPage(url);
    });

    // Handle browser back/forward
    window.addEventListener('popstate', function(e) {
        if (e.state) {
            document.querySelector('.content-container').innerHTML = e.state.content;
            updateActiveTab(e.state.page);
            initializePageScripts(e.state.page);
        }
    });

    // Load page content via AJAX
    async function loadPage(url) {
        try {
            const response = await fetch(url, {
                headers: {
                    'X-Requested-With': 'XMLHttpRequest'
                }
            });
            
            if (!response.ok) throw new Error('Network response was not ok');
            
            const html = await response.text();
            
            // Extract content from the response
            const tempDiv = document.createElement('div');
            tempDiv.innerHTML = html;
            const newContent = tempDiv.querySelector('.content-container').innerHTML;

            // Update the page
            document.querySelector('.content-container').innerHTML = newContent;
            
            // Update history and state
            const newState = {
                page: url,
                content: newContent
            };
            
            window.history.pushState(newState, '', url);
            currentState = newState;
            
            // Update active tab and reinitialize page-specific scripts
            updateActiveTab(url);
            initializePageScripts(url);

        } catch (error) {
            console.error('Error loading page:', error);
            // Fallback to traditional navigation on error
            window.location.href = url;
        }
    }

    // Update active tab in navigation
    function updateActiveTab(url) {
        const tabs = document.querySelectorAll('.nav-tabs .nav-item');
        tabs.forEach(tab => {
            tab.classList.remove('active');
            if (tab.getAttribute('href') === url) {
                tab.classList.add('active');
            }
        });
    }

    // Initialize page-specific scripts
    function initializePageScripts(url) {
        switch(url) {
            case '/breed-search':
                initializeBreedSearch();
                break;
            case '/voting':
                initializeVoting();
                break;
            case '/favorites':
                initializeFavorites();
                break;
        }
    }

    // Page-specific initializations
    function initializeBreedSearch() {
        if (typeof Swiper !== 'undefined') {
            const swiper = new Swiper('.swiper', {
                slidesPerView: 1,
                spaceBetween: 0,
                loop: true,
                autoplay: {
                    delay: 3000,
                    disableOnInteraction: false,
                    pauseOnMouseEnter: true,
                },
                pagination: {
                    el: '.swiper-pagination',
                    clickable: true,
                },
            });

            const swiperElement = document.querySelector('.swiper');
            if (swiperElement) {
                swiperElement.addEventListener('mouseenter', () => {
                    swiper.autoplay.stop();
                });
                swiperElement.addEventListener('mouseleave', () => {
                    swiper.autoplay.start();
                });
            }
        }

        // Handle breed selection form
        const breedForm = document.getElementById('breed-form');
        if (breedForm) {
            breedForm.addEventListener('submit', async function(e) {
                e.preventDefault();
                const formData = new FormData(breedForm);
                const response = await fetch('/breed-search', {
                    method: 'POST',
                    body: formData,
                    headers: {
                        'X-Requested-With': 'XMLHttpRequest'
                    }
                });
                const html = await response.text();
                const tempDiv = document.createElement('div');
                tempDiv.innerHTML = html;
                document.querySelector('.content-container').innerHTML = 
                    tempDiv.querySelector('.content-container').innerHTML;
                initializeBreedSearch(); // Reinitialize after content update
            });
        }
    }

    function initializeVoting() {
        const votingForms = document.querySelectorAll('form[action="/voting"]');
        votingForms.forEach(form => {
            form.addEventListener('submit', async function(e) {
                e.preventDefault();
                const formData = new FormData(form);
                await fetch('/voting', {
                    method: 'POST',
                    body: formData,
                    headers: {
                        'X-Requested-With': 'XMLHttpRequest'
                    }
                });
                // Reload voting page content
                await loadPage('/voting');
            });
        });
    }

    function initializeFavorites() {
        // Initialize any favorites-specific functionality
        const favsTab = document.querySelector('a[href="#favs-tab"]');
        const favoritesSection = document.getElementById('favorites-section');
        
        if (favsTab && favoritesSection) {
            favsTab.addEventListener('click', function(e) {
                e.preventDefault();
                if (favoritesSection.style.display === "none") {
                    favoritesSection.style.display = "block";
                } else {
                    favoritesSection.style.display = "none";
                }
            });
        }
    }
});