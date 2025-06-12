import React, { useEffect, useState } from 'react';
import './App.css';
import TaskList from './components/TaskList';
import TaskForm from './components/TaskForm';

interface Task {
  id: number;
  title: string;
  content: string;
  createdAt: string;
  status: string;
}

const App: React.FC = () => {
  const [tasks, setTasks] = useState<Task[] | null>(null);

  useEffect(() => {
    fetch('http://localhost:8080/tasks')
      .then((response) => response.json())
      .then((data) => setTasks(data));
  }, []);

  const addTask = async (title: string, content: string) => {
    try {
      const response = await fetch('http://localhost:8080/tasks', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ title, content }),
      });
      const newTask = await response.json();
      setTasks((prevTasks) => (prevTasks ? [...prevTasks, newTask] : [newTask]));
    } catch (error) {
      console.error('Error adding task:', error);
    }
  };

  const deleteTask = async (id: number) => {
    try {
      await fetch(`http://localhost:8080/tasks/${id}`, {
        method: 'DELETE',
      });
      setTasks((prevTasks) => prevTasks?.filter((task) => task.id !== id) || null);
    } catch (error) {
      console.error('Error deleting task:', error);
    }
  };

  return (
    <div className="App">
      <h1>Task Management</h1>
      <TaskForm onAddTask={addTask} />
      {tasks ? <TaskList tasks={tasks} onDelete={deleteTask} /> : <p>Loading tasks...</p>}
    </div>
  );
};

export default App;
