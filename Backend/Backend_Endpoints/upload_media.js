const form = document.getElementById('uploadForm');
const status = document.getElementById('status');

form.addEventListener('submit', async (event) => {
    event.preventDefault(); // Prevent page reload

    const formData = new FormData(form);

    console.log("Called to upload meida");

    try {
        const response = await fetch('/upload_image', {
            method: 'POST',
            body: formData,
        });

        if (response.ok) {
            const result = await response.json();
            status.textContent = `Success: ${result.message}`;
        } else {
            const error = await response.json();
            status.textContent = `Error: ${error.error}`;
        }
    } catch (err) {
        status.textContent = `Error: ${err.message}`;
    }
});
