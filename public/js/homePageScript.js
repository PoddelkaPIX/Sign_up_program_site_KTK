let select_filter = document.getElementsByClassName("select_filter")
for (let select of select_filter) {
    select.onclick = ()=>{
        let drop_down
        if (select.id == "level"){
            drop_down = document.querySelector("#level ~ .drop_down_items")
        }else if (select.id == "type"){
            drop_down = document.querySelector("#type ~ .drop_down_items")
        }else if (select.id == "form"){
            drop_down = document.querySelector("#form ~ .drop_down_items")
        }else if (select.id == "place"){
            drop_down = document.querySelector("#place ~ .drop_down_items")
        }

        if (drop_down.classList.value.includes("hide")){
            drop_down.classList.remove("hide")
        }else{
            drop_down.classList.add("hide")
        }
    }
}



Сhoosing_filter(".level_filter_option", "#level")
Сhoosing_filter(".type_filter_option", "#type")
Сhoosing_filter(".form_filter_option", "#form")
Сhoosing_filter(".place_filter_option", "#place")

function Сhoosing_filter(option, select){
    let filter_option = document.querySelectorAll(option)
    for (let op of filter_option){
        op.onclick = ()=>{
            let drop_down = op.parentElement
            drop_down.classList.add("hide")
            let select_filter= document.querySelector(select)
            select_filter.textContent = op.textContent
            console.log(op.textContent)

        }
    }
}
