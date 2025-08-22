import 'dart:convert';

import 'package:client/notes/domain/models/notes.dart';
import 'package:flutter/material.dart' hide Text;
import 'package:http/http.dart' as http;

class ContextModel extends ChangeNotifier {
  ContextModel();

  final List<Note> _notes = [];
  List<Note> get notes => _notes;
  List<Bookmark> get bookmarks => _notes.whereType<Bookmark>().toList();
  List<Text> get texts => _notes.whereType<Text>().toList();

  Future<void> loadNotes() async {
    final response = await http.get(
      Uri.parse('http://localhost:8080/api/notes'),
    );
    if (response.statusCode == 200) {
      final List<dynamic> data = jsonDecode(response.body);
      _notes.clear();
      _notes.addAll(
        data.map((item) {
          switch (item['Kind']) {
            case 'bookmark':
              return Bookmark.fromJson(item);
            case 'text':
              return Text.fromJson(item);
            default:
              throw Exception('Unknown note kind: ${item['Kind']}');
          }
        }),
      );
      notifyListeners();
    } else {
      throw Exception('Failed to load notes');
    }
  }

  Future<void> addNote(Note note) async {
    _notes.add(note);
    notifyListeners();
    try {
      final response = await http.post(
        Uri.parse('http://localhost:8080/api/notes'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode(note),
      );
      if (response.statusCode == 201) {
        // Note added successfully
      } else {
        throw Exception('Failed to add note');
      }
    } catch (e) {
      _notes.remove(note);
      notifyListeners();
      rethrow; // Re-throw the exception to handle it in the UI
    }
  }

  Future<void> addBookmark(Bookmark bookmark) async {
    print("Adding bookmark: ${bookmark.toJson()}");
    _notes.add(bookmark);
    notifyListeners();
    try {
      final response = await http.post(
        Uri.parse('http://localhost:8080/api/notes'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode(bookmark),
      );
      if (response.statusCode == 201) {
        // Bookmark added successfully
      } else {
        throw Exception('Failed to add bookmark');
      }
    } catch (e) {
      _notes.remove(bookmark);
      notifyListeners();
      rethrow; // Re-throw the exception to handle it in the UI
    }
  }
}
