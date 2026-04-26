// Minimal front‑end logic (fetch API examples)
document.addEventListener('DOMContentLoaded', () => {
    const app = document.getElementById('app');
    fetch('/api/tasks')
        .then(r => r.json())
        .then(tasks => {
            app.innerHTML = '<pre>' + JSON.stringify(tasks, null, 2) + '</pre>';
        })
        .catch(err => console.error('Error loading tasks:', err));
});
