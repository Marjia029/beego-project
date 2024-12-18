<!DOCTYPE html>
<html>
<head>
    <title>Cat Breeds</title>
    <!-- Swiper CSS -->
    <link rel="stylesheet" href="https://unpkg.com/swiper/swiper-bundle.min.css" />
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
        }
        h1 {
            text-align: center;
            margin: 20px 0;
        }
        .swiper {
            width: 80%;
            margin: 20px auto;
            height: 500px; /* Set fixed height for the slider */
        }
        .swiper-slide {
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            text-align: center;
            padding: 20px;
            border: 1px solid #ccc;
            border-radius: 10px;
            background-color: #f9f9f9;
            box-sizing: border-box;
            height: 100%; /* Fill the container */
            overflow: hidden; /* Prevent overflow of content */
        }
        .swiper-slide img {
            max-width: 200px;
            max-height: 200px;
            border-radius: 10px;
            margin-bottom: 15px;
            object-fit: cover;
        }
        .swiper-slide h3 {
            font-size: 20px;
            margin-bottom: 10px;
            white-space: nowrap; /* Prevent long titles from wrapping */
            overflow: hidden;
            text-overflow: ellipsis;
        }
        .swiper-slide p {
            font-size: 14px;
            line-height: 1.5;
            margin: 5px 0;
            max-height: 80px; /* Limit description height */
            overflow: hidden;
            text-overflow: ellipsis;
        }
    </style>
</head>
<body>
    <h1>Cat Breeds Slider</h1>

    <!-- Swiper Slider -->
    <div class="swiper">
        <div class="swiper-wrapper">
            {{range .Breeds}}
                <div class="swiper-slide">
                    <img src="{{.ImageURL}}" alt="{{.Name}}" onerror="this.style.display='none';">
                    <h3>{{.Name}}</h3>
                    <p><strong>Origin:</strong> {{.Origin}}</p>
                    <p><strong>Temperament:</strong> {{.Temperament}}</p>
                    <p>{{.Description}}</p>
                </div>
            {{end}}
        </div>
        <!-- Add Pagination -->
        <div class="swiper-pagination"></div>
        <!-- Add Navigation -->
        <div class="swiper-button-prev"></div>
        <div class="swiper-button-next"></div>
    </div>

    <!-- Swiper JS -->
    <script src="https://unpkg.com/swiper/swiper-bundle.min.js"></script>
    <script>
        // Initialize Swiper
        var swiper = new Swiper('.swiper', {
            loop: true, // Enable infinite scrolling
            autoplay: {
                delay: 3000, // Auto slide every 3 seconds
                disableOnInteraction: false, // Continue auto-sliding after user interaction
            },
            pagination: {
                el: '.swiper-pagination',
                clickable: true, // Allow clicking on pagination dots
            },
            navigation: {
                nextEl: '.swiper-button-next',
                prevEl: '.swiper-button-prev',
            },
            slidesPerView: 1, // Show one slide at a time
            centeredSlides: true, // Center the active slide
        });
    </script>
</body>
</html>
