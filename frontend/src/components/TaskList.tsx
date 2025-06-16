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
              <select
                value={task.status}
                onChange={(e) => onStatusChange(task.id, e.target.value)}
              >
                <option value="registered">未着手</option>
                <option value="doing">進行中</option>
                <option value="completed">完了</option>
              </select>
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
