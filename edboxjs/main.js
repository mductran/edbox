import {Terminal} from "@xterm/xterm";
import {FitAddon} from "@xterm/addon-fit"

const term = new Terminal();
const fitAddon = new FitAddon();

term.loadAddon(fitAddon);
term.open(document.getElementById("terminal"));

fitAddon.fit();
window.addEventListener("resize", () => fitAddon.fit());
