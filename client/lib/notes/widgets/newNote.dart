import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../models/contextModel.dart';
import '../domain/models/notes.dart';

class NewNoteCard extends StatefulWidget {
  const NewNoteCard({super.key});

  @override
  State<NewNoteCard> createState() => _NewNoteCardState();
}

class _NewNoteCardState extends State<NewNoteCard> {
  String _noteContent = '';
  final String _noteKind = '';
  late final ContextModel contextModel;

  @override
  void initState() {
    super.initState();
    contextModel = Provider.of<ContextModel>(context, listen: false);
  }

  void _saveNote() {
    final note = Note(Content: _noteContent, Kind: _noteKind);
    contextModel.addNote(note);
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
