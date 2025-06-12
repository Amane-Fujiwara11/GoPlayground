import React from 'react';

interface Task {
  id: number;
  title: string;
  content: string;
  completed: boolean;
}

interface TaskListProps {
  tasks: Task[];
  onDelete: (id: number) => void;
  onCompleteChange: (id: number, completed: boolean) => void;
}


const TaskList: React.FC<TaskListProps> = ({ tasks, onDelete, onCompleteChange }) => {
  return (
    <table className="TaskTable">
      <thead>
        <tr>
          <th>ID</th>
          <th>Title</th>
          <th>Content</th>
          <th>Status</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {tasks.map((task) => (
          <tr key={task.id}>
            <td>{task.id}</td>
            <td>{task.title}</td>
            <td>{task.content}</td>
            <td>
              <input
                type="checkbox"
                checked={task.completed}
                onChange={(e) => onCompleteChange(task.id, e.target.checked)}
              />
            </td>
            <td>
              <button onClick={() => onDelete(task.id)}>Delete</button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default TaskList;
