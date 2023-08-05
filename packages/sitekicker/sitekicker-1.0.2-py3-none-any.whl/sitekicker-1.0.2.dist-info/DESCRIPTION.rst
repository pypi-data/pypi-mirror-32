Copyright (c) 2018 Minghao Ni

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

Description-Content-Type: text/markdown
Description: [![TravisCI Build Status](https://travis-ci.org/nmhnmh/sitekicker.svg)](https://travis-ci.org/nmhnmh/sitekicker)
        
        This is my **personal** static site generator, it lacks testings and documents at the moment.
        If you need a static site generator, find one [here](https://www.staticgen.com/) with good community support.
        
        **SiteKicker** is yet another static site builder written in Python3.
        
        Current Project Status: **Beta, usable, but no testing and documentation**
        
        ## Feature Todos
        
        - [ ] Flow chart, sequence chart support
        - [ ] Fix incremental build
        - [ ] Trigger page reload when page changed(use websocket)
        - [ ] post-processor to process <a> internal cross-link <a href="id:entry-id-here">A Cross Link</a>, report when the cross-link is broken
        - [ ] add pagination support
        - [ ] Generate ToC, add perm link to headers
        - [ ] More robust error handling and reporting
        
        # Documentation Todos
        
        - [ ] Add documentation
        - [ ] Setup home page with Github Pages
        
        ## Test Todos
        
        - [ ] Add Unit Test And Integration Tests for the Project
        - [ ] Add continuous test with [TravisCI](https://travis-ci.org/) for Mac and Linux
        - [ ] Add continuous test with [Appvoyor](https://www.appveyor.com/) for Windows
        
        ## sitekicker.yml
        ```yml
        # Name of the site
        name: An Awesome Site
        # Base URL for the site, will be used to generate absolute urls
        base_url: https://example.org
        # Directory where build output will be saved, could be relative path or absolute path
        output_dir: .dist
        # Directory that contains layout/templates, default: templates, optional, supported template format is jinja2
        template_dir: templates
        # Directories that will be copied, such as folders with assets or binary files
        copy_dirs:
          - assets
        ```
        
        ## folder.yml
        ```yml
        # The options set in this file will be applied to all entries inside the folder where this file is found,
        # we refer to these entries as 'affected items of this file' below.
        # This is a good place to set some common options and flags.
        # You could also add any custom options below, prefix and tags are special because they has special meaning
        
        # The prefix will be prepend to all items affected by this file, 'article.html' will be 'abc/article.html'
        # if multiple prefix specified along the way, they will be contatenated and prefixed to the final url,
        # so if two prefix 'a' and 'b' specified, then the final url will be '/a/b/article.html'
        prefix: abc
        # The tags listed here will be added to all items affected by this file, tags specified at different places
        # will be merged, duplicate tags will be removed, original order of tags will be maintained
        tags:
          - global tag 1
          - global tag 2
        ```
        
        ## Entry Front Matter
        ```yml
        # In this file your specified options for the entry, beside some predefined ones like 'id', 'title', 'date'
        # you can add your own custom options, and use it inside your templates, options specified here will override
        # options specified in meta.yml in parent folder, except 'prefix' and 'tags', the former will be concatenated,
        # the later will be merged
        
        # A unique id to identify the entry, no special chars, space will be substitued with hyphens, optional
        # when not set, will try to use file name as id, will emit an error when it is not possible
        id: some-thing-as-name
        # Title of the entry, mandatory, may contain any characters
        title: Sitekicker is another Static Site Generator
        # Date of the writting, mandatory, in the format of YYYY-MM-DD
        date: 2016-10-20
        # Date of update, optional
        update_date: 2016-11-20
        # Tags that applies to this entry, optional
        # current entry will inherit all tags in its parent folders,
        # if folder 'a' contains tag 'a', folder 'a/b' contains tag 'b'
        # entry 'a/b/entry.md' contains tag 'c', then enventually the entry will
        # have there tags: 'a', 'b', 'c'
        tags:
          - tag1
          - tag2
          - tag3
        ```
        
Platform: UNKNOWN
Classifier: Development Status :: 4 - Beta
Classifier: Programming Language :: Python
Classifier: Programming Language :: Python :: 3 :: Only
Classifier: Programming Language :: Python :: 3.5
Classifier: Programming Language :: Python :: 3.6
Classifier: Programming Language :: Python :: 3.7
Classifier: Environment :: Console
Classifier: Intended Audience :: Developers
Classifier: License :: OSI Approved :: MIT License
Classifier: Topic :: Software Development
Classifier: Topic :: System :: Networking
Classifier: Topic :: Terminals
Classifier: Topic :: Text Processing :: Markup :: HTML
Classifier: Topic :: Text Processing
Classifier: Topic :: Utilities
