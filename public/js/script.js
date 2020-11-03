const urlInputBox = document.querySelector('#input');
const generateUrlButton = document.querySelector('#submitbutton');
const form = document.querySelector('#form')
const link = document.querySelector('#link')
 
form.addEventListener('submit', (event) => {
    event.preventDefault()
    const url = urlInputBox.value.trim();
    getShortUrl(url);
});
 
function getShortUrl(url) {
    fetch(`/getshorturl`, {
        method: 'POST',
        headers: {
            'Content-Type': undefined
        },
        body: JSON.stringify({
            url
        })
    })
    .then((resp) => resp.json())
    .then((result) => {
        console.log(result)
        if (result.hasOwnProperty('Response')) {
            renderURLs(result['Response']);
        }
    })
    .catch((error) => {
        console.log(error);
    });
}
 
function renderURLs(response) {
    const {
        RedirectURL, ShortURL
    } = response;
    link.innerHTML = ShortURL;
    link.setAttribute('href', RedirectURL);
}