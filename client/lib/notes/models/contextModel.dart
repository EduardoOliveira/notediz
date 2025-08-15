import 'dart:convert';

import 'package:client/notes/domain/models/notes.dart' show Note;
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

class ContextModel extends ChangeNotifier {
  ContextModel();

  final List<Note> _notes = [];
  List<Note> get notes => _notes;

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
}
