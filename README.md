# simplej2

Simplej2 is very simple tool to process existing jinja2 templates while using
a JSON file for the necessary variables used in the template.

## Usage

The main purpose of `simplej2` is to process a folder containing Jinja2 templates
and save the generated files in a target directory, where the same directory
structure is generated as in the source folder. If the source is a directory,
only files ending on `.j2` are processed. For the resulting files, this suffix
is stripped.

If the source is a single file, the file ending does not matter.

```shell
simplej2 <value.json> <src> <destDir>
```

The `value.json` file does have the following structure:

```json
{
  "variable1": true,
  "variable2": "foobar"
}

```

Given the following template,

```none
{% if variable1 %}
{{ variable2 }}
{% endif %}
```

a resulting file with the following content is produced.

```none
foobar
```

The rendering is done with [gonja](https://github.com/noirbizarre/gonja), which
is originally based on [pongo2](https://github.com/flosch/pongo2), but does
support Jinja2 template syntax instead of Django template syntax (for differences
see <https://jinja.palletsprojects.com/en/2.10.x/switching/#django>).
