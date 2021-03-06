# go-akomantoso
AkomaNtoso go-libs for pardocs + dundocs data extraction.
For now, will mix up the usage; the input will be the extracted multiline content per question/answer.

NOTE: It is only using a subset of AKN; as per described in 

Should pass validation of AKN 3.0 schema; as per in schema.xsd; using example 
--> https://github.com/lestrrat-go/libxml2 (needs C; skip!)
--> https://github.com/terminalstatic/go-xsd-validate (needs C; skip!)
--> https://github.com/krolaw/xsd (needs C; skip!)

Would need something like --> https://xmlschema.readthedocs.io/en/latest/intro.html ; for native validation

Java Validators 
- https://github.com/amouat/xsd-validator
- https://github.com/docjason/XmlValidate

Sample given by SayIt does NOT comply with the latest AKN standard!

```bash
# Works --> https://github.com/amouat/xsd-validator
$ ./xsdv.sh ../go-akomantoso/schema.xsd ../go-akomantoso/testdata/the-tempest.an.xml
```
If just to extract into struct; this looks OK --> https://github.com/wagner-aos/go-xsd

Use https://github.com/shabbyrobe/cmdy to run this as a standalone on the spilit output from
go-pardocs + go-dundocs ..

## How to convert data to the standard we use – Akoma Ntoso

Akoma Ntoso1 is a comprehensive XML schema for several Parliamentary document types such as bills, acts, and debates. Various bodies around the world are starting to use or interoperate with Akoma Ntoso to model their data. Whilst it was designed for Parliamentary document types, the schema is general enough that it can be used for many different types of debate.

SayIt can import a subset of Akoma Ntoso, and below we describe which aspects of it we currently cover. You can export an Akoma Ntoso representation of any section on SayIt by adding .an to the end of any section URI, for example Shakespeare’s The Tempest: https://shakespeare.sayit.mysociety.org/the-tempest.an.

-  https://akomantoso.io/faq/why-is-akoma-ntoso-in-xml-and-not-html-or-json-or-pdf/

## Intro

- https://github.com/musale/akn-primer

## Standard (Full)

- http://docs.oasis-open.org/legaldocml/akn-core/v1.0/os/part1-vocabulary/akn-core-v1.0-os-part1-vocabulary.html#_Toc523925021
- 

## Schema Viewer

- http://schema.akomantoso.com/element/writtenStatements

## Examples

- https://leveson.sayit.mysociety.org/hearing-13-july-2012/mr-hugh-tomlinson-qc
- https://charles-taylor.sayit.mysociety.org/hearing-22-january-2013
- https://akomantoso.io/tag/hansard/


## Test data

In testdata folder ...