// Проверка что человек авторизован. Если нет, то очищаем всё на экране
if (!document.cookie.includes("session_token")) {
    document.getElementById("main_window").remove()
}

// Перекрашивание блока с заголовкам и сатусом программы в нужный цвет
Switch_color_status("Ожидание", "#FFFF00")
Switch_color_status("Активно", "#00FF00")

// Отображение окна в зависимости от вкладки после перезагрузки
Switch_view()

switch_icon_status_listener("Подал заявление", "/assets/images/clock.png")


include("/assets/js/controlPanelProgramListView.js")

function Switch_view() {
    let switch_but = document.getElementsByName("select")
    let control_panel = document.getElementById("control_panel")
    let archive = document.getElementById("archive")
    for (let i of switch_but) {
        if (i.checked == true && i.id == "control_input") {
            archive.classList.add("hide")
            control_panel.classList.remove("hide")
        } else if (i.checked == true && i.id == "library_input") {
            control_panel.classList.add("hide")
            archive.classList.remove("hide")
        }
    }

}

function Create_new_program() {
    let programs = document.querySelectorAll(".DPO")
    for (let j of programs) {
        j.classList.add("hide")
    }
    let create_form = document.getElementById("creat_program_form")
    create_form.classList.remove("hide")
}

function Switch_color_status(status, color) {
    let statuses = document.querySelectorAll(".status_lable")
    let program_status = document.getElementsByClassName(status)
    console.log(program_status.value)
    for (const i of statuses) {
        if (i.textContent == status) {
            i.style.background = color
            i.style.fontSize = "22px"
        }
    }
    for (const i of program_status) {
        i.style.background = color
        i.style["boxShadow"] = "0 0 4px rgba(0, 0, 0, .6) inset"
    }
}

function switch_icon_status_listener(name_value, path) {
    let image_listener_statuses = document.getElementsByClassName(name_value)
    for (let i of image_listener_statuses) {
        i.src = path
    }
}

function include(url) {
    var script = document.createElement('script');
    script.src = url;
    document.getElementsByTagName('head')[0].appendChild(script);
}

function Send(method, uri, data, callback) {
    let xhr = new XMLHttpRequest();
    xhr.open(method, uri);

    xhr.onload = function(event) {
        callback(JSON.parse(this.response));
    }
    if (data) {
        console.log(data);
        xhr.setRequestHeader("Content-Type", "application/json; charset=utf-8");
        xhr.setRequestHeader("X-Requested-With", "XMLHttpRequest");
        xhr.send(JSON.stringify(data));
    } else {
        xhr.send();
    }
}