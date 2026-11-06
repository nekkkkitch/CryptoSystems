import './style.css';
import './app.css';

import { ChangeParams } from '../wailsjs/go/app/App';
import { Cipher } from '../wailsjs/go/app/App';
import { Decipher } from '../wailsjs/go/app/App';

window.changeParams = function () {
    try{
        let p = document.getElementById("input-p").value
        let q = document.getElementById("input-q").value

        let pub = document.getElementById("public")
        let priv = document.getElementById("private")
        
        ChangeParams(Number(p), Number(q)).then((result) => {
            let pu = result["pu"]
            let pr = result["pr"] 
            pub.innerHTML = "Public: " + pu
            priv.innerHTML = "Private: " + pr 
        })
    }catch(err){
        console.error(err)
    }
}

window.cipher = function () {
    let input = document.getElementById("input-to-cipher").value
    let ciphered = document.getElementById("input-to-decipher")

    // Call App.Greet(name)
    try {
        Cipher(input)
            .then((result) => {
                // Update result with data back from App.Greet()
                ciphered.value = result;
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};

window.decipher = function () {
    let input = document.getElementById("input-to-decipher").value
    let deciphered = document.getElementById("input-to-cipher")

    // Call App.Greet(name)
    try {
        Decipher(input)
            .then((result) => {
                // Update result with data back from App.Greet()
                deciphered.value = result;
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};
