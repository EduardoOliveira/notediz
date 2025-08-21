import 'package:client/notes/domain/models/bookmarks.dart';
import 'package:flutter/material.dart';

class TextCard extends StatelessWidget {
  const TextCard({super.key, required this.bookmark});

  final Bookmark bookmark;

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(bookmark.url, style: Theme.of(context).textTheme.titleLarge),
            const SizedBox(height: 8.0),
            Text(bookmark.kind, style: Theme.of(context).textTheme.titleMedium),
          ],
        ),
      ),
    );
  }
}
