let show_status = document.getElementById("show_status");
let box_menu = document.querySelector(".box_menu");
let login_but = document.getElementById("login_but");


if (document.cookie.includes("session_token")) {
    show_status.textContent = "Вы Администратор"
    document.getElementById("admin_panel").classList.remove("hide")
    document.getElementById("authorization").classList.add("hide")
} else {
    show_status.textContent = "Вы гость"
    document.getElementById("authorization").classList.remove("hide")
    document.getElementById("admin_panel").classList.add("hide")
}

function show_menu() {
    let class_menu = box_menu.classList.value
    if (class_menu.includes("hide")) {
        box_menu.classList.remove("hide")
    } else {
        box_menu.classList.add("hide")
    }
}

login_but.onclick = function() {
    let is_correct = true
    let login_input = document.getElementById("login")
    let password_input = document.getElementById("password")
    if (!login_input || login_input.value.length != 10) {
        login_input.classList.add("incorrect")
        is_correct = false
    }
    if (!password_input || password_input.value.length < 8) {
        password_input.classList.add("incorrect")
        is_correct = false
    }
    if (!is_correct) {
        return
    }

    let data = JSON.stringify({
        Login: login_input.value,
        Password: password_input.value
    })

    let xhr = new XMLHttpRequest()
    xhr.open("POST", "/login")
    xhr.send(data)
    console.log("Отправлено: ", login_input.value, password_input.value)

}