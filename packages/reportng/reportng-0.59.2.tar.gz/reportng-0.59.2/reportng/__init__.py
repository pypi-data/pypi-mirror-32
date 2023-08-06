# coding: utf-8
"""
Python report generator that wraps around bootstrap 4 using dominate.
Usage is simple. Follows header, section..., footer structure. reportng
relies on JS for some of its dynamic properties and has been developed
using modern browsers.
"""
import os
import logging
from requests import get
from collections import OrderedDict
import dominate
import dominate.tags as tag
from dominate.util import raw

# ugly way to address unicode encode issues
import sys
if sys.version[0] == '2':
    reload(sys)
    sys.setdefaultencoding('utf-8')
    import rnghelpers as rng
elif sys.version[0] == '3':
    from . import rnghelpers as rng

__author__ = 'securisec'
__version__ = '0.59.2'


class ReportWriter:
    """
    The main class that is used. Modifiable parameters are report_name,
    theme, brand and highlight_color

    Example with enabling code highlighting:
        >>> r = reporng.ReportWriter('my report', 'securisec', code=True)
    """

    def __init__(self, report_name, brand, asciinema=False, pbar=True,
                 search=True, code=False, theme_preview=False, **kwargs):
        """
        Assign theme and report name

        :param str brand: Name of the company/tester
        :param str report_name: Name of report. Default is Sample report
        :param bool theme_preview: Set to True to see a theme preview dropdown
        :param bool asciinema: Set to true to use asciinema's in report. Default is False
        :param bool code: Set to True in order to use code highlighting
        :param bool pbar: Set to True to enable the progress bar
        :param bool search: Set to False to disable search and highlight

        Example:
            >>> import reportng
            >>> r = reportng.ReportWriter(report_name='myreport', brand='securisec')
            >>> report = r.header(...)
        """
        self.report_name = report_name
        self.brand = brand
        self.document = dominate.document(title=self.report_name)
        self.asciinema = asciinema
        self.code = code
        self.pbar = pbar
        self.search = search
        self.theme_preview = theme_preview

    def report_header(self, theme='lux', highlight_color='#f1c40f', navbar_bg='primary',
                      custom_css='', **kwargs):
        """
        Controls the link and script tags in the head. This method should always be called
        at the top

        :param str theme: Name of any bootswatch theme. Default is lux
        :param str highlight_color: any rgb color. default is #f1c40f
        :param str script: warg Pass additional JS/jquery to add to header
        :param str custom_css: Pass additional custom CSS. This is in the header, and controls style for the whole report
        :param str navbar_bg: Controls the color of the navbar.
        :param bool bootstrap: Set to True to get default bootstrap theme. Does not work with local files
        :return: The head tag for the report.

        Example showing how to change the default theme:
            >>> r = report.report_header(theme='flatly')
        """

        if len(self.report_name) > 40:
            logging.warning('A report_name greater than 40 characters can \
            can cause the navbar to expand and break some functionality.')

        with self.document.head as report_head:
            # link and script builder for bootstrap 4
            tag.comment('Created using reportng by securisec')
            tag.meta(charset="utf-8", name="viewport",
                     content="width=device-width, initial-scale=1")
            # main style components
            tag.script(src=rng.JSCSS.jquery)
            tag.script(src=rng.JSCSS.popper_js)
            tag.script(src=rng.JSCSS.bs4_js)
            if not self.search == False:
                tag.script(src=rng.JSCSS.mark_js)

            # JS for tooltip
            tag.command('JS for tooltip')
            tag.script(raw(
                rng.JSCustom.tooltip_js
            ))

            # JS for mark_js
            tag.comment('JS for mark.js')
            tag.script(raw(
                rng.JSCustom.markjs_script
            ))

            # JS to populate the navbar dropdown
            tag.comment('JS to populate the navbar dropdown')
            tag.script(raw(
                rng.JSCustom.populate_navbar_onload
            ))

            # script that allows for smooth scrolling and adds padding for navbar
            tag.comment(
                'script that allows for smooth scrolling and adds padding for navbar')
            tag.script(raw(
                rng.JSCustom.smoothscroll_navbar_pad
            ))

            # js to filter in the dropdown menu
            tag.comment('js to filter in the dropdown menu')
            tag.script(raw(
                rng.JSCustom.dropdown_filter
            ))

            # user insterted JS
            if 'script' in kwargs:
                tag.comment('User inserted JS')
                tag.script(raw(
                    kwargs.get('script')
                ))

            # bootswatch style sheets
            tag.comment('style sheets')
            if 'bootstrap' in kwargs:
                if kwargs['bootstrap']:
                    bootswatch = 'https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/css/bootstrap.min.css'
            elif theme != 'lux':
                bootswatch = "https://bootswatch.com/4/%s/bootstrap.min.css" % theme
            else:
                bootswatch = rng.JSCSS.bootswatch
            tag.link(rel="stylesheet", type="text/css",
                     href=bootswatch, id='bootswatch')
            tag.link(href=rng.JSCSS.font_awesome, rel="stylesheet")

            # constructing this way to avoid loading un needed js and css
            # css for asciinema
            if self.asciinema == True:
                tag.comment('css for asciinema')
                tag.link(rel="stylesheet", type="text/css",
                         href=rng.JSCSS.asciinema_css)

            # css and js for highlight.js
            if self.code == True:
                tag.comment('css and js for highlight.js')
                tag.link(rel="stylesheet", href=rng.JSCSS.highlightjs_css)
                tag.script(src=rng.JSCSS.highlightjs_js)
                tag.script(raw(
                    """
                    hljs.initHighlightingOnLoad();
                    """
                ))

            # script for progress bar
            if self.pbar == True:
                tag.comment('js for progress bar')
                tag.script(
                    src=rng.JSCSS.progressbar_js)
                tag.script(raw(rng.JSCustom.progress_bar))

            # search highlight color control
            tag.comment('search highlight color control')
            # tag.style('span.highlight{background:  %s;}' %
            #           highlight_color)
            tag.style(raw(
                """
                mark {background: %s;}
                mark.current {background: orangered;}
                """ % highlight_color
            ))

            if custom_css != '':
                tag.comment('Custom CSS starts here')
                tag.style(raw(custom_css))

            # Navbar on top with 2 margin to seperate from first jumbotron. add class mb-2
            with tag.nav(
                    _class="navbar navbar-expand-lg navbar-dark bg-%s sticky-top" % rng.HelperFunctions.color_to_tag(
                            navbar_bg)):
                tag.a(self.brand, _class="navbar-brand", href="#")
                # sets the report title on the navbar
                tag.span(self.report_name, _class="navbar-text text-secondary")
                # theme previewer
                if self.theme_preview:
                    tag.comment('Theme previewer')
                    raw(rng.CustomHTML.themes_preview)
                # Button for responsive navbar
                with tag.button(_class="navbar-toggler",
                                type="button", data_toggle="collapse", data_target="#navbarid",
                                aria_controls="navbarColor01", aria_expanded="false", aria_label="Toggle navigation"):
                    tag.span(_class="navbar-toggler-icon")

                # Search box and button on navbar
                # make sure to include the word context to div/p tags to make it searchable
                with tag.div(_class="navbar-collapse collapse justify-content-md-end", id="navbarid"):
                    # ul class to house the navbar navigation items
                    with tag.ul(_class="navbar-nav"):
                        with tag.li(_class="nav-item"):
                            # add dropdown menu to house h1 tags from sections
                            with tag.div(_class="dropdown"):
                                tag.button('sections', _class="btn btn-secondary btn-block dropdown-toggle",
                                           type="button", id="dropdownMenuButton", data_toggle="dropdown",
                                           aria_haspopup="true", aria_expanded="false")
                                with tag.ul(_class="dropdown-menu dropdown-menu-right",
                                            aria_labelledby="dropdownMenuButton", id="ddmenu",
                                            style="max-height: 300px; height: auto; overflow: scroll"):
                                    # input box for filtering dropdown
                                    tag.input(_class="form-control-sm", id="ddfilter",
                                              type="text", placeholder="Filter..")
                        # highlight box form starts here
                        # input for search box
                        if not self.search == False:
                            tag.input(_class="form-control mr-sm-2", type="search", placeholder="Search",
                                      data_toggle="tooltip", data_placement="bottom",
                                      title="Regex capable. Case sensitive.")
                            # Show search hit count
                            tag.span("0", id="searchcount",
                                     style="color:%s; font-size: initial; padding-right: 8; align-self: center;" % highlight_color)
                            raw(
                                """
                                <button data-search="next" class="btn btn-sm btn-secondary">&darr;</button>
                                <button data-search="prev" class="btn btn-sm btn-secondary">&uarr;</button>
                                """
                                # <button data-search="clear" class="btn btn-sm btn-secondary">✖</button>
                            )
            if self.theme_preview:
                # theme preview jquery
                tag.comment('theme preview jquery')
                tag.script(raw(
                    rng.JSCustom.themes_preview
                ))
        return str(report_head)

    def report_section(self, title, content, html_content='', pre_tag=True, tag_color='default',
                       title_color=True, overflow=rng.CSSControl.css_overflow,
                       text_color='primary', h2_title=False, **kwargs):
        """
        This form the main body of the report

        :param str title: The h1/header title of the section
        :param bool pre_tag: Default is True and treats content as monospaced. Set to False to use p tag
        :param str content: The content for this section
        :param str tag_color: The severity color of the section.
        :param str text_color: Controls the color of the text. Use red, blue, green, yellow etc
        :param str overflow: Allows to control the style of the div container. Defaults to scroll on overflow. Set to empty string to have all content show
        :param bool h2_title: Set to True if the title should be converted to a h2 tag. Will not apprear in section and cannot modify colors
        :param bool title_color: Controls if the header background or text is colored. Default is True and lets background color.
        :param tuple alert: Kwarg Create a dismissable alert box. First value of tuple is the color, and the second is the message
        :param bool section: Kwarg Set to True to append the section to the preceding section. Default is false
        :param tuple reference: Kwarg Adds a small button which hrefs to a user supplied link. Is a tuple. First value is color, second is link
        :param dict badge: Kwarg Adds badges. Key is the color, and value is the message.
        :param str html_content: Insert raw html to be added to the end of the section
        :param dict modal: Create a modal. Is a dictionary. Valid keys are button, title and content

        :return: a jumbotron object
        :raises NotValidTag: Raises exception if a valid tag is not used

        Example show how to use a red title with only colored text
            >>> r += report.report_section('some title', content, tag_color='warning', titble_bg=False)
        """

        if title_color:
            color = 'bg'
        else:
            color = 'text'
        if tag_color not in rng.HelperFunctions.valid_tags:
            raise rng.NotValidTag(
                'Valid tags are primary secondary success danger warning info')

        # create a space between body jumbotrons
        tag.br()
        if 'section' in kwargs:
            style = rng.CSSControl.sticky_section_css
        else:
            style = rng.CSSControl.not_sticky_section
        # creates the jumbotron. User dictates if it is pre or p tag
        with tag.div(_class="jumbotron container context reportng-report-section-class",
                     style=style) as r:  # padding mods
            # can change the text color, or the background color
            if h2_title:
                tag.h2(title)
            else:
                t = tag.h1(title, _class="%s-%s" %
                                         (color, rng.HelperFunctions.color_to_tag(tag_color)),
                           id="%s" % rng.HelperFunctions.id_with_random(5, title))
            if 'reference' in kwargs:
                t.add(rng.HelperFunctions.ref_button(kwargs.get('reference')))
            # creates dismissable alert box
            if 'alert' in kwargs:
                rng.HelperFunctions.make_alert(kwargs.get('alert'))
            # creates a reference button with link
            with tag.div(_class="container", style=overflow):
                if pre_tag:
                    tag.pre(content, _class="text-%s" %
                                            rng.HelperFunctions.color_to_tag(text_color))
                else:
                    tag.p(content, _class="text-%s" %
                                          rng.HelperFunctions.color_to_tag(text_color))
            if 'badge' in kwargs:
                rng.HelperFunctions.create_badges(kwargs.get('badge'))
            if html_content != '':
                raw(html_content)
            if 'modal' in kwargs:
                if isinstance(kwargs.get('modal'), dict):
                    rng.HelperFunctions.make_modals(
                        title.replace(' ', ''), kwargs.get('modal'))
        return str(rng.HelperFunctions.convert_to_string(r))

    def report_section_collapsible(self, title, content='', header_color='default', raw_html='', pre_tag=True, **kwargs):
        """
        Is used to create a collapsible jumbotron container which is collapsed by default.

        :param str title: title of the section
        :param str content: content for the section, defaults to ''
        :param str header_color: color of the header, defaults to 'default'
        :param str raw_html: use raw html content instead of generated content, defaults to ''
        :param bool pre_tag: if set to false, uses a p tag, defaults to True
        :param dict kwargs: Applicable kwargs are alert and badge. Refer to report_section for format
        """

        color = 'bg-%s' % rng.HelperFunctions.color_to_tag(header_color)
        return rng.HelperFunctions.accordian_collapse(color, title,
                                                      content, pre_tag, raw_html, **kwargs)

    def report_image_carousel(self, *args, **kwargs):
        """
        Is used to create an image carousel with optional captions

        :param str \**args: A list of image paths
        :param str \**kwargs: Kwargs handle image captions and must be in the same order as args
        :return: image jumbotron carousel container
        :raises IndexError: If the number of kwargs is not equal to args

        Example:
            >>> r += report.report_image_carousel('example.jpg', 'example1.png',
            >>>                                 img1_caption='Beautiful', img2_caption='Wonderful')
        """
        # dict order for image captions
        if kwargs:
            ordered_kwargs = list(OrderedDict(kwargs).values())
        # create jumbotron container
        with tag.div(_class="jumbotron jumbomargin container reportng-image-carousel-class",
                     style=rng.CSSControl.sticky_section_css) as i:
            with tag.div(_class="carousel slide", id="carousel_controls", data_interval="false",
                         data_ride="carousel"):
                # add the carousel image indicator based on the number of images
                rng.HelperFunctions.slide_indicator(len(args))
                with tag.div(_class="carousel-inner"):
                    # iterate over *image_links
                    for index_num, image in enumerate(args):
                        # so that the first image is set to active
                        if index_num == 0:
                            with tag.div(_class="carousel-item active").add(
                                    tag.a(href=image, target="_blank")):
                                tag.img(
                                    src=image, _class="img-fluid img-thumbnail rounded mx-auto d-block")
                                if kwargs:
                                    tag.div(_class="carousel-caption").add(
                                        tag.p(ordered_kwargs[index_num]))
                        # images 2+
                        else:
                            with tag.div(_class="carousel-item").add(
                                    tag.a(href=image, target="_blank")):
                                tag.img(
                                    src=image, _class="img-fluid img-thumbnail rounded mx-auto d-block")
                                try:
                                    if kwargs:
                                        tag.div(_class="carousel-caption reportng-image-caption-class").add(
                                            tag.p(ordered_kwargs[index_num]))
                                except IndexError:
                                    logging.exception(
                                        'All captions needs to be set')
                    # carousel button
                    with tag.a(_class="carousel-control-prev", href="#carousel_controls",
                               role="button", data_slide="prev"):
                        tag.span(_class="carousel-control-prev-icon",
                                 aria_hidden="true")
                        tag.span("Previous", _class="sr-only")
                    with tag.a(_class="carousel-control-next", href="#carousel_controls",
                               role="button", data_slide="next"):
                        tag.span(_class="carousel-control-next-icon",
                                 aria_hidden="true")
                        tag.span("Next", _class="sr-only")
        return str(i)

    def report_asciinema(self, asciinema_link, title='', **kwargs):
        """
        Section creates a jumbotron to house an asciinema

        :param str asciinema_link: Link to asciinema. Could be http/s or local files
        :param str title: Set the title of the asciinema. If set, it will create its own section. If not, it will append to previous section
        :param tuple alert: Kwarg Create a dismissable alert box. First value of tuple is the color, and the second is the message
        :param bool section: Kwarg Set to True to append the cards section to the preceding section. Default is false
        :raises ObjectNotInitiated: Raises exception when the correct flags are not set in ReportWriter

        Example:
            >>> r += report.report_asciinema('https://asciinema.org/a/XvEb7StzQ3C1BAAlvn9CDvqLR', title='asciinema')
        """
        logging.warning('This method only works with asciinema links because of the way\n \
            browsers enforce CORS')
        # checks to see if asciinema has been intialized
        if not self.asciinema:
            raise rng.ObjectNotInitiated(
                'To integrate asciinema, set asciinema=True in ReportWriter')

        # TODO: write a check here that validates the asciinema url

        # hacky way to bypass the CORS problem
        try:
            url = get('%s.json' % asciinema_link).url
        except:
            logging.warning(
                'Need internet to get the proper url for %s' % asciinema_link)

        # controls if sticky or not
        if 'section' in kwargs:
            style = rng.CSSControl.sticky_section_css
        else:
            style = rng.CSSControl.not_sticky_section

        with tag.div(_class="jumbotron jumbomargin container reportng-acsiinema-class",
                     style=style) as a:
            if title != '':
                tag.h1(title, id="%s" %
                                 rng.HelperFunctions.id_with_random(5, title))
            # create dismissable alert box
            if 'alert' in kwargs:
                rng.HelperFunctions.make_alert(kwargs.get('alert'))
            with tag.div(_class="container", style="text-align: center;" + style):
                raw('<asciinema-player src="%s"></asciinema-player>' % url)
                tag.script(
                    src=rng.JSCSS.asciinema_js)
                tag.a('Asciinema link', _class="btn btn-secondary row justify-content-center btn-sm",
                      role="button", href=asciinema_link, target="_blank")
        return str(a)

    def report_code_section(self, title, code, **kwargs):
        """
        This section can use used to add code containers that will be lexed and highlighted using highlight.js

        :param str title: Title of the code section.
        :param str code: Code. Use pre and code tags so multiline code is fine
        :param bool section: Kwarg Set to True to append the cards section to the preceding section. Default is false
        :param tuple alert: Kwarg Create a dismissable alert box. First value of tuple is the color, and the second is the message
        :param tuple reference: Kwarg Adds a small button which hrefs to a user supplied link. Is a tuple. First value is color, second is link
        :param dict badge: Kwarg Adds badges. Key is the color, and value is the message.
        :param dict modal: Create a modal. Is a dictionary. Valid keys are button, title and content
        :return: a string code section
        :raises ObjectNotInitiated: Raises exception when the correct flags are not set in ReportWriter

        Example of how to get code from file:
            >>> with open('somefile.py', 'r') as f:
            >>>     data = f.read()
            >>> r += report.report_code_section('my py code', data)
        """
        if not self.code:
            raise rng.ObjectNotInitiated(
                'To integrate code highlighting, set code=True in ReportWriter')
        if 'section' in kwargs:
            style = rng.CSSControl.sticky_section_css
        else:
            style = rng.CSSControl.not_sticky_section
        with tag.div(_class="jumbotron container context reportng-code-section-class",
                     style=style) as c:  # padding mods
            t = tag.h1(title, id="%s" %
                                 rng.HelperFunctions.id_with_random(5, title))
            if 'reference' in kwargs:
                t.add(rng.HelperFunctions.ref_button(kwargs.get('reference')))
            # create dismissable alert box
            if 'alert' in kwargs:
                rng.HelperFunctions.make_alert(kwargs.get('alert'))
            with tag.div(_class="container", style="max-height: 70%; overflow: auto; margin-bottom: 20"):
                tag.pre().add(tag.code(code))
                if 'badge' in kwargs:
                    rng.HelperFunctions.create_badges(kwargs.get('badge'))
            if 'modal' in kwargs:
                if isinstance(kwargs.get('modal'), dict):
                    rng.HelperFunctions.make_modals(
                        title.replace(' ', ''), kwargs.get('modal'))
        return str(c)

    def report_captions(self, content, text_color='primary', **kwargs):
        """
        Simple method to added some center aligned text.

        :param str content: content to add
        :param str text_color: Controls the color of the text. Example, red, blue, yellow, green
        :param str custom_html: Insert raw html to be added to the end of the section

        Example:
            >>> r += report.report_captions('This is my caption')
        """
        if 'section' in kwargs:
            style = rng.CSSControl.sticky_section_css
        else:
            style = "margin-top:-30;"
        with tag.div(_class="container text-center reportng-captions-class", style=style) as s:
            tag.p(content, _class="text-%s" %
                                  rng.HelperFunctions.color_to_tag(text_color))
            if 'custom_html' in kwargs:
                raw(kwargs.get('custom_html'))
        return str(rng.HelperFunctions.convert_to_string(s))

    def report_table(self, *args, **kwargs):
        """
        Helps create tables

        :param tuple \*args: Rows of data as tuples. Ex ('data1', 'data2', ...)
        :param str \**kwargs: Kwargs to control title, header, etc.
        :param tuple header: Kwarg The header for the table. Accepts a tuple
        :param str title: Kwarg The title of the section
        :param bool section: Kwarg Set to false to append table to previous section
        :param str header_color: Kwarg Sets the color of the header. Defaults to dark
        :param bool tindex: Kwarg Sets an index column
        :param tuple alert: Kwarg Creats a dismissable alert box. Requires a tuple. First value is color and second value is message.
        :param dict modal: Create a modal. Is a dictionary. Valid keys are button, title and content

        :return: A table object

        :raises TypeError: Raises exception if not a tuple or all lenghts are not the same
        :raises NotValidTag: If not a valid bootstrap color
        :raises TableError: When there are issues creating a table

        Example showing how to change the default theme:
            >>> r = report.report_table(('data1', 'demo1'), ('data2', 'demo2'),
            >>>                         header=('header1', 'header2'), title='demo')
        """
        if 'section' in kwargs:
            style = rng.CSSControl.sticky_section_css
        else:
            style = "padding-bottom:3; padding-top:40;"
        if 'header_color' in kwargs:
            header_color = kwargs.get('header_color')
            if header_color not in rng.HelperFunctions.valid_tags:
                raise rng.NotValidTag('Not a valid header color')
        else:
            header_color = 'dark'
        # creates optional header
        if 'header' in kwargs:
            table_header = kwargs.get('header')
        # Check to make sure it is args
        if not isinstance(args, tuple):
            raise TypeError('Not a tuple')
        # Saves length of first arg
        check_length = len(args[0])
        # Run checks agains header
        if not isinstance(table_header, tuple):
            raise TypeError('Headers are not a tuple')
        elif len(table_header) != check_length:
            raise rng.TableError('Header not the same length as data')
        # starts building the table
        with tag.div(_class="jumbotron container context reportng-table-class", style=style) as c:  # padding mods
            if 'title' in kwargs:
                tag.h1(kwargs.get('title'), id="%s" %
                                               rng.HelperFunctions.id_with_random(5, kwargs.get('title')))
            # create dismissable alert box
            if 'alert' in kwargs:
                rng.HelperFunctions.make_alert(kwargs.get('alert'))
            with tag.div(_class="container", style="overflow-x:auto; max-height: 70%; overflow: auto;"):
                with tag.table(_class="table table-striped display nowrap table-hover", style="width: 90%") as tables:
                    # Make table header
                    if table_header:
                        with tag.thead(_class="table-%s" % rng.HelperFunctions.color_to_tag(header_color)).add(tag.tr()):
                            if kwargs.get('tindex') == True:
                                tag.th('Index')
                            for h in range(len(table_header)):
                                tag.th(table_header[h], scope='col')
                    for r in range(len(args)):
                        with tag.tr():
                            if kwargs.get('tindex') == True:
                                tag.td(str(r + 1))
                            # Checks length of subsequent args with first arg
                            if len(args[r]) != check_length:
                                raise rng.TableError(
                                    'Length of all tuples has to be the same')
                            for t in range(len(args[r])):
                                tag.td(args[r][t])
                if 'modal' in kwargs:
                    if isinstance(kwargs.get('modal'), dict):
                        rng.HelperFunctions.make_modals(
                            kwargs.get('title').replace(' ', ''), kwargs.get('modal'))
        return rng.HelperFunctions.convert_to_string(c)

    def report_cards(self, *args, **kwargs):
        """
        Functions that allow adding multiple cards to a jumbotron

        :param tuple \**args: Tuples that creates cards. The first value of the
            tuple is used to color the card, second value is the header for the
            card and the third is passed to a p tag for content
        * tuple values : Are in order fo ('color', 'title', 'content')
        :param bool section: Kwarg Set to True to append the cards section to the preceding section. Default is false
        :param str title: Kwarg Sets an optional title for cards
        :param bool border_only: Kwarg Changes the look of the cards
        :param tuple alert: Kwarg Creates a dismissable alert box. Requires a tuple. First value is color and second value is message.
        :raises TypeError: Raises TypeError if args is not a tuple
        :raises TooManyValues: Raises exception when the number of values in tuple is not 3

        Example:
            >>> r += report.report_cards(('primary', 'header1', 'some text'),
            >>>                       ('success', 'header2', 'some other text'),
            >>>                        title='My cards')
        """

        # Check to see if args is a tuple
        if not isinstance(args, tuple):
            raise TypeError('Use tuples only')

        # if the kwarg border exists, this set bool value
        if 'border_only' in kwargs:
            border = True
        else:
            border = False

        # control if stick to previous section or not
        if 'section' in kwargs:
            style = rng.CSSControl.sticky_section_css
        else:
            style = rng.CSSControl.not_sticky_section

        with tag.div(_class="jumbotron container context",
                     style=style) as c:  # padding mods
            if 'title' in kwargs:
                tag.h1(kwargs.get('title'))
            # create dismissable alert box
            if 'alert' in kwargs:
                rng.HelperFunctions.make_alert(kwargs.get('alert'))
            with tag.div(_class="row justify-content-center"):
                for i in range(len(args)):
                    # Checks to make sure corrent number of values in tuple
                    if len(args[i]) != 3:
                        raise rng.TooManyValues(
                            'Only pass three values to each tuple')
                    k = args[i][0]
                    h = args[i][1]
                    v = args[i][2]
                    rng.HelperFunctions.make_cards(border, k, h, v)
        return str(c)

    def report_footer(self, message='', **kwargs):
        """
        Returns the footer object. Supports social media

        :param str message: A message in the footer
        :param str twitter: Kwarg Twitter link
        :param str github: Kwarg Github link
        :param str linkedin: Kwarg LinkedIn link
        :param str email: Kwarg email address
        :param str custom_html: Insert raw html to be added to the end of the section

        Example with some social media:
            >>> r += report.report_footer(message='securisec', twitter='https://twitter.com/securisec')
        """
        # creates the footer
        with tag.footer(_class="page-footer reportng-footer-class") as footer:
            with tag.div(_class="container"):
                with tag.div(_class="row"):
                    with tag.div(_class="mb-4"):
                        # Looks through valid kwargs and creates appropiate a tag
                        for key, value in sorted(kwargs.items()):
                            if key == 'twitter':
                                tag.a(_class="icons-sm tw-ic", href=value, target="_blank").add(
                                    tag.i(_class="fab fa-twitter fa-2x white-text mr-md-4"))
                            elif key == 'github':
                                tag.a(_class="icons-sm gh-ic", href=value, target="_blank").add(
                                    tag.i(_class="fab fa-github fa-2x white-text mr-md-4"))
                            elif key == 'linkedin':
                                tag.a(_class="icons-sm li-ic", href=value, target="_blank").add(
                                    tag.i(_class="fab fa-linkedin fa-2x white-text mr-md-4"))
                            elif key == 'email':
                                tag.a(_class="icons-sm", href=value).add(
                                    tag.i(_class="fas fa-at fa-2x white-text mr-md-4"))
                        # i tag for user message
                        tag.span(message, style="font-size: 125%;")
                if 'custom_html' in kwargs:
                    raw(kwargs.get('custom_html'))

        return str(footer)

    def report_list_group(self, items, **kwargs):
        """
        Creates a list group. Each arg is a line

        :param str title: Title of the section
        :param list items: Each arg passed will be treated as a new list item
        :param bool section: Kwarg Set to True to append the cards section to the preceding section. Default is false
        :param tuple alert: Kwarg Create a dismissable alert box. First value of tuple is the color, and the second is the message
        :param tuple reference: Kwarg Adds a small button which hrefs to a user supplied link. Is a tuple. First value is color, second is link
        :param dict badge: Kwarg Adds badges. Key is the color, and value is the message.
        :param str custom_html: Insert raw html to be added to the end of the section
        :param dict modal: Create a modal. Is a dictionary. Valid keys are button, title and content
        :raises NoValidTag: Raises not valid tag

        Example:
            >>> r += report.report_list_group('First entry', 'Second entry', ...)
        """
        if not kwargs.get('title'):
            raise rng.NotValidTag('Need a title')
        title = kwargs.get('title')
        if not isinstance(items, list):
            raise rng.NotValidTag('Items have to be in the form of a list')

        if 'section' in kwargs:
            style = rng.CSSControl.sticky_section_css
        else:
            style = rng.CSSControl.css_overflow
        with tag.div(_class="jumbotron container context reportng-list-group-class", style=style) as r:
            t = tag.h1(title, id="%s" %
                                 rng.HelperFunctions.id_with_random(5, title))
            # creates a reference button with link
            if 'reference' in kwargs:
                t.add(rng.HelperFunctions.ref_button(kwargs.get('reference')))
            # creates dismissable alert box
            if 'alert' in kwargs:
                rng.HelperFunctions.make_alert(kwargs.get('alert'))

            with tag.ul(_class="list-group"):
                for i in range(len(items)):
                    tag.li(items[i],
                           _class="list-group-item d-flex justify-content-between align-items-center text-primary")

            if 'badge' in kwargs:
                rng.HelperFunctions.create_badges(kwargs.get('badge'))
            if 'custom_html' in kwargs:
                raw(kwargs.get('custom_html'))
            if 'modal' in kwargs:
                if isinstance(kwargs.get('modal'), dict):
                    rng.HelperFunctions.make_modals(
                        title.replace(' ', ''), kwargs.get('modal'))
        return rng.HelperFunctions.convert_to_string(r)

    def report_section_custom_html(self, html):
        """
        Add any custom html inside a jumbotron.

        :param str html: Any custom html
        """
        with tag.div(_class="jumbotron container context reportng-custom-html-class", style='padding:0') as c:
            raw(html)
        return rng.HelperFunctions.convert_to_string(c)

    def report_save(self, path, all_objects):
        """
        Saves the html file to disk

        :param str all_objects: The tally of all the objects
        :param str path: path to save the file

        Example:
            >>> report_save(report, '/tmp/demo.html')
        """
        with open(os.path.expanduser(path), 'w+') as save:
            save.write(str(all_objects))


class Assets:
    """
    Assets allows one to either download and map all dependent CSS and JS files, or
    use existing CSS and JS files
    """
    @staticmethod
    def local(rel_path):
        """
        This method allows one to map locally available asset files automatically.
        Themes are dicated by the locally available file

        :param str rel_path: The relative path from the report file. Usally is ./assets/

        Example:
            >>> from reportng import ReportWriter, Assets
            >>> Assets.local(rel_path='/tmp/assets/')
            >>> r = ReportWriter('Title', 'securisec')
        """
        change = vars(rng.JSCSS)
        for k, v in change.items():
            if not '__' in k:
                local_file = v.split('/')[-1]
                setattr(rng.JSCSS, k, rel_path + local_file)

    @staticmethod
    def download(download_path, rel_path, theme='lux'):
        """
        This method is used to download all online assests like JS/CSS locally. This method
        also will change all the src and href links to the local files

        :param str download_path: Path to save the files in
        :param str rel_path: Relative path from where the html will be saved
        :param str theme: The name of the bootswatch theme. Defaults to Lux

        Example:
            >>> from reportng import ReportWriter, Assets
            >>> Assets.download(save_path='/tmp/assets/', rel_path='./assets/')
            >>> r = ReportWriter('Title', 'securisec')
        """

        logging.warning(
            'Some files like font-awesome (all.css) does not work unless put into specific folders')
        # Check to make sure path is a dir
        if not os.path.isdir(download_path):
            raise TypeError('Path is not a directory')

        change = vars(rng.JSCSS)
        for k, v in change.items():
            if not '__' in k:
                local_file = v.split('/')[-1]
                if not os.path.exists(download_path + local_file):
                    with open(download_path + local_file, 'w+') as f:
                        if 'https://bootswatch.com/4/' in v:
                            v = v.replace('lux', theme)
                            local_file = v.split('/')[-1]
                        headers = {
                            'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36'
                        }
                        f.write(get(v, headers=headers).text)
                        logging.info('Downloaded %s to %s' %
                                     (v, download_path))
                        setattr(rng.JSCSS, k, rel_path + local_file)

# TODO: add a brand image that is resized with the navbar
# TODO: keep the image jumbotron static no matter the size of the picture
# TODO: make header method mandatory
# TODO: weird how the title of the code section is getting outside
