package types

/*
	Description for Root:
		If we have several files: 5_1625481283.png and 5_1625617232.png, the Root for these two files is going to be '5'.
		So passing the root we assume everything before the underscore '_'.
*/
type Root string
