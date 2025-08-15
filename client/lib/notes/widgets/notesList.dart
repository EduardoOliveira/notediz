import 'package:client/notes/widgets/noteCard.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'newNote.dart';
import '../models/contextModel.dart';
import 'package:flutter_staggered_grid_view/flutter_staggered_grid_view.dart';

class NotesList extends StatelessWidget {
  const NotesList({super.key});
  @override
  Widget build(BuildContext context) {
    return Consumer<ContextModel>(
      builder: (context, contextModel, child) => MasonryGridView.count(
        crossAxisCount: 4,
        mainAxisSpacing: 4,
        crossAxisSpacing: 4,
        itemBuilder: (context, index) {
          if (index == 0) {
            // Show the new note card at the top
            return NewNoteCard();
          }
          if (index - 1 < contextModel.notes.length) {
            final note = contextModel.notes[index - 1];
            return NoteCard(note: note);
          }
          return Placeholder();
          //return Tile(index: index, extent: (index % 5 + 1) * 100);
        },
      ),
    );
  }
}
