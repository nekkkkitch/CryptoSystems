import './style.css';
import './app.css';

import { SetPublic } from '../wailsjs/go/app/App';
import { GetBig } from '../wailsjs/go/app/App';
import { CalcZ } from '../wailsjs/go/app/App';

let BigA = document.getElementById("big-a")
let BigB = document.getElementById("big-b")
let ZA = document.getElementById("z-a")
let ZB = document.getElementById("z-b")

window.setPublic = function () {
     try{
            let b = document.getElementById("input-b").value
            let m = document.getElementById("input-m").value
            SetPublic(Number(b), Number(m))
        }catch(err){
            console.error(err)
        }
}

window.sendA = function(){
    try {
        let a = document.getElementById("input-a-small").value
        GetBig(Number(a)).then((result) => {
            console.log(result)
            BigA.innerHTML = result
        })
    }catch(err){
        console.error(err)
    }
}

window.sendB = function(){
    try {
        let b = document.getElementById("input-b-small").value
        GetBig(Number(b)).then((result) => {
            console.log(result)
            BigB.innerHTML = result
        })
    }catch(err){
        console.error(err)
    }
}

window.calcZA = function(){
    try {
        let bigB = document.getElementById("big-b").innerHTML
        let smallA = document.getElementById("input-a-small").value
        CalcZ(Number(bigB), Number(smallA)).then((result) => {
            console.log(result)
            ZA.innerHTML = result
        })
    }catch(err){
        console.error(err)
    }
}

window.calcZB = function(){
    try {
        let bigA = document.getElementById("big-a").innerHTML
        let smallB = document.getElementById("input-b-small").value
        CalcZ(Number(bigA), Number(smallB)).then((result) => {
            console.log(result)
            ZB.innerHTML = result
        })
    }catch(err){
        console.error(err)
    }
}