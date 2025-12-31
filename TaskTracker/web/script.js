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
    // Hide all content & deactivate buttons
    document.querySelectorAll('.install-content').forEach(c => c.classList.remove('active'));
    document.querySelectorAll('.tab-btn').forEach(b => b.classList.remove('active'));

    // Show selected content & activate button
    document.getElementById(`install-${os}`).classList.add('active');
    
    // Map OS to button index
    const indexMap = { 'linux': 0, 'windows': 1, 'mac': 2, 'source': 3 };
    document.querySelectorAll('.tab-btn')[indexMap[os]].classList.add('active');

    // Update Terminal Title
    const titles = {
        'linux': 'bash — linux',
        'windows': 'powershell — administrator',
        'mac': 'zsh — macos',
        'source': 'bash — developer'
    };
    document.getElementById('term-title').innerText = titles[os];
}