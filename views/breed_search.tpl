<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Search Cat Breeds</title>
    <link rel="stylesheet" href="https://unpkg.com/swiper/swiper-bundle.min.css" />
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            text-align: center;
        }
        h1 {
            margin: 20px 0;
        }
        .search-container {
            margin: 20px 0;
        }
        .breed-info {
            margin: 10px 20px;
            max-width: 800px;
            margin-left: auto;
            margin-right: auto;
        }
        .breed-info p {
            line-height: 1.6;
        }
        .swiper {
            width: 100%;
            max-width: 500px;
            margin: 20px auto;
            height: 300px;
        }
        .swiper-slide {
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            text-align: center;
            padding: 10px;
            background-color: #f9f9f9;
            border-radius: 10px;
            height: 100%;
            background-size: contain;
            background-repeat: no-repeat;
            background-position: center;
            opacity: 0;
            transition: opacity 0.3s ease;
            position: absolute;
            width: 100%;
            top: 0;
            left: 0;
            z-index: 1;
        }
        .swiper-slide-active {
            opacity: 1;
            z-index: 2;
        }
        .swiper-slide img {
            display: none;
        }
        @media (max-width: 600px) {
            .swiper {
                height: 250px;
                max-width: 350px;
            }
            .breed-info {
                padding: 0 10px;
            }
        }
    </style>
</head>
<body>
    <h1>Search Cat Breeds</h1>
    <div class="search-container">
        <form method="POST" action="/breed-search">
            <select name="breed_id" required>
                <option value="">Select a breed</option>
                {{range .Breeds}}
                    <option value="{{.ID}}">{{.Name}}</option>
                {{end}}
            </select>
            <button type="submit">Search</button>
        </form>
    </div>

    {{if .SelectedBreed}}
        <div class="breed-info">
            <h2>{{.SelectedBreed.Name}}</h2>
            <p>{{.SelectedBreed.Description}}</p>
        </div>
        <div class="swiper">
            <div class="swiper-wrapper">
                {{range .BreedImages}}
                    <div class="swiper-slide" style="background-image: url('{{.URL}}');">
                        <img src="{{.URL}}" alt="Breed Image" loading="lazy" />
                    </div>
                {{end}}
            </div>
            <div class="swiper-pagination"></div>
            <div class="swiper-button-prev"></div>
            <div class="swiper-button-next"></div>
        </div>
    {{else if .Error}}
        <p style="color: red;">{{.Error}}</p>
    {{end}}

    <script src="https://unpkg.com/swiper/swiper-bundle.min.js"></script>
    <script>
        var swiper = new Swiper('.swiper', {
            slidesPerView: 1,
            spaceBetween: 0,
            centeredSlides: true,
            loop: true,
            speed: 500,
            autoplay: {
                delay: 3000,
                disableOnInteraction: false,
                pauseOnMouseEnter: true,
            },
            effect: 'fade',
            fadeEffect: {
                crossFade: true
            },
            preloadImages: false,
            lazy: {
                loadPrevNext: true,
                loadPrevNextAmount: 1
            },
            pagination: {
                el: '.swiper-pagination',
                clickable: true,
            },
            navigation: {
                nextEl: '.swiper-button-next',
                prevEl: '.swiper-button-prev',
            },
            on: {
                init: function () {
                    let slides = this.slides;
                    slides.forEach((slide, index) => {
                        if (index === this.activeIndex) {
                            slide.style.opacity = 1;
                            slide.classList.add('swiper-slide-active');
                        } else {
                            slide.style.opacity = 0;
                            slide.classList.remove('swiper-slide-active');
                        }
                    });
                },
                slideChange: function () {
                    let slides = this.slides;
                    slides.forEach((slide, index) => {
                        if (index === this.activeIndex) {
                            slide.style.opacity = 1;
                            slide.classList.add('swiper-slide-active');
                        } else {
                            slide.style.opacity = 0;
                            slide.classList.remove('swiper-slide-active');
                        }
                    });
                }
            }
        });

        // Pause autoplay on hover
        document.querySelector('.swiper').addEventListener('mouseenter', () => {
            swiper.autoplay.stop();
        });
        document.querySelector('.swiper').addEventListener('mouseleave', () => {
            swiper.autoplay.start();
        });
    </script>
</body>
</html>