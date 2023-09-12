package code

var UsecaseTemplate string = `import 'package:dartz/dartz.dart';
import 'package:your_project/domain/%s/entity.dart';
import 'package:your_project/domain/%s/repository.dart';

class Get%s {
  final %sRepository repository;

  Get%s(this.repository);

  Future<Either<Failure, %sEntity>> call(String id) {
    return repository.get%s(id);
  }
}

// TODO: Implement other use cases (Get%ss, Add%s, Update%s, Delete%s)`
