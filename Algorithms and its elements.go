Algorithme nombre de caractère
Variables:
    c : character
    length := 0
    numWords := 0
    numVowels := 0

Begin

    Read(c)

    While c ≠ '.' do

        length := length + 1

        // Count vowels
        If c = 'a' or c = 'e' or c = 'i' or
           c = 'o' or c = 'u' or c = 'y' or
           c = 'A' or c = 'E' or c = 'I' or
           c = 'O' or c = 'U' or c = 'Y'
        Then
            numVowels := numVowels + 1
        EndIf

        // Count words
        If length = 1 Then
            numWords := 1
        EndIf

        If c = ' ' Then
            numWords := numWords + 1
        EndIf

        Read(c)

    EndWhile

    // Count the period
    length := length + 1

    Print(length)
    Print(numWords)
    Print(numVowels)

End
