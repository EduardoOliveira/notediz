import 'package:client/notes/domain/models/notes.dart';
import 'package:flutter/material.dart';

class NoteCard extends StatelessWidget {
  const NoteCard({super.key, required this.note});

  final Note note;

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(note.Content, style: Theme.of(context).textTheme.titleLarge),
            const SizedBox(height: 8.0),
            Text(note.Kind, style: Theme.of(context).textTheme.titleMedium),
          ],
        ),
      ),
    );
  }
}
