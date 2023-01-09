# Obscene-Vocabulary-Checker
Check sentences and replace the words that you find in the dictionary with ****.
The input includes lines with a sentence where the first word starts with a
capital letter. All words are separated by one space, and the sentence ends with
a period. Your task is to replace the obscene words with ****. Good luck!

## Objectives

   - Read the name of the file that contains taboo words from the user input;
   - Read the file;
   - Read sentences from the input. If the sentence contains taboo words, replace
   them with *, according to their length, and print the result. Otherwise, print
   the sentence. Repeat this until the exit command;
   - When the program receives the exit command, print Bye! and exit the program.

### Example

The greater-than symbol followed by a space (> ) represents the user input. Note that it's not part of the input.

Example 1:

The forbidden_words.txt contents:
disgusting
unpleasant
ugly
bad

```
> forbidden_words.txt
> It is a Bad and ugly sentence.
It is a *** and **** sentence.
> It is a Badge.
It is a Badge.
> exit
Bye!
```
