// --- CONFIGURATION ---
// Use your Ngrok URL here
const API_BASE = "https://kinematographic-supercolossally-hollie.ngrok-free.dev/api/v1";

// --- THEME LOGIC ---
const themeBtn = document.getElementById('theme-toggle');
const body = document.body;
const savedTheme = localStorage.getItem('theme') || 'light'; 
body.setAttribute('data-theme', savedTheme);
updateThemeIcon(savedTheme);

themeBtn.addEventListener('click', () => {
    const currentTheme = body.getAttribute('data-theme');
    const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
    body.setAttribute('data-theme', newTheme);
    localStorage.setItem('theme', newTheme);
    updateThemeIcon(newTheme);
});

function updateThemeIcon(theme) {
    const sun = document.querySelector('.icon-sun');
    const moon = document.querySelector('.icon-moon');
    if (theme === 'dark') {
        sun.style.display = 'block';
        moon.style.display = 'none';
    } else {
        sun.style.display = 'none';
        moon.style.display = 'block';
    }
}

// --- AUTOCOMPLETE LOGIC ---
const longUrlInput = document.getElementById('longUrl');
const suggestionsBox = document.getElementById('suggestions');

function getHistory() { return JSON.parse(localStorage.getItem('packetStreamHistory') || '[]'); }

function saveToHistory(original, short) {
    let history = getHistory();
    history = history.filter(item => item.original !== original);
    history.unshift({ original, short });
    if (history.length > 5) history.pop();
    localStorage.setItem('packetStreamHistory', JSON.stringify(history));
}

function showSuggestions(query) {
    const history = getHistory();
    suggestionsBox.innerHTML = '';
    if (history.length === 0) return;

    const matches = query 
        ? history.filter(item => item.original.toLowerCase().includes(query.toLowerCase()))
        : history; 

    if (matches.length === 0) {
        suggestionsBox.classList.add('hidden');
        return;
    }

    matches.slice(0, 5).forEach(item => {
        const div = document.createElement('div');
        div.className = 'suggestion-item';
        div.innerHTML = `<span class="s-original">${item.original}</span><span class="s-short">${item.short}</span>`;
        div.addEventListener('click', () => {
            longUrlInput.value = item.original;
            suggestionsBox.classList.add('hidden');
        });
        suggestionsBox.appendChild(div);
    });
    suggestionsBox.classList.remove('hidden');
}

longUrlInput.addEventListener('input', (e) => showSuggestions(e.target.value));
longUrlInput.addEventListener('focus', () => showSuggestions(longUrlInput.value));
document.addEventListener('click', (e) => {
    if (!longUrlInput.contains(e.target) && !suggestionsBox.contains(e.target)) {
        suggestionsBox.classList.add('hidden');
    }
});

// --- FORM SUBMISSION (WITH ROBUST ERROR HANDLING) ---
const form = document.getElementById('shorten-form');
const submitBtn = document.getElementById('submitBtn');
const resultArea = document.getElementById('result-area');
const errorArea = document.getElementById('error-area');
const resultInput = document.getElementById('resultInput');
const copyBtn = document.getElementById('copyBtn');
const openBtn = document.getElementById('openBtn');
const latencyDisplay = document.getElementById('latencyDisplay');

let errorTimeout;

form.addEventListener('submit', async (e) => {
    e.preventDefault();

    // Reset UI
    const originalText = submitBtn.innerText;
    submitBtn.innerText = "PROCESSING...";
    submitBtn.disabled = true;
    resultArea.classList.add('hidden');
    errorArea.classList.add('hidden');
    suggestionsBox.classList.add('hidden');
    if (errorTimeout) clearTimeout(errorTimeout);

    const payload = { 
        url: longUrlInput.value,
        customshortner: document.getElementById('customAlias').value || undefined,
        expiry: document.getElementById('expiry').value || undefined
    };

    try {
        const startTime = performance.now();

        const response = await fetch(API_BASE, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(payload)
        });

        // --- 1. SAFE RESPONSE PARSING ---
        // Check if the response is actually JSON before parsing
        const contentType = response.headers.get("content-type");
        let data = {};

        if (contentType && contentType.includes("application/json")) {
            data = await response.json();
        } else {
            // Server returned HTML/Text (e.g., 404 page or Nginx error)
            // Log raw text for debugging, but throw a clean error for user
            const text = await response.text();
            console.error("Non-JSON Response received:", text);
            throw new Error("INVALID_SERVER_RESPONSE");
        }

        // --- 2. HTTP STATUS HANDLING ---
        if (!response.ok) {
            let userMsg = "UNKNOWN_SYSTEM_ERROR";
            
            // Map status codes to "Neobrutalist" messages
            switch (response.status) {
                case 400: userMsg = "INVALID_INPUT_DETECTED"; break;
                case 404: userMsg = "API_ENDPOINT_NOT_FOUND"; break;
                case 409: userMsg = "ALIAS_ALREADY_TAKEN"; break; // Conflict
                case 422: userMsg = "UNPROCESSABLE_ENTITY"; break;
                case 429: userMsg = "RATE_LIMIT_EXCEEDED (WAIT)"; break;
                case 500: userMsg = "INTERNAL_SERVER_ERROR"; break;
                case 502: userMsg = "BAD_GATEWAY (SERVER_DOWN)"; break;
                default: userMsg = data.error ? data.error.replace(/ /g, '_').toUpperCase() : `ERROR_${response.status}`;
            }
            throw new Error(userMsg);
        }

        // --- 3. SUCCESS ---
        const endTime = performance.now();
        const latency = (endTime - startTime).toFixed(0);

        const fullShortLink = `http://${data.customshort}`;
        resultInput.value = fullShortLink;
        
        latencyDisplay.innerHTML = `${latency}ms`;
        document.getElementById('rateLimit').textContent = data.rate_limit || "∞";
        
        saveToHistory(payload.url, fullShortLink);
        resultArea.classList.remove('hidden');

    } catch (err) {
        // --- 4. NETWORK & GENERIC ERROR HANDLING ---
        let finalMsg = err.message;
        
        // Handle "Failed to fetch" (Offline or Server Down)
        if (err.name === 'TypeError' && err.message === 'Failed to fetch') {
            finalMsg = "CONNECTION_REFUSED (IS_SERVER_ONLINE?)";
        }

        errorArea.innerHTML = `[FATAL]: ${finalMsg}`;
        errorArea.classList.remove('hidden');

        // Auto-hide error after 5 seconds
        errorTimeout = setTimeout(() => {
            errorArea.classList.add('hidden');
        }, 5000);

    } finally {
        submitBtn.innerText = originalText;
        submitBtn.disabled = false;
    }
});

copyBtn.addEventListener('click', () => {
    resultInput.select();
    navigator.clipboard.writeText(resultInput.value);
    const originalHTML = copyBtn.innerHTML;
    copyBtn.innerHTML = `<i class="ri-check-line"></i>`;
    setTimeout(() => { copyBtn.innerHTML = originalHTML; }, 1500);
});

openBtn.addEventListener('click', () => {
    if(resultInput.value) window.open(resultInput.value, '_blank');
});