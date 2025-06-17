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
  onCompleteChange: (id: number, completed: boolean) => void;
}


const TaskList: React.FC<TaskListProps> = ({ tasks, onDelete, onStatusChange, onCompleteChange }) => {
  return (
    <div className="TaskCards">
      {tasks.map((task) => (
        <div key={task.id} className={`TaskCard ${task.status}`}>
          <h3>{task.title}</h3>
          <p>{task.content}</p>
            {task.status === "registered" && (
              <button
                className="StartButton"
                onClick={() => onStatusChange(task.id, "doing")}
              >
                着手
              </button>
            )}
            <input
              type="checkbox"
              className="TaskCheckbox"
              checked={task.status === "completed"}
              onChange={(e) => onCompleteChange(task.id, e.target.checked)}
            />
          <button onClick={() => onDelete(task.id)}>Delete</button>
        </div>
      ))}
    </div>
  );
};

export default TaskList;
