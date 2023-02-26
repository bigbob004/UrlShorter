function copyFunction() {
    var copyText = document.getElementById("shorted_url");
    copyText.select();
    copyText.setSelectionRange(0, 99999)
    navigator.clipboard.writeText(copyText.value);
    alert("Скопировал текст: " + copyText.value);
}