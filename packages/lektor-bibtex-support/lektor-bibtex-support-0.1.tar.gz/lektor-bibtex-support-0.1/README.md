# Lektor Bibtex Plugin

The plugin allows adding a list of publications generated from bibtex files to a page.


## Enabling the plugin

To enable the plugin add this to your project file:

```
[packages]
lektor-bibtex-support = 0.1
```

## Configuring

Create a file configs/bibtex-support.ini with a section called Bibtex. Define a variable files that
is a white space separated list of bibtex files. Put these files into the assets directory of your project.

```
[Bibtex]
files = A.bib B.bib
```

Optionally you can also create a template file that should leave in the template directory to render the bibtex entries. If you skip this entry a default template will be used that sorts the entries by year.

## Usage

You can add the the publication anywhere in your jinja template by calling

```
{{ list_publications(name=name, tag=tag, year=year, labels=labels, fname=fname)|safe }}
```

The arguments are optional and can be used for filtering.

By name: the name has to show up in the author list to be included

by year: only publications from this year

labels: white space separated list of bibtex labels (the name of each entry in the bibtex file)

fname: white space separated list files to search (by default all files will be searched)


## Javascript

For the default template, you can also add the following javascript to show some entries if you have a long list (relies on jquery):

```
$(document).ready(function()
	   {
	      $(".BIBTeX").hide();
	      $(".BIBTeXtoggle").click( function () {
		  $(this).parent().children(".BIBTeX").toggle(300);
		  return false;
		});
	      $(".BIBYear").hide();
	      $(".BIBYear:first").show();
	      $(".BIBYearheader").click( function () {
		  $(this).parent().children(".BIBYear").toggle(300);
		  return false;
		});
	      $(".bibshowall").click( function () {
		  $(".BIBYear").show();
		});
	      $(".bibhideall").click( function () {
		  $(".BIBYear").hide();
		  $(".BIBYear:first").show();
		});
           });
```
