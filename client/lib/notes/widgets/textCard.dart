import 'package:client/notes/domain/models/notes.dart' as notes;
import 'package:flutter/material.dart';

class TextCard extends StatelessWidget {
  const TextCard({super.key, required this.text});

  final notes.Text text;

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(text.content, style: Theme.of(context).textTheme.titleLarge),
            const SizedBox(height: 8.0),
            Text(text.kind, style: Theme.of(context).textTheme.titleMedium),
          ],
        ),
      ),
    );
  }
}
