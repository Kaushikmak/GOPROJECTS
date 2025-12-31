// Copies the git clone command to clipboard with visual feedback.
function copyCommand() {
    const commandText = document.getElementById("clone-command").innerText;
    const copyBtn = document.getElementById("copyBtn");
    const originalText = copyBtn.innerText;

    navigator.clipboard.writeText(commandText).then(() => {
        copyBtn.innerText = "Copied!";
        copyBtn.style.backgroundColor = "var(--accent)";
        copyBtn.style.color = "#fff";
        copyBtn.style.borderColor = "var(--accent)";
        
        setTimeout(() => {
            copyBtn.innerText = originalText;
            copyBtn.style.backgroundColor = ""; 
            copyBtn.style.color = "";
            copyBtn.style.borderColor = "";
        }, 2000);
    }).catch(err => {
        console.error('Failed to copy text: ', err);
    });
}

// Handles switching between OS installation tabs.
function switchOS(os) {
    document.querySelectorAll('.install-content').forEach(c => c.classList.remove('active'));
    document.querySelectorAll('.tab-btn').forEach(b => b.classList.remove('active'));
    document.getElementById(`install-${os}`).classList.add('active');
    
    const indexMap = { 'linux': 0, 'windows': 1, 'mac': 2, 'source': 3 };
    document.querySelectorAll('.tab-btn')[indexMap[os]].classList.add('active');

    const titles = {
        'linux': 'bash — linux',
        'windows': 'powershell — administrator',
        'mac': 'zsh — macos',
        'source': 'bash — developer'
    };
    document.getElementById('term-title').innerText = titles[os];
}

// ==========================================
//  INTERACTIVE DEMO LOGIC
// ==========================================

const demoScenarios = {
    add: {
        command: 'task-cli add "Finish the project documentation"',
        output: `<div class="output">Task added successfully (ID: 1)</div>`
    },
    list: {
        command: 'task-cli list',
        output: `<div class="output" style="white-space: pre; font-family: monospace;">
ID       STATUS       CREATED           DESCRIPTION
1        <span style="color:#ffbd2e">todo</span>         2025-01-15 10:00  Buy groceries
2        <span style="color:#27c93f">done</span>         2025-01-14 18:30  Setup Go environment
3        <span style="color:#ff5f56">in-progress</span>  2025-01-15 09:15  Write unit tests</div>`
    },
    update: {
        command: 'task-cli update 1 "Buy groceries and snacks"',
        output: `<div class="output">Task updated: 1</div>`
    },
    mark: {
        command: 'task-cli mark 1 in-progress',
        output: `<div class="output">Task 1 marked as in-progress</div>`
    },
    delete: {
        command: 'task-cli delete 2',
        output: `<div class="output">Task deleted: 2</div>`
    }
};

let currentDemoKey = 'add';
let isTyping = false;
let autoPlayInterval;
let demoOrder = ['add', 'list', 'update', 'mark', 'delete'];
let currentOrderIndex = 0;
let isManualMode = false;

document.addEventListener('DOMContentLoaded', () => {
    runDemoSequence('add');
});

function manualDemo(key) {
    isManualMode = true; // Stop auto-looping
    runDemoSequence(key);
}

function runDemoSequence(key) {
    if (isTyping && currentDemoKey === key) return; // Ignore if already running same demo

    // Update Buttons UI
    document.querySelectorAll('.demo-btn').forEach(btn => btn.classList.remove('active'));
    const btnIndex = demoOrder.indexOf(key);
    if (btnIndex >= 0) {
        document.querySelectorAll('.demo-btn')[btnIndex].classList.add('active');
    }

    currentDemoKey = key;
    isTyping = true;

    const terminalBody = document.getElementById('demo-content');
    const typewriterSpan = document.getElementById('typewriter-text');
    
    // Reset Terminal
    terminalBody.innerHTML = `
        <div class="line">
            <span class="prompt">user@dev:~$</span> 
            <span id="typewriter-text"></span><span class="cursor-blink">_</span>
        </div>
    `;

    const scenario = demoScenarios[key];
    const textToType = scenario.command;
    let charIndex = 0;

    // Clear any previous timeouts
    if (window.typewriterTimeout) clearTimeout(window.typewriterTimeout);
    if (window.nextStepTimeout) clearTimeout(window.nextStepTimeout);

    function type() {
        const targetSpan = document.getElementById('typewriter-text');
        if (!targetSpan) return;

        if (charIndex < textToType.length) {
            targetSpan.textContent += textToType.charAt(charIndex);
            charIndex++;
            // Randomize typing speed slightly for realism
            window.typewriterTimeout = setTimeout(type, Math.random() * 50 + 50); 
        } else {
            // Typing finished, show output after delay
            window.nextStepTimeout = setTimeout(() => {
                showOutput(scenario.output);
            }, 600);
        }
    }

    type();
}

function showOutput(outputHtml) {
    const terminalBody = document.getElementById('demo-content');
    
    // Remove cursor from current line
    const cursor = terminalBody.querySelector('.cursor-blink');
    if (cursor) cursor.remove();

    // Append Output
    const outputContainer = document.createElement('div');
    outputContainer.innerHTML = outputHtml;
    terminalBody.appendChild(outputContainer);

    // Append New Prompt Line
    const newPrompt = document.createElement('div');
    newPrompt.className = 'line';
    newPrompt.innerHTML = `<span class="prompt">user@dev:~$</span> <span class="cursor-blink">_</span>`;
    terminalBody.appendChild(newPrompt);

    isTyping = false;

    // If not manual mode, queue next demo
    if (!isManualMode) {
        window.nextStepTimeout = setTimeout(() => {
            currentOrderIndex = (currentOrderIndex + 1) % demoOrder.length;
            runDemoSequence(demoOrder[currentOrderIndex]);
        }, 3000); // Wait 3 seconds before next demo
    }
}