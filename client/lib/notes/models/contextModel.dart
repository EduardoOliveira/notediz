import 'package:client/notes/domain/models/notes.dart' show Note;
import 'package:flutter/material.dart';

class ContextModel extends ChangeNotifier {
  ContextModel();

  List<Note> _notes = [];
  List<Note> get notes => _notes;

  void addNote(Note note) {
    _notes.add(note);
    notifyListeners();
  }

  void removeNote(Note note) {
    _notes.remove(note);
    notifyListeners();
  }
}
