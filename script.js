document.addEventListener("DOMContentLoaded", function() {
    var form = document.querySelector('form');
    form.addEventListener('submit', function(event) {
        event.preventDefault();
        var messageInput = document.getElementById('message');
        var message = messageInput.value.trim();
        if (message !== "") {
            console.log("Item añadido:", message);
            form.submit();
        }
    });
});