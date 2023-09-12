package code

var RepositoryTemplate string = `import 'package:dartz/dartz.dart';
import 'package:your_project/domain/%s/entity.dart';

abstract class %sRepository {
  Future<Either<Failure, %sEntity>> get%s(String id);
  Stream<Either<Failure, List<%sEntity>>> get%ss();
  Future<Either<Failure, void>> add%s(%sEntity %s);
  Future<Either<Failure, void>> update%s(%sEntity %s);
  Future<Either<Failure, void>> delete%s(String id);
}`
