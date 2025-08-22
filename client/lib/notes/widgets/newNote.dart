import 'package:client/notes/domain/models/notes.dart' as notes;
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../models/contextModel.dart';

class NewNoteCard extends StatefulWidget {
  const NewNoteCard({super.key});

  @override
  State<NewNoteCard> createState() => _NewNoteCardState();
}

class _NewNoteCardState extends State<NewNoteCard> {
  String _noteContent = '';
  late final ContextModel contextModel;

  @override
  void initState() {
    super.initState();
    contextModel = Provider.of<ContextModel>(context, listen: false);
  }

  void _saveNote() {
    if (_noteContent.startsWith("http")) {
      final bookmark = notes.Bookmark(
        url: _noteContent,
        description: _noteContent,
        createdAt: DateTime.now(),
        updatedAt: DateTime.now(),
      );
      contextModel.addBookmark(bookmark);
      return;
    }
    contextModel.addNote(
      notes.Text(
        content: _noteContent,
        createdAt: DateTime.now(),
        updatedAt: DateTime.now(),
      ),
    );
  }

  void _noteChanged(String value) {
    setState(() {
      _noteContent = value;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          children: [
            TextField(
              minLines: 5,
              maxLines: null,
              onChanged: _noteChanged,
              decoration: InputDecoration(
                border: OutlineInputBorder(),
                labelText: 'Enter a note, link or task',
              ),
            ),
            Padding(
              padding: const EdgeInsets.only(top: 8.0),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  ElevatedButton(
                    onPressed: () {
                      // Handle save action
                      _saveNote();
                    },
                    child: Text('Save'),
                  ),
                  TextButton(
                    onPressed: () {
                      // Handle cancel action
                    },
                    child: Text('Cancel'),
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
