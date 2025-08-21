import 'dart:convert';

class Bookmark {
  final String id;
  final String kind;
  final String url;
  final String title;
  final String description;
  final DateTime createdAt;
  final DateTime updatedAt;

  Bookmark({
    this.id = "",
    this.kind = "bookmark",
    this.url = "",
    this.title = "",
    this.description = "",
    DateTime? createdAt,
    DateTime? updatedAt,
  }) : createdAt = createdAt ?? DateTime.now(),
       updatedAt = updatedAt ?? DateTime.now();

  factory Bookmark.fromRawJson(String str) =>
      Bookmark.fromJson(json.decode(str));

  String toRawJson() => json.encode(toJson());

  factory Bookmark.fromJson(Map<String, dynamic> json) => Bookmark(
    id: json["ID"],
    kind: json["Kind"] ?? "bookmark",
    url: json["URL"],
    title: json["Title"],
    description: json["Description"],
    createdAt: DateTime.parse(json["CreatedAt"]),
    updatedAt: DateTime.parse(json["UpdatedAt"]),
  );

  Map<String, dynamic> toJson() => {
    "ID": id,
    "Kind": kind,
    "URL": url,
    "Title": title,
    "Description": description,
    "CreatedAt": createdAt.toUtc().toIso8601String(),
    "UpdatedAt": updatedAt.toUtc().toIso8601String(),
  };
}
