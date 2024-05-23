document.addEventListener('DOMContentLoaded', function() {
    const elem = document.getElementById("sidebar").children[0]

    activeTab(elem)
    elem.classList.add("bg-slate-200")(elem)
});

function toggleMenu() {
    var elem = document.getElementById("panel")
    if (elem.classList.contains("toggle-menu")) {
        elem.classList.remove("toggle-menu")
        return
    }
    elem.classList.add("toggle-menu")
}

function onClickSidebar(elem) {
    const elems = document.getElementById("sidebar").children
    Array.from(elems).forEach(element => {
        if (element.classList.contains("bg-slate-200")) {
            element.classList.remove("bg-slate-200")
        }
    });

    
    elem.classList.add("bg-slate-200")
    activeTab(elem)

    /* close panel */
    toggleMenu()
}

async function activeTab(elem) {
    let tabIndex = elem.getAttribute("index")
    let titleElem = document.getElementById("title")
    titleElem.innerHTML = elem.innerText.toUpperCase()

    switch (parseInt(tabIndex)) {
        case 0:
            /* team collection */
            let teams = await getTeamCollection()
            document.getElementById("collection").innerHTML = teams
            break
        case 1:
            /* user collection */
            let collections = await getUserCollection("tmp12345")
            document.getElementById("collection").innerHTML = collections
            break
        case 2:
            /* guide */
            break
    }
}