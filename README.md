# terraform-provider-bitbucket
temp location for TrueMark Bitbucket Client


## Generation of BitBucket Client API: 

git clone https://github.com/OpenAPITools/openapi-generator-cli

git clone https://github.com/OpenAPITools/openapi-generator
 
```
    npm install -g @openapitools/openapi-generator-cli
    yarn add global @openapitools/openapi-generator-cli
``` 

You can verify that openapi-generator-cli is installed by running:
```
    $ npx openapi-generator-cli

    Usage: openapi-generator-cli <command> [<args>]
    Commands:
        version-manager  Manage used / installed generator version
        author           Utilities for authoring generators or customizing templates.
        batch            Generate code in batch via external configs.
        config-help      Config help for chosen lang
        generate         Generate code with the specified generator.
        help             Display help information about openapi-generator
        list             Lists the available generators
        meta             MetaGenerator. Generator for creating a new template set 
                         and configuration for Codegen.  The output will be based 
                         on the language you specify, and includes default templates to include.
        validate         Validate specification
        version          Show version information used in tooling
```

CONFIG OPTIONS
	disallowAdditionalPropertiesIfNotPresent
	    If false, the 'additionalProperties' implementation (set to true by default) is compliant with the OAS and JSON schema specifications. 
        If true (default), keep the old (incorrect) behaviour that 'additionalProperties' is set to false by default. (Default: true)
	        false - The 'additionalProperties' implementation is compliant with the OAS and JSON schema specifications.
	        true - Keep the old (incorrect) behaviour that 'additionalProperties' is set to false by default.

	enumClassPrefix
	    Prefix enum with class name (Default: false)

	generateInterfaces
	    Generate interfaces for api classes (Default: false)

	hideGenerationTimestamp
	    Hides the generation timestamp when files are generated. (Default: true)

	isGoSubmodule
	    whether the generated Go module is a submodule (Default: false)

	packageName
	    Go package name (convention: lowercase). (Default: openapi)

	packageVersion
	    Go package version. (Default: 1.0.0)

	prependFormOrBodyParameters
	    Add form or body parameters to the beginning of the parameter list. (Default: false)

	structPrefix
	    whether to prefix struct with the class name. e.g. DeletePetOpts => PetApiDeletePetOpts (Default: false)

	useOneOfDiscriminatorLookup
	    Use the discriminator's mapping in oneOf to speed up the model lookup. IMPORTANT: Validation (e.g. one and only one match in oneOf's schemas) will be skipped. (Default: false)

	withAWSV4Signature
	    whether to include AWS v4 signature support (Default: false)

	withXml
	    whether to include support for application/xml content type and include XML annotations in the model (works with libraries that provide support for JSON and XML) (Default: false)



```
    openapi-generator-cli generate -i https://api.bitbucket.org/swagger.json -g go --skip-validate-spec
```
