// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'notes.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_Note _$NoteFromJson(Map<String, dynamic> json) => _Note(
  Content: json['Content'] as String? ?? '',
  Kind: json['Kind'] as String? ?? '',
);

Map<String, dynamic> _$NoteToJson(_Note instance) => <String, dynamic>{
  'Content': instance.Content,
  'Kind': instance.Kind,
};
