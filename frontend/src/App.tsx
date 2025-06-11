import React, { useEffect, useState } from 'react';
import './App.css';
import TaskList from './components/TaskList';

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

  return (
    <div className="App">
      <h1>Task Management</h1>
      {tasks ? <TaskList tasks={tasks} /> : <p>Loading tasks...</p>}
    </div>
  );
};

export default App;
