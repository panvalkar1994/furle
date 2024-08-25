let shortenForm = document.getElementById('shorten-form');

shortenForm.addEventListener('submit', function(e) {
    e.preventDefault();
    let url = document.getElementById('url').value;
    console.log(url);
    let data = {
        url: url
    };

    let body = JSON.stringify(data);
    console.log(body);
    fetch('/api/v1/shorten', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: body
    })
    .then(response => response.json())
    .then(data => {
        showNewShortUrl(data);
        console.log("resp",data);  
    });
});


// TODO: Create a div to show the shortened URL and append it to the DOM
//  along with it should have a copy button
function showNewShortUrl(data) {
    let shortUrl = document.getElementById('shorturl');
    shortUrl.innerHTML = '';
    shortUrl.innerHTML = shortUrlElement(data.shortened);
}

function copyToClipboard() {
    let copyText = document.getElementById('short-url');
    navigator.clipboard.writeText(copyText.href);
    alert('Copied to clipboard');
}

const shortUrlElement = (shorturl)=> `
<div class="shorturl">
    <a href="${shorturl}" id="short-url">${shorturl}</a>
    <button onclick="copyToClipboard()">Copy</button>
</div>
`