ocument.addEventListener('DOMContentLoaded', function() {
    const taskInput = document.getElementById('taskInput');
    const addTaskButton = document.getElementById('addTaskButton');
    const taskList = document.getElementById('taskList');

    addTaskButton.addEventListener('click', addTask);

    function addTask() {
        const taskText = taskInput.value.trim();
        if (taskText !== '') {
            const taskItem = createTaskItem(taskText);
            taskList.appendChild(taskItem);
            taskInput.value = '';
        }
    }

    function createTaskItem(taskText) {
        const taskItem = document.createElement('li');
        taskItem.classList.add('task');

        const completeCheckbox = document.createElement('input');
        completeCheckbox.type = 'checkbox';
        completeCheckbox.classList.add('task-complete');
        taskItem.appendChild(completeCheckbox);

        const taskSpan = document.createElement('span');
        taskSpan.textContent = taskText;
        taskItem.appendChild(taskSpan);

        const deleteButton = document.createElement('button');
        deleteButton.textContent = 'Delete';
        deleteButton.classList.add('delete-task');
        taskItem.appendChild(deleteButton);

        deleteButton.addEventListener('click', function() {
            taskItem.remove();
        });

        completeCheckbox.addEventListener('change', function() {
            if (completeCheckbox.checked) {
                taskSpan.style.textDecoration = 'line-through';
            } else {
                taskSpan.style.textDecoration = 'none';
            }
        });

        return taskItem;
    }
});