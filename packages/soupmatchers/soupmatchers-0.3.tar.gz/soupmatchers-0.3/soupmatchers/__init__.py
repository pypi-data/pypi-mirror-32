"""Matchers for testing HTML.

Given a string containing some HTML you can assert that it contains
a particular tag with

    self.assertThat(the_html, HTMLContains(Tag('a link', 'a')))

You can also check particular attributes:

    Tag('a link', 'a', attrs={'class': 'foo'})

or that it matches a certain number of times:

    Tag('a link', 'a', count=2)

In addition you can assert that one part of the document exists within
another.

   the_body = Tag('the body', 'body', count=1)
   one_link = Tag('one link', 'a', count=1)
   self.assertThat(the_html, HTMLContains(one_link.within(the_body)))

See the README for more information.
"""

import bs4

from testtools import matchers
from testtools.content import Content
from testtools.content_type import ContentType


class StatusCodeMismatch(matchers.Mismatch):
    """The status code of a response differs from the expected."""

    def __init__(self, expected_code, actual_code, html=None):
        self.expected_code = expected_code
        self.actual_code = actual_code
        self.html = html

    def describe(self):
        return "Got response code %s, expected %s" % (str(self.actual_code),
                str(self.expected_code))

    def get_details(self):
        # TODO: don't assume html
        if self.html is not None:
            return {
                "html": Content(ContentType("text", "html"),
                        lambda: [self.html.encode("UTF-8")])
            }
        return {}

    def __eq__(self, other):
        return (self.expected_code == other.expected_code
                and self.actual_code == other.actual_code
                and self.html == other.html)

    def __ne__(self, other):
        return (self.expected_code != other.expected_code
                or self.actual_code != other.actual_code
                or self.html != other.html)


class ResponseHas(matchers.Matcher):
    """Assert that the response has the desired characteristics."""

    def __init__(self, status_code=200, content_matches=None):
        """Create a ResponseHas matcher.

        :param status_code: the value that the status_code attribute
            of the response should have.
        :type status_code: int
        :param content_matches: a matcher that should be checked against
            the content attribute of the response.
        :type content_matches: matchers.Matcher
        """
        self.status_code = status_code
        self.content_matches = content_matches

    def match(self, response):
        if int(response.status_code) != int(self.status_code):
            return StatusCodeMismatch(
                self.status_code, response.status_code,
                html=response.content)
        if self.content_matches is not None:
            return self.content_matches.match(response.content)
        return None

    def __str__(self):
        return "%s(status_code=%s, content_matches=%s)" % (
                str(self.__class__.__name__), str(self.status_code),
                str(self.content_matches))


class HTMLResponseHas(ResponseHas):
    """Assert that the HTML response has the desired characteristics."""

    # TODO: check that the response content type is HTML

    def __init__(self, status_code=200, html_matches=None):
        """Create a ResponseHas matcher.

        :param status_code: the value that the status_code attribute
            of the response should have.
        :type status_code: int
        :param html_matches: a matcher that should be checked against
            the parsed HTML version of the content attribute of the
            response.
        :type content_matches: matchers.Matcher
        """
        self.status_code = status_code
        self.content_matches = html_matches and HTMLContains(html_matches)


class HTMLContains(matchers.Matcher):
    """Assert that the HTML has certain contents.

    Pass a list of other matchers to the constructor to be matched
    against the parsed content.
    """

    def __init__(self, *args):
        self.matchers = args

    def match(self, content):
        if len(self.matchers) > 0:
            parsed_content = bs4.BeautifulSoup(content, "html.parser")
            for matcher in self.matchers:
                match = matcher.match(parsed_content)
                if match is not None:
                    return match
        return None

    def __str__(self):
        return "HTML contains [%s]" % (
            ", ".join([str(a) for a in self.matchers]),)


class TagMismatch(matchers.Mismatch):
    """TagMismatch indicates that a Tag matched the wrong number of times.

    Tags are searched for and the number that match the specification is
    checked against the expected number, and if the two differ this
    Mismatch is returned.
    """

    def __init__(self, tag, content, matches):
        self.tag = tag
        self.content = content
        self.matches = matches

    def describe(self):
        if len(self.matches) != 1:
            s = "s"
        else:
            s = ""
        base = "Matched %s time%s" % (len(self.matches), s)
        if len(self.matches) > 1:
            base += "\nThe matches were:"
            for match in self.matches:
                base += "\n   %s" % (match,)
        elif len(self.matches) > 0:
            base += "\nThe match was:\n   %s" % (self.matches[0],)
        extra_info = self.tag.get_extra_info([self.content], "")
        if len(extra_info) > 0:
            base += "\nHere is some information that may be useful:\n  "
            base += "\n  ".join(extra_info)
        return base

    def get_details(self):
        details = {
            "html": Content(
                ContentType("text", "html"),
                lambda: [str(self.content).encode("UTF-8")])
        }
        return details

    def __eq__(self, other):
        return (self.tag == other.tag
                and self.content == other.content
                and self.matches == other.matches)

    def __ne__(self, other):
        return (self.tag != other.tag
                or self.content != other.content
                or self.matches != other.matches)

    def __repr__(self):
        return  "<soupmatchers.TagMismatch object at %x attributes=%r>" % (
            id(self), self.__dict__)


class DocumentPart(matchers.Matcher):

    def __init__(self, count=None):
        self.count = count

    def __str__(self):
        raise NotImplementedError(self.__str__)

    @property
    def identifier(self):
        raise NotImplementedError(self.identifier)

    def get_matches(self, matchee):
        raise NotImplementedError(self.get_matches)

    def match(self, matchee):
        matches = self.get_matches(matchee)
        if ((self.count is not None and len(matches) != self.count)
                or (self.count is None and len(matches) < 1)):
            return TagMismatch(self, matchee, [str(m) for m in matches])
        return None

    def within(self, tag):
        """Returns a Within that will match if this is within `tag`.

        :param tag: the tag that the object must be within to match.
        :type tag: Tag
        :return: a matcher that will match if this is within the tag.
        :rtype: Within
        """
        if not isinstance(tag, Tag):
            raise ValueError("%s is not a tag" % tag)
        return Within(tag, self)

    def get_extra_info(self, html_list, identifier_suffix):
        """Get extra info that may be relevant to the user.

        Passed a list of interesting sub-trees, provide any information
        that would be useful to the user about this Matcher in those
        subtrees.

        :param html_list: the list of sub-trees.
        :type html_list: a list of BeautifulSoup objects.
        :param identifier_suffix: a suffix for the identifier for use
            when identifying this Matcher.
        :param identifier_suffix: str
        :return: a list of strings containing useful information.
        :rtype: list(str)
        """
        raise NotImplementedError(self.get_extra_info)


class _NotPassed(object):

    def __str__(self):
        return "Not passed"


_not_passed = _NotPassed()


class Tag(DocumentPart):
    """Assert that a particular tag occurs a certain number of times."""

    def __init__(self, identifier, tag_type, attrs=None, text=_not_passed,
            count=None):
        """Create a Tag matcher.

        identifier is a short string that makes sense to the user of
        the matcher. It can be anything, but should be unique to the
        particular options passed, and should be informative about
        what should be passed.

        tag_type selects the tag name(s) that will be matched. It can
        either be a str to match a single tag name, a regular expression to
        match any tag name that matches the regular expression, a list
        to match any tag name in the list, a dict to match any tag that
        is a key in the dict, or True, which matches any tag name.

        It is also possible to pass a callable as a tag name. The callable
        should take a bs4.Tag object as the argument and return a boolean,
        with True indicating that the tag should match.

        attrs is a dict that defines what attributes the tag should have.
        The keys are the names of attributes and the values define the
        values of that attribute that the key must have to match. The values
        of the dict can be strings to match that string, lists to match
        any of the values of the list, a dict to match any of the keys
        of the dict, True to match any value of the attribute, None
        to match only those tags that have no value for the attribute,
        or a callable which will be called with the attributes of the
        tag, or None if there are none for that tag, and should return
        True if the tag should match.

        text allows you to match the contents of a tag, rather than the
        tag itself. This is useful for e.g. matching the text that will
        be part of the anchor in a link. You can specify a string to
        match tags containing that string exactly, a regular expression
        to match tags containing strings that match the expression,
        a list to match any string in the list, or a dictionary to match
        any of the keys, True to match any tag with text contents, or None
        to match any tag without text contents, or a callable that will be
        called with the text contained by a tag and should return True
        to indicate that the tag should be matched.

        For instance given

           <a></a><a><b></b></a>

        a

            Tag("a", count=2)

        would match.

        :param tag_type: the name of the tag, e.g. "a", "image", ...
        :type tag_type: str, regex, list, dict, True or a callable
        :param attrs: attributes the tag must have.
        :type attrs: dict
        :param text: text contents that the tag must have
        :type text: str, regex, list, dict, True, None, or a callable
        :param count: the number of times tags matching the arguments must
            appear, default None, to match any number of times.
        :type count: int or None
        """
        super(Tag, self).__init__(count=count)
        self._identifier = identifier
        self.tag_type = tag_type
        self.attrs = attrs or {}
        self.text = text

    @property
    def identifier(self):
        return self._identifier

    def _check_text(self, candidates, text):
        if len(candidates) > 0 and text is not _not_passed:
            for candidate in candidates[:]:
                texts = candidate.find_all(text=text)
                if len(texts) < 1:
                    candidates.remove(candidate)

    def _get_matches(self, html, attrs, text):
        candidates = list(html.find_all(self.tag_type, attrs=attrs))
        self._check_text(candidates, text)
        return candidates

    def get_matches(self, html):
        return self._get_matches(html, self.attrs, self.text)

    def __str__(self):
        plural = ""
        if self.count is not None and self.count != 1:
            plural = "s"
        attrs = "..."
        if len(self.attrs) > 0:
            attrs = " ".join("%s='%s'" % a for a in sorted(self.attrs.items()))
            attrs += " ..."
        children = ""
        text = "..."
        if self.text is not _not_passed:
            import re
            if isinstance(self.text, type(re.compile(""))):
                text = "re.compile('%s') ..." % self.text.pattern
            else:
                text = "%s ..." % self.text
        expected_count = ""
        if self.count is not None:
            expected_count = " (expected %d time%s)" % (self.count, plural)
        template = "<%s %s>%s%s</%s>%s" % (
            self.tag_type, attrs, text, children, self.tag_type,
            expected_count)
        return 'Tag("%s", %s)' % (self.identifier, template)

    def _get_extra_info(self, html, results_cb):
        full_matches = [str(a) for a in self._get_matches(
                html, self.attrs, self.text)]
        if ((self.count is None and len(full_matches) < 1)
                or (self.count is not None
                    and self.count > len(full_matches))):
            def add_candidates(candidates, explanation):
                filtered_candidates = [c for c in candidates
                        if c not in full_matches]
                if len(filtered_candidates) < 1:
                    return
                results_cb(filtered_candidates, explanation)
            def get_matches(attrs, text, explanation):
                candidates = [str(a) for a in self._get_matches(
                    html, attrs, text)]
                add_candidates(candidates, explanation)
            for attr in sorted(self.attrs.keys()):
                attrs = self.attrs.copy()
                del attrs[attr]
                get_matches(attrs, self.text,
                        'attribute %s="%s"' % (attr, self.attrs[attr]))
            if self.text is not _not_passed:
                get_matches(self.attrs, _not_passed, 'text="%s"' % self.text)

    def get_extra_info(self, html_list, identifier_suffix):
        part_counts = {}
        def add_part_count(matches, explanation):
            part_counts.setdefault(explanation, 0)
            part_counts[explanation] += len(matches)
        for html in html_list:
            self._get_extra_info(html, add_part_count)
        extra_info = []
        for explanation in sorted(part_counts.keys()):
            extra_info.append('%d matches for "%s"%s when %s is not '
                    'a requirement.' % (part_counts[explanation],
                        self.identifier, identifier_suffix, explanation))
        return extra_info


def TagWithId(id, identifier=None, tag_type=True, attrs=None,
              text=_not_passed, count=1):
    """Match a tag with a particular id.

    :param id: the id that the tag must have to match.
    :type id: str
    :param identifier: override the human readable identifier of this tag.
    :type identifier: str or None
    :param tag_type: the tag type that the tag with the given id must
        have. The default is to allow any type of tag.
    :type tag_type: str, regex, dict, callable, or True (default True)
    :param attrs: a dict containing the attributes that the tag must
        have to match.
    :type attrs: dict
    :param text: the text that should be in the matched tag.
    :param count: the number of times that the tag should be found
        in the document to match. Default is 1 as ids shouldn't be shared.
    :type count: int or None to match at least once.
    """
    if identifier is None:
        identifier = '#'+id
    if attrs is None:
        attrs = {}
    attrs['id'] = id
    return Tag(identifier, tag_type, attrs=attrs, text=text, count=count)


class Within(DocumentPart):
    """Matches if one part of the document exists within another."""

    def __init__(self, outer, inner):
        super(Within, self).__init__()
        self.outer = outer
        self.inner = inner

    def __str__(self):
        return "%s within %s" % (self.inner, self.outer)

    @property
    def identifier(self):
        return '"%s" within "%s"' % (self.inner.identifier,
                self.outer.identifier)

    def get_matches(self, matchee):
        candidates = self.outer.get_matches(matchee)
        if len(candidates) < 1:
            return candidates
        for candidate in candidates:
            matches = self.inner.get_matches(candidate)
            if len(matches) > 0:
                return [candidate]
        return []

    def get_extra_info(self, html_list, identifier_suffix):
        extra_info = []
        outer_matches = []
        inner_matches = []
        for html in html_list:
            outer_matches += self.outer.get_matches(html)
            inner_matches += self.inner.get_matches(html)
        extra_info.append('%d matches for "%s"%s in the document.'
            % (len(inner_matches), self.inner.identifier, identifier_suffix))
        extra_info.append('%d matches for "%s"%s in the document.'
            % (len(outer_matches), self.outer.identifier, identifier_suffix))
        extra_info += self.inner.get_extra_info(outer_matches,
                ' within "%s"' % self.outer.identifier + identifier_suffix)
        extra_info += self.outer.get_extra_info(outer_matches,
                ' within "%s"' % self.outer.identifier + identifier_suffix)
        return extra_info


class Form(Tag):
    """Matches a form tag with a set of elements."""

    def __init__(self, identifier, action=None, method=None, attrs=None):
        if method is not None:
            if attrs is None:
                attrs = {}
            attrs['method'] = method
        if action is not None:
            if attrs is None:
                attrs = {}
            attrs['action'] = action
        super(Form, self).__init__(identifier, "form", attrs=attrs)


class Select(DocumentPart):
    """Matches a select tag with a set of choices."""

    def __init__(self, identifier, name=None, attrs=None, choices=None,
            selected=None):
        super(Select, self).__init__()
        if name is not None:
            if attrs is None:
                attrs = {}
            attrs["name"] = name
        self.select_matcher = Tag(identifier + " <select> tag", "select",
                attrs=attrs)
        self._identifier = identifier
        choice_matchers = []
        if choices is not None:
            for choice_value, choice_text in choices.items():
                attrs = {"value": choice_value}
                if selected is not None and choice_value == selected:
                    attrs["selected"] = "selected"
                choice_matchers.append(
                        Tag(identifier + " %s option" % choice_value,
                            "option", attrs=attrs, text=choice_text))
        self.choice_matchers = matchers.MatchesAll(*choice_matchers)

    def __str__(self):
        return ("[%s] within %s"
            % (", ".join([str(a) for a in self.choice_matchers.matchers]),
                self.select_matcher))

    @property
    def identifier(self):
        return self._identifier

    def get_matches(self, html):
        candidates = self.select_matcher.get_matches(html)
        if len(candidates) < 1:
            return candidates
        matches = []
        for candidate in candidates:
            found_all = True
            for matcher in self.choice_matchers.matchers:
                submatches = matcher.get_matches(candidate)
                if len(submatches) < 1:
                    found_all = False
                    break
            if found_all:
                matches.append(candidate)
        return matches

    def get_extra_info(self, html_list, identifier_suffix):
        extra_info = []
        select_matches = []
        for html in html_list:
            select_matches += self.select_matcher.get_matches(html)
        extra_info.append('%d matches for "%s"%s in the document.'
                % (len(select_matches), self.select_matcher.identifier,
                    identifier_suffix))
        extra_info += self.select_matcher.get_extra_info(
                html_list, identifier_suffix)
        within_clause = (' within "%s"' % self.select_matcher.identifier
                + identifier_suffix)
        if len(select_matches) > 0:
            for matcher in self.choice_matchers.matchers:
                total_matches = []
                for select_match in select_matches:
                    total_matches += matcher.get_matches(select_match)
                extra_info.append('%d matches for "%s"%s in the document.'
                        % (len(total_matches), matcher.identifier,
                            identifier_suffix))
                extra_info += matcher.get_extra_info(select_matches,
                        within_clause)
        return extra_info
