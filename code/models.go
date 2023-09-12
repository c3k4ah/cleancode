package code

var ModelTemplate = `import 'package:your_project/domain/%s/entity.dart';
	
class %sModel extends %sEntity {
  %sModel({
    required String id,
    required String title,
    required DateTime createdAt,
    required DateTime updatedAt,
  }) : super(
          id: id,
          title: title,
          createdAt: createdAt,
          updatedAt: updatedAt,
        );

  @override
  Map<String, dynamic> toMap() {
    // TODO: Implement toMap
  }

  @override
  static %sEntity fromMap(Map<String, dynamic> map) {
    // TODO: Implement fromMap
  }

  @override
  String toJson() {
    // TODO: Implement toJson
  }

  @override
  static %sEntity fromJson(String json) {
    // TODO: Implement fromJson
  }

  @override
  %sEntity copyWith({
    String? id,
    String? title,
    DateTime? createdAt,
    DateTime? updatedAt,
  }) {
    // TODO: Implement copyWith
  }
}`
