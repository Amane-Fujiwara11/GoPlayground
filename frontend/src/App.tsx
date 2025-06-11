import React from 'react';
import logo from './logo.svg';
import './App.css';
import TaskList from './components/TaskList';
import TaskForm from './components/TaskForm';

interface Task {
  id: number;
  title: string;
  content: string;
}

const App: React.FC = () => {
  const [tasks, setTasks] = useState<Task[] | null>(null);

  useEffect(() => {
    fetch('http://localhost:8080/tasks')
      .then((response) => response.json())
      .then((data) => setTasks(data));
  }, []);

  const addTask = (title: string, content: string) => {
    fetch('http://localhost:8080/tasks', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ title, content }),
    })
      .then((response) => response.json())
      .then((newTask) => setTasks((prevTasks) => (prevTasks ? [...prevTasks, newTask] : [newTask])));
  };

  const deleteTask = (id: number) => {
    fetch(`http://localhost:8080/tasks/${id}`, {
      method: 'DELETE',
    }).then(() => setTasks((prevTasks) => prevTasks?.filter((task) => task.id !== id) || null));
  };

  return (
    <div className="App">
      <h1>Task Management</h1>
      <TaskForm onAddTask={addTask} />
      {tasks ? <TaskList tasks={tasks} onDelete={deleteTask} /> : <p>Loading tasks...</p>}
    </div>
  );
}

export default App;
