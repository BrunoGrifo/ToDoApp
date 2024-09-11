document.addEventListener('DOMContentLoaded', function() {
    const submitButton = document.getElementById('submit-button');
    const form = document.getElementById('create-task-form');

    if (submitButton && form) {
        submitButton.addEventListener('click', function(event) {
            event.preventDefault();  // Prevent default form submission

            const formData = new FormData(form);
            console.log("check");

            const title = formData.get('title');
            const description = formData.get('description');
            const csrfToken = document.getElementById('csrf-token').value;

            const data = new URLSearchParams({
                title: title,
                description: description
            }).toString();

            console.log(data);

            fetch('/task', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                    'X-CSRF-Token': csrfToken
                },
                body: data
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
                // Optionally, handle success by updating the UI or redirecting
            })
            .catch(error => {
                // Handle error
                console.error('Error:', error);
            });
        });
    }
});
