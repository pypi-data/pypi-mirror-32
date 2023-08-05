"""
These are the converter functions for the NTCIR-10 Math converter package.
"""

from logging import getLogger
from multiprocessing import Pool
from pathlib import Path
from zipfile import ZipFile

from lxml import etree
from lxml.etree import Element, SubElement, QName
from tqdm import tqdm


LOGGER = getLogger(__name__)
PARAGRAPH_XPATH = "|".join([
    ".//xhtml:div[contains(concat(' ', normalize-space(@class), ' '), ' para ')]",
    ".//xhtml:div[contains(concat(' ', normalize-space(@class), ' '), ' caption ') and .//m:*]",
    ".//xhtml:div[contains(concat(' ', normalize-space(@class), ' '), ' bibblock ') and .//m:*]",
    ".//xhtml:table[.//m:*]",
    ".//xhtml:p[contains(concat(' ', normalize-space(@class), ' '), ' p ') and .//m:*]",
    ".//*[contains(concat(' ', normalize-space(@class), ' '), ' title ') and .//m:*]",
])
NAMESPACES = {
    "xhtml": "http://www.w3.org/1999/xhtml",
    "m": "http://www.w3.org/1998/Math/MathML",
}


def get_judged_identifiers(input_file):
    """
    Extracts the names of judged documents and the identifiers of their judged elements from
    relevance in the NTCIR-10 Math format.

    Parameters
    ----------
    input_file : file
        The input file containing relevance judgements in the NTCIR-10 Math format.

    Yields
    ------
    (str, str)
        Judged document names and the judged element identifiers
    """
    for line in tqdm(list(input_file)):
        identifier = line.split(' ')[2]
        document_basename, element_id = identifier.split('#')
        document = Path(document_basename).with_suffix(".xhtml")
        yield document.name, element_id


def convert_judgements(input_file, output_file, identifier_map):
    """
    Converts relevance judgements from the NTCIR-10 Math format to the NTCIR-11 Math-2, and the
    NTCIR-12 MathIR format.

    Parameters
    ----------
    input_file : file
        The input file containing relevance judgements in the NTCIR-10 Math format.
    output_file : file
        The output file that will contain relevance judgements in the NTCIR-11 Math-2, and the
        NTCIR-12 MathIR format.
    identifier_map : dict of (str, str)
        A mapping between element identifiers, and paragraph identifiers.
    """
    for line in tqdm(list(input_file)):
        topic, unused, identifier, score = line.split(' ')
        output_file.write(
            "%s %s %s %d\n" % (
                topic, unused, identifier_map[identifier], int(score)))


def process_dataset(input_root_dir, output_root_dir=None, judged_identifiers=None, num_workers=1):
    """
    Processes the NTCIR-10 Math dataset, building a mapping between element identifiers, and
    paragraph identifiers, and optionally also building an equivalent dataset in the NTCIR-11
    Math-2, and NTCIR-12 MathIR format.

    Parameters
    ----------
    input_root_dir : pathlib.Path
        The input directory containing the NTCIR-10 Math dataset.
    output_root_dir : pathlib.Path or None, optional
        The output directory that will contain the dataset from the input directory converted to the
        NTCIR-11 Math-2, and the NTCIR-12 MathIR format. If None, no conversion will be performed.
    judged_identifiers : dict of (str, set of str) or None, optional
        The names of judged documents and the identifiers of their judged elements. This constrains
        the documents that are actually processed when we are not building a dataset, and the
        elements whose identifiers are recorded in the mapping between element identifiers, and
        paragraph identifiers.
    num_workers : int, optional
        The number of processes that will process the documents in the NTCIR-10 Math dataset.

    Returns
    -------
    dict of (str, str)
        A mapping between element identifiers, and paragraph identifiers.
    """
    dataset_identifier_map = {}
    if output_root_dir:
        LOGGER.info("Converting dataset %s -> %s", input_root_dir, output_root_dir)
    if judged_identifiers is not None:
        LOGGER.info("Building a mapping between element identifiers, and paragraph identifiers")
    with Pool(num_workers) as pool:
        input_files = [
            input_file for input_file in input_root_dir.glob("**/*.xhtml")
            if output_root_dir or judged_identifiers is None
            or input_file.name in judged_identifiers]
        for document_identifier_map in tqdm(pool.imap_unordered(
                _process_document_worker,
                (
                    (
                        input_file, output_root_dir,
                        input_file.relative_to(input_root_dir).with_suffix(""),
                        judged_identifiers[input_file.name]
                        if judged_identifiers is not None
                        and input_file.name in judged_identifiers else None
                    )
                    for input_file in input_files
                )), total=len(input_files)):
            dataset_identifier_map.update(document_identifier_map)
    return dataset_identifier_map


def _process_document_worker(args):
    input_file, output_root_dir, output_dir, judged_element_identifiers = args
    document_identifier_map = {}
    LOGGER.debug("Processing document %s", input_file)
    if output_root_dir:
        LOGGER.debug("Creating directory %s", output_root_dir / output_dir)
        (output_root_dir / output_dir).mkdir(parents=True)
    with input_file.open("rt") as f:
        input_tree = etree.parse(f)
    input_all_paragraphs = input_tree.xpath(PARAGRAPH_XPATH, namespaces=NAMESPACES)
    input_paragraphs = set(input_all_paragraphs)
    for input_paragraph in input_all_paragraphs:
        for input_paragraph_descendant in input_paragraph.xpath(
                PARAGRAPH_XPATH, namespaces=NAMESPACES):
            if input_paragraph_descendant in input_paragraphs:
                LOGGER.debug(
                    "Skipping paragraph %s in document %s, because there exists an ancestor",
                    input_paragraph_descendant.attrib["id"]
                    if "id" in input_paragraph_descendant.attrib else "(unknown id)",
                    input_file)
                input_paragraphs.remove(input_paragraph_descendant)
    for input_paragraph_num, input_paragraph in enumerate(input_paragraphs):
        if "id" not in input_paragraph.attrib:
            LOGGER.warning(
                "Skipping a paragraph in document %s, because it lacks an identifier", input_file)
            continue
        LOGGER.debug(
            "Processing paragraph %s in document %s", input_paragraph.attrib["id"], input_file)
        output_file = Path(
            "%s_1_%d" % (
                str(input_file.with_suffix("").name), input_paragraph_num
            )).with_suffix(".xhtml.zip")
        for input_element in input_paragraph.findall(".//*[@id]"):
            if judged_element_identifiers is not None and \
                    input_element.attrib["id"] in judged_element_identifiers:
                ntcir10_identifier = "%s#%s" % (
                    str(input_file.with_suffix("").name), input_element.attrib["id"])
                ntcir11_12_identifier = str(output_file.with_suffix("").with_suffix("").name)
                LOGGER.debug("Mapping %s -> %s", ntcir10_identifier, ntcir11_12_identifier)
                if ntcir10_identifier in document_identifier_map \
                        and document_identifier_map[ntcir10_identifier] != ntcir11_12_identifier:
                    LOGGER.warning(
                        "Duplicate element with element id %s occurs in paragraphs %s, and %s",
                        ntcir10_identifier, document_identifier_map[ntcir10_identifier],
                        ntcir11_12_identifier)
                else:
                    document_identifier_map[ntcir10_identifier] = ntcir11_12_identifier
        if output_root_dir:
            LOGGER.debug("Creating ZIP archive %s", output_root_dir / output_dir / output_file)
            with ZipFile((output_root_dir / output_dir / output_file).open("wb"), "w") as zip_file:
                LOGGER.debug(
                    "Creating archived file %s/%s",
                    (output_root_dir / output_dir / output_file),
                    output_file.with_suffix("").name)
                html = Element(QName(NAMESPACES["xhtml"], "html"), {}, {None: NAMESPACES["xhtml"]})
                head = SubElement(html, QName(NAMESPACES["xhtml"], "head"))
                SubElement(head, QName(NAMESPACES["xhtml"], "meta"), {
                    "http-equiv": "Content-Type",
                    "content": "application/xhtml+xml; charset=UTF-8"
                })
                body = SubElement(html, QName(NAMESPACES["xhtml"], "body"))
                body.append(input_paragraph)
                zip_file.writestr(output_file.with_suffix("").name, etree.tostring(html))
    return document_identifier_map
