
# Notes Tool

## What Does the Tool Do?

The Notes Tool is a simple command-line utility that helps you manage a collection of notes. You can use it to create, view, edit, and delete individual notes within a specified file. The tool also allows for the deletion of an entire collection if needed, making it an easy and efficient way to manage your notes directly from the terminal.

## Tool Usage with Examples

### Basic Usage

To use the tool, run the following command:

```bash
./notestool [COLLECTION_NAME]
```

- **COLLECTION_NAME**: The name of the file where your notes will be stored. If the file does not exist, it will be created.

### Example Command

```bash
./notestool mynotes.txt
```

This command opens or creates the `mynotes.txt` file. Once opened, the tool will display a menu of options for managing your notes.

### Menu Options

1. **Show Notes**
   - **Description**: Displays all notes currently stored in the collection.
   - **Example**: Press `1` to view all notes in `mynotes.txt`.

2. **Add a Note**
   - **Description**: Prompts you to enter a new note, which is saved with a timestamp.
   - **Example**: Press `2`, enter your note, and it will be saved in the collection.

3. **Delete a Note**
   - **Description**: Removes a specific note from the collection based on its number.
   - **Example**: Press `3`, enter the number of the note to delete, and it will be removed.

4. **Edit a Note**
   - **Description**: Allows you to modify an existing note.
   - **Example**: Press `4`, select the note number, and edit the text.

5. **Delete the Collection**
   - **Description**: Permanently deletes the entire notes collection after confirmation.
   - **Example**: Press `5` to remove `mynotes.txt`. You'll be prompted to confirm before the file is deleted.

6. **Exit**
   - **Description**: Saves all changes and exits the program.
   - **Example**: Press `6` to save your notes and close the tool.

### Sample Workflow

Here’s a step-by-step example of how you might use the tool:

1. **Add a Note**: 
   - Press `2` to add a note. Enter "Buy groceries". The note is saved with a timestamp.

2. **Show Notes**: 
   - Press `1` to display the notes. You'll see something like:
     ```
     001 - Buy groceries [Last modified: 14:32:10 18.08.2024]
     ```

3. **Edit a Note**: 
   - Press `4` to edit a note. Enter `1` to select the first note, and change it to "Buy groceries and cook dinner".

4. **Delete a Note**: 
   - Press `3` to delete a note. Enter `1` to remove the first note.

5. **Exit**: 
   - Press `6` to save your changes and exit the program.

## Explanation of How the Data is Stored

The notes are stored in a plain text file specified by **COLLECTION_NAME**. Each note is saved on a new line with an associated timestamp indicating when it was last modified.

### Data Storage Details

- **File Structure**:
  - Each note is stored on a separate line within the file.
  - Example content of `mynotes.txt`:
    ```
    Buy groceries [Last modified: 14:32:10 18.08.2024]
    Finish the report [Last modified: 16:45:30 18.08.2024]
    ```

- **Adding Notes**:
  - New notes are appended to the end of the file, along with a timestamp.

- **Editing Notes**:
  - When a note is edited, the corresponding line in the file is updated with the new text and timestamp.

- **Deleting Notes**:
  - When a note is deleted, the corresponding line is removed from the file.

- **Deleting the Collection**:
  - If the collection is deleted, the entire file is removed from the system.

The tool automatically saves any changes you make to the notes, ensuring your data is preserved between sessions.
