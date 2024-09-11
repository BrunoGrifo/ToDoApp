document.querySelectorAll('.delete-button').forEach(button => {
    button.addEventListener('click', function() {
        const taskId = this.getAttribute('data-task-id');
        const csrfToken = this.getAttribute('data-csrf-token');
        console.log(csrfToken)
        console.log("yo")
        if (confirm('Are you sure you want to delete this task?')) {
            fetch(`/task/${taskId}`, {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                    'X-CSRF-Token': csrfToken
                },
            })
            .then(response => {
                if (response.ok) {
                    // Remove the gallery item from the DOM
                    this.parentElement.remove();
                } else {
                    alert('Failed to delete task');
                }
            })
            .catch(error => {
                console.error('Error:', error);
                alert('An error occurred');
            });
        }
    });
});

document.addEventListener('DOMContentLoaded', () => {
    // Select all gallery-item elements
    const galleryItems = document.querySelectorAll('.gallery-item');

    galleryItems.forEach(item => {
        // Add click event listener to each gallery-item
        item.addEventListener('click', () => {
            // Retrieve the task ID from the data attribute
            const taskId = item.getAttribute('data-task-id');

            // Construct the URL for the GET request
            const url = `/task/${taskId}`;

            // Redirect to the URL
            window.location.href = url;
        });
    });
});

document.getElementById('update-button').addEventListener('click', function() {
    const form = document.getElementById('update-task-form');
    const formData = new FormData(form);
    console.log("check");
    const csrfToken = document.getElementById('csrf-token').value;

    fetch('/task', {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
            'X-CSRF-Token': csrfToken
        },
        body: new URLSearchParams(formData).toString()
    })
    .then(response => {
        console.log("check");
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        // Handle successful response
        console.log('Success:', data);
    })
    .catch(error => {
        // Handle error
        console.error('Error:', error);
    });
});