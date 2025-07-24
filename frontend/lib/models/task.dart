// lib/models/task.dart
class Task {
  final int id;
  String title;
  String description;
  bool completed;

  Task({required this.id, required this.title,required this.description, this.completed = false});

  factory Task.fromJson(Map<String, dynamic> json) {
    return Task(
      id: json['id'],
      title: json['title'] ?? " ",
      description: json['description'] ?? " ",
      completed: json['is_done'] ?? false,
    );
  }

  Map<String, dynamic> toJson() => {
    'id': id,
    'title': title,
    'completed': completed,
  };
}
