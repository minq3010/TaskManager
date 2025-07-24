import 'package:flutter/material.dart';
import '../models/task.dart';
import '../widgets/task_tile.dart';
import '../services/task_service.dart'; // giả sử bạn có file này

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});
  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  List<Task> tasks = [];

  @override
  void initState() {
    super.initState();
    TaskService.fetchTasks().then((value) {
      setState(() => tasks = value);
    });
  }

  void _addTask(String title) async {
    final newTask = await TaskService.createTask(title);
    setState(() => tasks.add(newTask));
  }

  void _toggleTask(Task task) async {
    task.completed = !task.completed;
    await TaskService.updateTask(task);
    setState(() {});
  }

  void _deleteTask(Task task) async {
    await TaskService.deleteTask(task.id);
    setState(() => tasks.remove(task));
  }

  void _openAddTaskModal() {
    showDialog(
      context: context,
      builder: (context) {
        String newTitle = "";
        return AlertDialog(
          title: const Text("New Task"),
          content: TextField(
            onChanged: (val) => newTitle = val,
            decoration: const InputDecoration(hintText: "Enter task title"),
          ),
          actions: [
            TextButton(
              onPressed: () {
                Navigator.of(context).pop();
              },
              child: const Text("Cancel"),
            ),
            ElevatedButton(
              onPressed: () {
                if (newTitle.trim().isNotEmpty) {
                  _addTask(newTitle.trim());
                  Navigator.of(context).pop();
                }
              },
              child: const Text("Add"),
            ),
          ],
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text("Task Manager")),
      body: tasks.isEmpty
          ? const Center(child: CircularProgressIndicator())
          : ListView(
              children: tasks.map((task) {
                return TaskTile(
                  task: task,
                  onToggle: () => _toggleTask(task),
                  onDelete: () => _deleteTask(task),
                );
              }).toList(),
            ),
      floatingActionButton: FloatingActionButton(
        onPressed: _openAddTaskModal,
        child: const Icon(Icons.add),
      ),
    );
  }
}
