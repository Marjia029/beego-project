{{define "navbar"}}
<nav class="nav-tabs">
    <a href="/voting" class="nav-item {{if eq .ActiveTab "voting"}}active{{end}}">
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M12 20V4M4 12l8-8 8 8"/> <!-- Upward arrow -->
        </svg>
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M12 4v16M4 12l8 8 8-8"/> <!-- Downward arrow -->
        </svg>
        Voting
    </a>
    <a href="/breed-search" class="nav-item {{if eq .ActiveTab "breeds"}}active{{end}}">
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <circle cx="11" cy="11" r="8"/>
            <path d="M21 21l-4.35-4.35"/>
        </svg>
        Breeds
    </a>
    <a href="#favs-tab" class="nav-item {{if eq .ActiveTab "favs"}}active{{end}}" id="favs-tab">
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78L12 21.23l8.84-8.84a5.5 5.5 0 0 0 0-7.78z"/>
        </svg>
        Favs
    </a>
</nav>
{{end}}