document.addEventListener('DOMContentLoaded', () => {
    const searchForm = document.querySelector('form[action="/search"]');
    if (searchForm) {
        searchForm.addEventListener('submit', function(event) {
            event.preventDefault();
            const query = this.querySelector('input[name="query"]').value.trim();
            if (query) {
                window.location.href = `/search?query=${encodeURIComponent(query)}`;
            } else {
                alert('Please enter a search query.');
            }
        });
    }
});