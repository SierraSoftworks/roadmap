# Schema
The Road Map file format is defined in a [JSONSchema](https://json-schema.org/) schema
which includes all of the supported fields and examples of their use. Many modern editors,
including [VS Code](https://code.visualstudio.com) include support for using JSONSchema files
to provide in-editor suggestions and validation.

You'll find the schema itself hosted on [schemas.sierrasoftworks.com](https://schemas.sierrasoftworks.com)
and you can directly reference it in your editor using the URL below.

```
https://schemas.sierrasoftworks.com/roadmap/v1/schema.json
```

## Using the Schema

### VS Code
If you have installed the [YAML extension](https://marketplace.visualstudio.com/items?itemName=redhat.vscode-yaml)
for Visual Studio code, you can add the `# yaml-language-server` directive to set the schema.

```yaml
# yaml-language-server: $schema=https://schemas.sierrasoftworks.com/roadmap/v1/schema.json
```

## roadmap.schema.json

<ClientOnly>
    <Schema url="https://raw.githubusercontent.com/SierraSoftworks/schemas/main/src/roadmap/v1/schema.json" />
</ClientOnly>
