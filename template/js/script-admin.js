function toggleSections(showId) {
    var showSection = document.getElementById(showId);
    var hideSection = document.getElementById(currentSection);

    if (showSection && hideSection && showId !== currentSection) {
        showSection.style.display = 'block';
        hideSection.style.display = 'none';
        currentSection = showId;
    }
}

window.onload = function() {
    var postsSection = document.getElementById('posts');
    if (postsSection) {
        postsSection.style.display = 'block';
        currentSection = 'posts';
    }
}