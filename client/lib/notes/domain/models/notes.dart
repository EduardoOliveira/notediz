import 'dart:convert';

Function isAKind = (String kind) {
  switch (kind) {
    case 'note':
      return (Note note) => note is Note;
    case 'bookmark':
      return (Note note) => note is Bookmark;
    case 'text':
      return (Note note) => note is Text;
    default:
      return (Note note) => false;
  }
};

sealed class Note {
  String kind;
  String id;
  List<String> tags;
  DateTime createdAt;
  DateTime updatedAt;

  Note({
    required this.kind,
    this.id = '',
    this.tags = const [],
    required this.createdAt,
    required this.updatedAt,
  });
}

class Bookmark extends Note {
  final String url;
  final String title;
  final String description;

  Bookmark({
    super.id,
    super.tags,
    required this.url,
    this.title = '',
    this.description = '',
    DateTime? createdAt,
    DateTime? updatedAt,
  }) : super(
         kind: "bookmark",
         createdAt: createdAt ?? DateTime.now(),
         updatedAt: updatedAt ?? DateTime.now(),
       );

  factory Bookmark.fromRawJson(String str) =>
      Bookmark.fromJson(json.decode(str));

  String toRawJson() => json.encode(toJson());

  factory Bookmark.fromJson(Map<String, dynamic> json) {
    if (json["Kind"] != "bookmark") {
      throw Exception('Invalid note kind');
    }
    return Bookmark(
      id: json["ID"],
      tags: List<String>.from(json["Tags"] ?? []),
      url: json["URL"],
      title: json["Title"],
      description: json["Description"],
      createdAt: DateTime.parse(json["CreatedAt"]),
      updatedAt: DateTime.parse(json["UpdatedAt"]),
    );
  }

  Map<String, dynamic> toJson() => {
    "ID": id,
    "Kind": kind,
    "Tags": tags,
    "URL": url,
    "Title": title,
    "Description": description,
    "CreatedAt": createdAt.toUtc().toIso8601String(),
    "UpdatedAt": updatedAt.toUtc().toIso8601String(),
  };
}

class Text extends Note {
  final String content;

  Text({
    super.id,
    super.tags,
    required this.content,
    DateTime? createdAt,
    DateTime? updatedAt,
  }) : super(
         kind: "text",
         createdAt: createdAt ?? DateTime.now(),
         updatedAt: updatedAt ?? DateTime.now(),
       );

  factory Text.fromRawJson(String str) => Text.fromJson(json.decode(str));

  String toRawJson() => json.encode(toJson());

  factory Text.fromJson(Map<String, dynamic> json) {
    if (json["Kind"] != "text") {
      throw Exception('Invalid note kind');
    }
    return Text(
      id: json["ID"],
      tags: List<String>.from(json["Tags"] ?? []),
      content: json["Content"],
      createdAt: DateTime.parse(json["CreatedAt"]),
      updatedAt: DateTime.parse(json["UpdatedAt"]),
    );
  }

  Map<String, dynamic> toJson() => {
    "ID": id,
    "Kind": kind,
    "Tags": tags,
    "Content": content,
    "CreatedAt": createdAt.toUtc().toIso8601String(),
    "UpdatedAt": updatedAt.toUtc().toIso8601String(),
  };
}
