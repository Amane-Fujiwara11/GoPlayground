import React, { useEffect, useState } from 'react';
import './App.css';
import TaskList from './components/TaskList';
import TaskForm from './components/TaskForm';

interface Task {
  id: number;
  title: string;
  content: string;
  completed: boolean;
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

  const updateTaskStatus = async (id: number, status: string) => {
    try {
      await fetch(`http://localhost:8080/tasks/${id}/status`, {
        method: 'PATCH',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ status }),
      });
      setTasks((prevTasks) =>
        prevTasks
          ? prevTasks.map((task) =>
              task.id === id ? { ...task, status } : task
            )
          : null
      );
    } catch (error) {
      console.error('Error updating task status:', error);
    }
  };

  const updateTaskCompletion = async (id: number, completed: boolean) => {
    try {
      await fetch(`http://localhost:8080/tasks/${id}/complete`, {
        method: 'PATCH',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ completed }),
      });
      setTasks((prevTasks) =>
        prevTasks
          ? prevTasks.map((task) =>
              task.id === id ? { ...task, completed } : task
            )
          : null
      );
    } catch (error) {
      console.error('Error updating task completion:', error);
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
      {tasks ? (
        <TaskList
          tasks={tasks}
          onDelete={deleteTask}
          onCompleteChange={updateTaskCompletion}
        />
      ) : (
        <p>Loading tasks...</p>
      )}
    </div>
  );
};

export default App;
