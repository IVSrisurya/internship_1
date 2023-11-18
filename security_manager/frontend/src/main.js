import "./style.css";
import "./app.css";

import logo from "./assets/images/logo-universal.png";
// import { Greet } from "../wailsjs/go/main/App";
import { Greet, DisableWindowsUpdates, DisableCommandPrompt, DisableFileDownloads, BlockWebsite, SetLockScreenTimeout } from "../wailsjs/go/main/App";



document.getElementById("logo").src = logo;

let nameElement = document.getElementById("name");
let resultElement = document.getElementById("result");
let timeoutInput = document.getElementById("timeoutInput");

window.greet = function () {
    let name = nameElement.value;

    if (name === "") return;

    try {
        Greet(name)
            .then((result) => {
                resultElement.innerText = result;
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};

window.disableWindowsUpdates = function () {
    DisableWindowsUpdates().then((result) => {
        console.log(`disableWindowsUpdates result: ${result}`);
        resultElement.innerText = `disableWindowsUpdates result: ${result}`;
    }
    ).catch((err) => {
        console.error(err);
        resultElement.innerText = `disableWindowsUpdates error: ${err}`;
    });
};

window.disableCommandPrompt = function () {
    DisableCommandPrompt().then((result) => {
        console.log(`disableCommandPrompt result: ${result}`);
        resultElement.innerText = `disableCommandPrompt result: ${result}`;
    }
    ).catch((err) => {
        console.error(err);
        resultElement.innerText = `disableCommandPrompt error: ${err}`;
    });
};

window.disableFileDownloads = function () {
    DisableFileDownloads().then((result) => {
        console.log(`disableFileDownloads result: ${result}`);
        resultElement.innerText = `disableFileDownloads result: ${result}`;
    }
    ).catch((err) => {
        console.error(err);
        resultElement.innerText = `disableFileDownloads error: ${err}`;
    });
    
};

window.blockWebsite = function () {
    const websiteInput = document.getElementById('websiteInput');
    const website = websiteInput.value;
    BlockWebsite(website).then((result) => {
        console.log(`blockWebsite result: ${result}`);
        resultElement.innerText = `blockWebsite result: ${result}`;
    }
    ).catch((err) => {
        console.error(err);
        resultElement.innerText = `blockWebsite error: ${err}`;
    });
    
};

window.setLockScreenTimeout = function () {
    const timeoutInput = document.getElementById('timeoutInput');
    const timeout = parseInt(timeoutInput.value, 10);
    SetLockScreenTimeout(timeout).then((result) => {
        console.log(`setLockScreenTimeout result: ${result}`);
        resultElement.innerText = `setLockScreenTimeout result: ${result}`;
    }
    ).catch((err) => {
        console.error(err);
    });
};
