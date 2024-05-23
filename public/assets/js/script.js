function toggleMenu() {
    var elem = document.getElementById("sidebar")
    if (elem.classList.contains("toggle-menu")) {
        elem.classList.remove("toggle-menu")
        return
    }
    elem.classList.add("toggle-menu")
}