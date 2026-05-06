async function update() {
    const res = await fetch('/api/me', { method: 'POST' });
    const { name } = await res.json();
    const userEl = document.getElementById('user');
    const msgEl = document.getElementById('message');
    
    if (name) {
        userEl.textContent = `Hi, ${name} `;
        msgEl.textContent = 'Logged in';
        document.getElementById('name').style.display = 'none';
        document.getElementById('login').style.display = 'none';
        document.getElementById('logout').style.display = 'inline';
    } else {
        userEl.textContent = '';
        msgEl.textContent = 'login - message';
        document.getElementById('name').style.display = 'inline';
        document.getElementById('login').style.display = 'inline';
        document.getElementById('logout').style.display = 'none';
    }
}

document.getElementById('login').onclick = async () => {
    const name = document.getElementById('name').value;
    const res = await fetch('/api/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name })
    });
    if (res.ok) {
        update();
    }
};

document.getElementById('logout').onclick = async () => {
    const res = await fetch('/api/logout', { method: 'POST' });
    if (res.ok) {
        update();
    }
};

update();
