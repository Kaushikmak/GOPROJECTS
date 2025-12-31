/**
 * Copies the installation command to the system clipboard.
 */
function copyCommand() {
    const commandText = document.getElementById("clone-command").innerText;
    const copyBtn = document.getElementById("copyBtn");

    navigator.clipboard.writeText(commandText).then(() => {
        const originalText = copyBtn.innerText;
        copyBtn.innerText = "Copied!";
        copyBtn.style.backgroundColor = "var(--accent)";
        copyBtn.style.color = "#fff";
        
        setTimeout(() => {
            copyBtn.innerText = originalText;
            copyBtn.style.backgroundColor = ""; 
            copyBtn.style.color = "";
        }, 2000);
    }).catch(err => {
        console.error('Failed to copy text: ', err);
    });
}

/**
 * Switches the displayed installation instructions based on OS selection.
 */
function switchOS(os) {
    // 1. Hide all content
    document.querySelectorAll('.install-content').forEach(c => c.classList.remove('active'));

    // 2. Deactivate all buttons
    document.querySelectorAll('.tab-btn').forEach(b => b.classList.remove('active'));

    // 3. Show selected content
    document.getElementById(`install-${os}`).classList.add('active');

    // 4. Activate selected button
    const buttons = document.querySelectorAll('.tab-btn');
    if (os === 'linux') buttons[0].classList.add('active');
    if (os === 'windows') buttons[1].classList.add('active');
    if (os === 'mac') buttons[2].classList.add('active');
    if (os === 'source') buttons[3].classList.add('active');

    // 5. Update Terminal Title
    const titles = {
        'linux': 'bash — linux',
        'windows': 'powershell — administrator',
        'mac': 'zsh — macos',
        'source': 'bash — developer'
    };
    document.getElementById('term-title').innerText = titles[os];
}

console.log("TaskTracker Landing Page | Loaded Successfully");