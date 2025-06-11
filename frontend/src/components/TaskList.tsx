import React from 'react';
import TaskItem from './TaskItem';

interface Task {
  id: number;
  title: string;
  content: string;
}

interface TaskListProps {
  tasks: Task[];
  onDelete: (id: number) => void;
}


const TaskList: React.FC<TaskListProps> = ({ tasks, onDelete }) => {
  return (
    <div>
      {tasks.map((task) => (
        <TaskItem key={task.id} task={task} onDelete={onDelete} />
      ))}
    </div>
  );
};

export default TaskList;
