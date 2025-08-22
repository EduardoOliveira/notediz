import 'package:client/notes/domain/models/notes.dart' as notes;
import 'package:client/notes/widgets/bookmarkCard.dart';
import 'package:client/notes/widgets/textCard.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'newNote.dart';
import '../models/contextModel.dart';
import 'package:flutter_staggered_grid_view/flutter_staggered_grid_view.dart';

class NotesList extends StatefulWidget {
  const NotesList({super.key});

  @override
  State<NotesList> createState() => _NotesListState();
}

class _NotesListState extends State<NotesList> {
  late Future<void> _loadNotesFuture;

  @override
  void initState() {
    super.initState();
    _loadNotesFuture = Provider.of<ContextModel>(
      context,
      listen: false,
    ).loadNotes();
  }

  @override
  Widget build(BuildContext context) {
    return Consumer<ContextModel>(
      builder: (context, contextModel, child) => FutureBuilder(
        future: _loadNotesFuture,
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return Center(child: CircularProgressIndicator());
          } else if (snapshot.hasError) {
            return Center(child: Text('Error: ${snapshot.error}'));
          }
          return MasonryGridView.count(
            crossAxisCount: 4,
            mainAxisSpacing: 4,
            crossAxisSpacing: 4,
            itemCount: contextModel.notes.length + 1,
            itemBuilder: (context, index) {
              if (index == 0) {
                // Show the new note card at the top
                return NewNoteCard();
              }
              if (index - 1 < contextModel.notes.length) {
                notes.Note n = contextModel.notes[index - 1];
                switch (n) {
                  case notes.Bookmark():
                    return BookmarkCard(bookmark: n);
                  case notes.Text():
                    return TextCard(text: n);
                  default:
                    print("Unknown note type: ${n.kind}");
                    return Placeholder();
                }
              }
              return Placeholder();
              //return Tile(index: index, extent: (index % 5 + 1) * 100);
            },
          );
        },
      ),
    );
  }
}
