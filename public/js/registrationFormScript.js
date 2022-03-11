let reg = document.getElementById("registration")

function Show_registration_form(title_program, id) {
    let span = document.getElementById("registration_title_program")
    let program_id = document.getElementById("id")
    reg.classList.remove("hide")
    span.textContent = title_program
    program_id.value = id

    document.body.style.overflow = "hidden"
}

function Close_registration_form() {
    reg.classList.add("hide")
    document.body.style.overflow = "auto"
}