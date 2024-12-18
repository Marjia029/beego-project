<!DOCTYPE html>
<html>
<head>
    <title>Breed Images</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            text-align: center;
        }
        h1 {
            margin: 20px 0;
        }
        .images-container {
            display: flex;
            flex-wrap: wrap;
            justify-content: center;
            gap: 10px;
            margin-top: 20px;
        }
        .image-item {
            width: 200px;
            height: 200px;
            overflow: hidden;
            border-radius: 10px;
            box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
        }
        .image-item img {
            width: 100%;
            height: 100%;
            object-fit: cover;
        }
    </style>
</head>
<body>
    <h1>Breed Images</h1>
    {{if .Error}}
        <p style="color: red;">{{.Error}}</p>
    {{else}}
        <div class="images-container">
            {{range .BreedImages}}
                <div class="image-item">
                    <img src="{{.URL}}" alt="Breed Image">
                </div>
            {{end}}
        </div>
    {{end}}
</body>
</html>
