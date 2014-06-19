replacer
========

Replacer is a simple console application that does two way text replacement based on configuration files.

#Usage
The application is run via the following command:
```
./replacer {options-file.json} {destination}
```

The options file contains the replacement rules for a given replacement operation.

```
{
  "Brand": "SampleBrand",
  "Token": "%%",
  "DefaultDirection": "reverse",
  "Filters": [
    { "Type": "Directory", "filterPattern" : "test" },
    { "Type": "Directory", "filterPattern" : "code" },
    { "Type": "File", "filterPattern" : "cs" },
    { "Type": "File", "filterPattern" : "txt" }
  ],
  "Values": [
    { "Source": "CompanyName", "value": "Mega Corp" },
    { "Source": "Copyright", "value": "Copyright 2012.  All Rights Reserved." }
  ]
}
```

**Brand**

The brand is used purely for display purposes at this time.

**Token**

The token represents the start and end blocks for a given replacement.  When the above configuration file is used, `%%CompanyName%%` gets replaced with `Mega Corp`.

**DefaultDirection**

There are two valid options:

1. forward
2. reverse

Forward does the replacement as you would expect.  `%%CompanyName%%` gets replaced with `Mega Corp`.  Reverse attempts to undo the replacement.  It searches for all instances of `Mega Corp` and inserts `%%CompanyName%%`.

Reverse direction can cause issues when the replacement value (Mega Corp in this case) isn't very distinct or multiple source rules have the same value.

**Filters**

There are two types of filters:

1. Directory
2. File

Directory filters are exclusion based.  `{ "Type": "Directory", "filterPattern" : "test" }` will exclude any directory named `test`.

File filters are inclusive.  If the filter pattern is set to `"filterPattern" : "cs"`, only files with a `.cs` extension will be replaced.  In the above example, all `.cs` and `.txt` files will have replacements made except for directories named `test` or `code`.

**Values**

The values are the actual replacement rules to be run.  The above example will replace the `%%CompanyName%%` and `%%Copyright%%` placeholders with their respective values.
