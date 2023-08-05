import types

from defusedxml.sax import parse
from rdflib import Graph, URIRef
from rdflib.exceptions import ParserError
from StringIO import StringIO
from xml.sax import ContentHandler, SAXParseException


from validator.constants import APPLICATIONS, APPROVED_APPLICATIONS


class RDFException(Exception):
    """Exception thrown when the RDF parser encounters a problem."""

    def __init__(self, message=None, orig_exception=None):
        if message is None and orig_exception is not None:
            message = orig_exception.getMessage()

        super(RDFException, self).__init__(message)
        self.orig_exception = orig_exception

    def line(self):
        return (self.orig_exception.getLineNumber() if self.orig_exception else
                None)


class AddonRDFEntity(object):
    """
    A "resolved" entity within an RDF file in an add-on. For use by SAX during
    the entity resolution process.
    """

    def getByteStream(self):
        yield None

    def getSystemId(self):
        return ''


class AddonRDFEntityResolver(object):
    """
    An entity resolver to be used by SAX for resolving internal entity
    references.
    """

    def __init__(self, err):
        self.err = err

    def resolveEntity(self, public, system):
        if system.startswith('data:'):
            self.err.warning(
                    err_id=('rdf', 'entity_resolver', 'data_uri'),
                    warning='`data:` URIs are not permitted in `install.rdf`.',
                    filename='install.rdf')
        elif system.startswith('chrome://'):
            self.err.warning(
                    err_id=('rdf', 'entity_resolver', 'chrome_uri'),
                    warning='`chrome://` URI referenced before initialization.',
                    description='A chrome URI was referenced before the '
                                'browser chrome was initialized.',
                    filename='install.rdf')
        else:
            self.err.warning(
                    err_id=('rdf', 'entity_resolver', 'remote_uri'),
                    warning='Remote URI referenced from `install.rdf`.',
                    description='Remote URIs should not be used within '
                                '`install.rdf` files.',
                    filename='install.rdf')

        return AddonRDFEntity()


class RDFParser(object):
    """Parser wrapper for RDF files."""

    def __init__(self, err, data, namespace=None):
        self.err = err
        self.manifest = u'urn:mozilla:install-manifest'
        self.namespace = namespace or 'http://www.mozilla.org/2004/em-rdf'

        if (hasattr(data, 'read') and hasattr(data, 'readline') or
            isinstance(data, StringIO)
        ):
            # It could be a file-like object, let's read it so that we can
            # wrap it in StringIO so that we can re-open at any time
            data.seek(0)
            data = data.read()

        try:
            # Use an empty ContentHandler, we just want to make sure it parses.
            parse(StringIO(data), ContentHandler())
        except SAXParseException as ex:
            # Raise the SAX parse exceptions so we get some line info.
            raise RDFException(orig_exception=ex)

        from rdflib.plugins.parsers import rdfxml
        orig_create_parser = rdfxml.create_parser

        try:
            # Patch rdflib to not resolve URL entities.
            def create_parser(*args, **kwargs):
                parser = orig_create_parser(*args, **kwargs)
                parser.setEntityResolver(AddonRDFEntityResolver(err))
                return parser
            rdfxml.create_parser = create_parser

            # Load up and parse the file in XML format.
            graph = Graph()
            graph.parse(StringIO(data), format='xml')
            self.rdf = graph

        except ParserError as ex:
            # Re-raise the exception in a local exception type.
            raise RDFException(message=ex.message)
        except SAXParseException as ex:
            # Raise the SAX parse exceptions so we get some line info.
            raise RDFException(orig_exception=ex)
        finally:
            # If we fail, we don't want to sully up the creation function.
            rdfxml.create_parser = orig_create_parser

    def uri(self, element, namespace=None):
        'Returns a URIRef object for use with the RDF document.'

        if namespace is None:
            namespace = self.namespace

        return URIRef('%s#%s' % (namespace, element))

    def get_root_subject(self):
        'Returns the BNode which describes the topmost subject of the graph.'

        manifest = URIRef(self.manifest)

        if list(self.rdf.triples((manifest, None, None))):
            return manifest
        else:
            return self.rdf.subjects(None, self.manifest).next()

    def get_object(self, subject=None, predicate=None):
        """Eliminates some of the glue code for searching RDF. Pass
        in a URIRef object (generated by the `uri` function above or
        a BNode object (returned by this function) for either of the
        parameters."""

        # Get the result of the search
        results = self.rdf.objects(subject, predicate)
        as_list = list(results)

        # Don't raise exceptions, value test!
        if not as_list:
            return None

        return as_list[0]

    def get_objects(self, subject=None, predicate=None):
        """Same as get_object, except returns a list of objects which
        satisfy the query rather than a single result."""

        # Get the result of the search
        results = self.rdf.objects(subject, predicate)
        return list(results)

    def get_applications(self):
        """Return the list of supported applications."""
        applications = []

        # Isolate all of the bnodes referring to target applications
        for target_app in self.get_objects(None,
                                           self.uri('targetApplication')):
            applications.append({
                'guid': self.get_object(target_app, self.uri('id')),
                'min_version': self.get_object(target_app,
                                               self.uri('minVersion')),
                'max_version': self.get_object(target_app,
                                               self.uri('maxVersion'))})
        return applications
