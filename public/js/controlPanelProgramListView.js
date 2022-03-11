let confirm = document.getElementsByClassName("confirm_listener_but")
if (confirm) {
    for (let con of confirm) {
        con.onclick = function() {
            Send("POST", "/confirm/listener", con.dataset.id, response => {
                console.log("Колбек")
                let img_status = document.getElementById("img_status_listener: " + con.dataset.id)
                img_status.remove()
            });
        }
    }
}


let programs = document.getElementsByName("title_program")
if (programs) {
    for (let prog of programs) {
        prog.onclick = function() {
            // Уберает checked при повторном нажатии с radioButton
            var siblings = document.querySelectorAll("input[type='radio'][name='" + prog.name + "']");
            for (var i = 0; i < siblings.length; i++) {
                if (siblings[i] != prog)
                    siblings[i].oldChecked = false;
            }
            if (prog.oldChecked)
                prog.checked = false;
            prog.oldChecked = prog.checked;
            //
            // Уберает class hide у выбраной программы для вывода на экран
            if (prog.checked) {
                let programs = document.querySelectorAll(".DPO")
                for (let j of programs) {
                    if (j.id != prog.value) {
                        j.classList.add("hide")
                    } else {
                        j.classList.remove("hide")
                    }
                }
            }
            //
        }
    }
}

function Show_select_passport() {
    const selected_program = document.querySelectorAll('input[name="title_program"]')
    for (const i of selected_program) {
        if (i.checked) {
            let programs = document.querySelectorAll(".DPO")
            for (let j of programs) {
                if (j.id != i.value) {
                    j.classList.add("hide")
                } else {
                    j.classList.remove("hide")
                }
            }
        }
    }
}