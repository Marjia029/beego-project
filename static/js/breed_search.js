document.addEventListener('DOMContentLoaded', function() {
    // Initialize Swiper
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

    // Pause autoplay on hover
    const swiperElement = document.querySelector('.swiper');
    if (swiperElement) {
        swiperElement.addEventListener('mouseenter', () => {
            swiper.autoplay.stop();
        });
        swiperElement.addEventListener('mouseleave', () => {
            swiper.autoplay.start();
        });
    }
});