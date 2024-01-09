function toggleCompletion(item) {
    item.classList.toggle("completed");
}

function addItem() {
    var newItemInput = document.getElementById("newItemInput");
    var newItemText = newItemInput.value.trim();

    if (newItemText !== "") {
        var checklist = document.getElementById("checklist");
        var newItem = document.createElement("li");
        newItem.className = "item";
        newItem.textContent = newItemText;
        newItem.onclick = function () {
            toggleCompletion(this);
        };

        checklist.appendChild(newItem);
        newItemInput.value = "";
    }
}