document.addEventListener('DOMContentLoaded', function() {
    const todoForm = document.getElementById('todo-form');
    const todoTitle = document.getElementById('todo-title');
    const todoList = document.getElementById('todo-list');
    const logoutButton = document.getElementById('logout-button');

    function loadTodos() {
        fetch('/todos')
            .then(response => response.json())
            .then(data => {
                todoList.innerHTML = '';
                data.forEach(todo => {
                    const todoItem = document.createElement('div');
                    todoItem.className = 'todo-item';
                    todoItem.innerHTML = `
                        <span>${todo.title}</span>
                        <button onclick="deleteTodo(${todo.ID})">Delete</button>
                    `;
                    todoList.appendChild(todoItem);
                });
            })
            .catch(error => console.error('Error fetching todos:', error));
    }

    todoForm.addEventListener('submit', function(event) {
        event.preventDefault();
        fetch('/todos', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ title: todoTitle.value, status: false })
        })
        .then(response => response.json())
        .then(data => {
            loadTodos();
            todoTitle.value = '';
        })
        .catch(error => console.error('Error creating todo:', error));
    });

    window.deleteTodo = function(id) {
        const todoItem = document.querySelector(`.todo-item button[onclick="deleteTodo(${id})"]`).parentElement;
        todoItem.classList.add('hidden');
        setTimeout(() => {
            fetch(`/todos/${id}`, {
                method: 'DELETE'
            })
            .then(response => response.json())
            .then(data => {
                loadTodos();
            })
            .catch(error => console.error('Error deleting todo:', error));
        }, 500);
    };

    logoutButton.addEventListener('click', function() {
        fetch('/logout', {
            method: 'POST'
        })
        .then(response => {
            if (response.redirected) {
                window.location.href = response.url;
            }
        })
        .catch(error => console.error('Error logging out:', error));
    });

    loadTodos();
});
