document.addEventListener("DOMContentLoaded", function () {
  const form = document.getElementById("form");
  const taskList = document.getElementById("taskList");

  form.addEventListener("submit", function (e) {
    e.preventDefault();
    const newTask = document.getElementById("newTask").value;
    addTask(newTask);
    form.reset();
  });

  function addTask(task) {
    const li = document.createElement("li");
    li.className = "collection-item";
    li.appendChild(document.createTextNode(task));
    taskList.appendChild(li);
    saveTaskInLocalStorage(task);
  }

  function saveTaskInLocalStorage(task) {
    let tasks;
    if (localStorage.getItem("tasks") === null) {
      tasks = [];
    } else {
      tasks = JSON.parse(localStorage.getItem("tasks"));
    }
    tasks.push(task);
    localStorage.setItem("tasks", JSON.stringify(tasks));
  }

  function getTasks() {
    let tasks;
    if (localStorage.getItem("tasks") === null) {
      tasks = [];
    } else {
      tasks = JSON.parse(localStorage.getItem("tasks"));
    }
    tasks.forEach(function (task) {
      const li = document.createElement("li");
      li.className = "collection-item";
      li.appendChild(document.createTextNode(task));
      taskList.appendChild(li);
    });
  }

  getTasks();
});
