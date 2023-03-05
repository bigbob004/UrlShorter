async function create_short_url() {
    let og_url = document.getElementById('input_og_url').value
    let url = 'http://localhost:8080/create';
    let data = {
        url: og_url
    };
    let resp = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json;charset=utf-8'
        },
        body: JSON.stringify(data)
    });
    console.log(resp)
    let shorted_url = await resp.text(); // читаем ответ в формате text

    console.log(shorted_url)
}