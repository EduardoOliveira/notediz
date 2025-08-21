import 'dart:convert';

import 'package:client/notes/domain/models/notes.dart' show Note;
import 'package:client/notes/domain/models/bookmarks.dart' show Bookmark;
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

class ContextModel extends ChangeNotifier {
  ContextModel();

  final List<dynamic> _notes = [];
  List<dynamic> get notes => _notes;
  List<Bookmark> get bookmarks => _notes.whereType<Bookmark>().toList();
  //List<Note> get notes => _notes.whereType<Note>().toList();

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
        Uri.parse('http://localhost:8080/api/bookmarks'),
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
