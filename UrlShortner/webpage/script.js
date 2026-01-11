// --- CONFIGURATION ---
const API_BASE = "http://16.171.231.55:3000/api/v1"; 

// --- THEME LOGIC ---
const themeBtn = document.getElementById('theme-toggle');
const body = document.body;
const curtain = document.getElementById('theme-curtain');

const savedTheme = localStorage.getItem('theme') || 'dark';
body.setAttribute('data-theme', savedTheme);

themeBtn.addEventListener('click', () => {
    // 1. Start Animation
    curtain.classList.remove('curtain-down');
    void curtain.offsetWidth; // Force Reflow
    curtain.classList.add('curtain-down');

    // 2. WAIT for curtain to cover (400ms)
    setTimeout(() => {
        // DISABLE TRANSITIONS to prevent flicker behind curtain
        body.classList.add('no-transition');

        // SWITCH THEME
        const currentTheme = body.getAttribute('data-theme');
        const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
        body.setAttribute('data-theme', newTheme);
        localStorage.setItem('theme', newTheme);

    }, 400); // Sync with CSS keyframe "40%"

    // 3. CLEANUP after animation finishes (800ms)
    setTimeout(() => {
        body.classList.remove('no-transition');
    }, 800);
});

// --- FORM LOGIC ---
const form = document.getElementById('shorten-form');
const submitBtn = document.getElementById('submitBtn');
const resultArea = document.getElementById('result-area');
const errorArea = document.getElementById('error-area');

let errorTimeout; // To track the timer

form.addEventListener('submit', async (e) => {
    e.preventDefault();

    submitBtn.textContent = "Routing Packets...";
    submitBtn.disabled = true;
    resultArea.classList.add('hidden');
    
    // Clear previous errors immediately
    errorArea.classList.add('hidden');
    errorArea.textContent = "";
    if (errorTimeout) clearTimeout(errorTimeout);

    const longUrl = document.getElementById('longUrl').value;
    const customAlias = document.getElementById('customAlias').value;
    const expiry = document.getElementById('expiry').value;

    const payload = { url: longUrl };
    if (customAlias) payload.customshortner = customAlias;
    if (expiry) payload.expiry = expiry;

    try {
        const response = await fetch(API_BASE, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(payload)
        });

        const data = await response.json();

        if (!response.ok) throw new Error(data.error || "Server Error");

        const fullShortLink = `http://${data.customshort}`;
        
        document.getElementById('finalLink').href = fullShortLink;
        document.getElementById('finalLink').textContent = fullShortLink;
        document.getElementById('rateLimit').textContent = `Quota: ${data.rate_limit} left`;
        document.getElementById('resetTime').textContent = `Resets in: ${data.rate_limit_reset} min`;

        resultArea.classList.remove('hidden');

    } catch (err) {
        errorArea.textContent = `Error: ${err.message}`;
        errorArea.classList.remove('hidden');
        
        // --- AUTO CLEAR ERROR AFTER 10 SECONDS ---
        errorTimeout = setTimeout(() => {
            errorArea.classList.add('hidden');
            errorArea.textContent = "";
        }, 10000); // 10000ms = 10s

    } finally {
        submitBtn.textContent = "Generate Short Link";
        submitBtn.disabled = false;
    }
});