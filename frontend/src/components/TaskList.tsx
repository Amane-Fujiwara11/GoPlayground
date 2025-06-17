import React from 'react';

interface Task {
  id: number;
  title: string;
  content: string;
  status: string;
}

interface TaskListProps {
  tasks: Task[];
  onDelete: (id: number) => void;
  onStatusChange: (id: number, status: string) => void;
}


const TaskList: React.FC<TaskListProps> = ({ tasks, onDelete, onStatusChange }) => {
  return (
    <div className="TaskCards">
      {tasks.map((task) => (
        <div key={task.id} className="TaskCard">
          <h3>{task.title}</h3>
          <p>{task.content}</p>
          <select
            value={task.status}
            onChange={(e) => onStatusChange(task.id, e.target.value)}
            style={{
              appearance: "none",
              border: "1px solid #ccc",
              padding: "5px",
              borderRadius: "4px",
              backgroundColor: "#fff",
              cursor: "pointer",
            }}
          >
            <option value="registered">未着手</option>
            <option value="doing">進行中</option>
            <option value="completed">完了</option>
          </select>
          <button onClick={() => onDelete(task.id)}>Delete</button>
        </div>
      ))}
    </div>
  );
};

export default TaskList;
