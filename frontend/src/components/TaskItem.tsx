import React from 'react';

interface Task {
  id: number;
  title: string;
  content: string;
}

interface TaskItemProps {
  task: Task;
}

const TaskItem: React.FC<TaskItemProps> = ({ task }) => {
  return (
    <div>
      <h3>{task.title}</h3>
      <p>{task.content}</p>
    </div>
  );
};

export default TaskItem;
