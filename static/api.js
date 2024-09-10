document.querySelectorAll('.delete-button').forEach(button => {
    button.addEventListener('click', function() {
        const taskId = this.getAttribute('data-task-id');
        const csrfToken = this.getAttribute('data-csrf-token');
        console.log(csrfToken)
        console.log("yo")
        if (confirm('Are you sure you want to delete this task?')) {
            fetch(`/task?id=${taskId}`, {
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