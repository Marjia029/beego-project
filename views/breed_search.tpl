<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat Browser</title>
    <link rel="stylesheet" href="https://unpkg.com/swiper/swiper-bundle.min.css" />
    <link rel="stylesheet" href="/static/css/main.css">
    <link rel="stylesheet" href="/static/css/breed_search.css">
</head>
<body>
    <div class="content-container">
        {{template "navbar" .}}
    
        <div class="search-container">
            <select name="breed_id" onchange="this.form.submit()" form="breed-form">
                {{range .Breeds}}
                    <option value="{{.ID}}" {{if eq .ID $.SelectedBreed.ID}}selected{{end}}>
                        {{.Name}}
                    </option>
                {{end}}
            </select>
            <form id="breed-form" method="POST" action="/breed-search"></form>
        </div>
    
        {{if .SelectedBreed}}
            <div class="swiper">
                <div class="swiper-wrapper">
                    {{range .BreedImages}}
                    <div class="swiper-slide">
                        <img src="{{.URL}}" alt="Breed Image">
                    </div>
                    {{end}}
                </div>
                <div class="swiper-pagination"></div>
            </div>
    
            <div class="breed-info">
                <h2 class="breed-title">
                    {{.SelectedBreed.Name}} 
                    <span class="breed-origin">{{if .SelectedBreed.Origin}}({{.SelectedBreed.Origin}}){{end}}</span>
                    <span class="breed-id">{{.SelectedBreed.ID}}</span>
                </h2>
                <p class="breed-description">{{.SelectedBreed.Description}}</p>
                {{if .SelectedBreed.WikipediaURL}}
                    <a href="{{.SelectedBreed.WikipediaURL}}" target="_blank" class="wiki-link">WIKIPEDIA</a>
                {{end}}
            </div>
            
        {{end}}
    

    </div>
    <script src="https://unpkg.com/swiper/swiper-bundle.min.js"></script>
    <script src="/static/js/breed_search.js"></script>
</body>
</html>
