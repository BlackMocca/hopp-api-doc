document.addEventListener('DOMContentLoaded', function() {
    const elem = document.getElementById("sidebar").children[0]

    activeTab(elem)
    elem.classList.add("bg-slate-200")
    elem.classList.add("text-[#000000]")
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
        if (element.classList.contains("text-[#000000]")) {
            element.classList.remove("text-[#000000]")
        }
    });

    
    elem.classList.add("bg-slate-200")
    elem.classList.add("text-[#000000]")
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
            let collections = await getUserCollection()
            document.getElementById("collection").innerHTML = collections
            break
        case 2:
            /* guide */
            break
    }
}

async function logout() {
    await signout()

    window.location.href = "/login"
}

function triggerFileInput(elementId) {
    document.getElementById(elementId).click()
}

async function onChangeImportCollection(elem) {
    if (elem.files.length > 0) {
        const file = elem.files[0]
        if (file) {
            let mycollections = await importCollection(file)
            document.getElementById("collection").innerHTML = mycollections
        }
    } 
}
