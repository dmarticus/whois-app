const form = document.getElementById('whoisForm');
form.onsubmit = (event) => {
    event.preventDefault();
    const data = document.getElementById('input').value;
    const div = document.getElementById('resultDiv');
    const xhr = new XMLHttpRequest();
    xhr.open('POST', '/whois', true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function() {
        if(this.readyState !== XMLHttpRequest.DONE) {
            return;
        }
        const response = JSON.parse(this.responseText);
        if(response.error) {
            div.innerHTML = `<p>${response.error}</p>`;
        } else {
            div.innerHTML = `<pre>${response.result}</pre>`;
        }
    };
    xhr.send(`data=${encodeURIComponent(data)}`);
}
