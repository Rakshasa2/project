import React from 'react';
import { View, Text, StyleSheet, TouchableOpacity } from 'react-native';

const TaskItem = ({ task }) => (
  <View style={styles.taskContainer}>
    <Text style={styles.taskText}>{task}</Text>
  </View>
);

const TaskList = () => {
  const tasks = [
    "Принять витамины до 10:00",
    "Завтрак до 10:00",
    "Встреча в баре до 16:00",
    "Забрать сына из школы до 17:30",
    "Купить подарок жене до 18:30"
  ];

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Доброе утро!</Text>
      <Text style={styles.subtitle}>Ваши задачи на сегодня:</Text>
      {tasks.map((task, index) => (
        <TaskItem key={index} task={task} />
      ))}
      <TouchableOpacity style={styles.addButton}>
        <Text style={styles.addButtonText}>Добавить задачу</Text>
      </TouchableOpacity>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#F5F5F5',
    padding: 20
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    marginBottom: 10
  },
  subtitle: {
    fontSize: 16,
    color: '#555',
    marginBottom: 20
  },
  taskContainer: {
    backgroundColor: '#FFF',
    padding: 15,
    borderRadius: 10,
    marginBottom: 10,
    elevation: 1 
  },
  taskText: {
    fontSize: 16,
  },
  addButton: {
    backgroundColor: '#000',
    padding: 15,
    borderRadius: 10,
    alignItems: 'center',
    marginTop: 20
  },
  addButtonText: {
    color: '#FFF',
    fontSize: 16,
  },
});
export default TaskList;
