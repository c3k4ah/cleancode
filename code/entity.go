package code

var EntityTemplate string = `import 'package:equatable/equatable.dart';

abstract class %sEntity extends Equatable {
  final String id;
  final String title;
  final DateTime createdAt;
  final DateTime updatedAt;

  %sEntity({
    required this.id,
    required this.title,
    required this.createdAt,
    required this.updatedAt,
  });

  Map<String, dynamic> toMap();
  static %sEntity fromMap(Map<String, dynamic> map);
  String toJson();
  static %sEntity fromJson(String json);
  %sEntity copyWith({
    String? id,
    String? title,
    DateTime? createdAt,
    DateTime? updatedAt,
  });
}`
