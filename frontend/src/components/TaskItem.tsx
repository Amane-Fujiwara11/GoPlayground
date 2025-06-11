import React from 'react';

interface Task {
  id: number;
  title: string;
  content: string;
}

interface TaskItemProps {
  onDelete: (id: number) => void;
  task: Task;
}

const TaskItem: React.FC<TaskItemProps> = ({ task, onDelete }) => {
  return (
    <div>
      <h3>{task.title}</h3>
      <p>{task.content}</p>
      <button onClick={() => onDelete(task.id)}>Delete</button>
    </div>
  );
};

export default TaskItem;
