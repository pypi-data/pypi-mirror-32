#!/usr/bin/env python
# coding: utf-8


from __future__ import print_function
from __future__ import unicode_literals
import argparse
import logging
from pykeepass import PyKeePass
from easypysmb import EasyPySMB


logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger(__name__)


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument(
        '-p', '--password',
        help='Password of the KDBX database',
        required=True
    )
    parser.add_argument(
        '-k', '--keyfile',
        help='Keyfile to unlock the KDBX database',
        required=False
    )
    parser.add_argument(
        '-d', '--database',
        # type=argparse.FileType('r'),
        help='Database (KDBX file)',
        required=True
    )
    parser.add_argument(
        '-o', '--outfile',
        # type=argparse.FileType('w'),
        help='File to write the updated database to',
        required=False
    )
    parser.add_argument(
        '-D', '--destination',
        help='Group where to write the new entry to (path)',
        default='',
        required=False
    )
    parser.add_argument(
        '-f', '--force',
        help='Force the creation of a new entry, even if there already is '\
            'one with the same title',
        action='store_true',
        default=False,
        required=False
    )
    parser.add_argument(
        '-e', '--entry',
        help='Name (title) of the new entry',
        required=True
    )
    parser.add_argument(
        '-U', '--entry-username',
        help='Username to put in the new entry',
        required=True
    )
    parser.add_argument(
        '-P', '--entry-password',
        help='Password to put in the new entry',
        required=True
    )
    parser.add_argument(
        '--entry-url',
        help='URL of the new entry',
        required=False
    )
    parser.add_argument(
        '-N', '--entry-notes',
        help='Notes for the new entry',
        required=False
    )
    parser.add_argument(
        '-T', '--entry-tags',
        help='Tags for the entry',
        action='append',
        required=False
    )
    return parser.parse_args()


def smb_retrieve(samba_path):
    # smb://DOMAIN;USER:PASSWORD@SERVER:SHARE/PATH
    logger.info('Retrieve database from Samba share')
    e = EasyPySMB(samba_path)
    f = e.retrieve_file()
    e.close()
    return f.name


def smb_send(db_file, samba_path):
    logger.info('Upload {} to {}'.format(db_file, samba_path))
    e = EasyPySMB(samba_path)
    res = e.store_file(db_file)
    e.close()
    return res


def write_entry(kdbx_file, kdbx_password, group_path,
                entry_title, entry_username, entry_password, entry_url,
                entry_notes, entry_tags, kdbx_keyfile=None,
                force_creation=False, outfile=None):
    logging.info(
        'Attempt to write entry "{}: {}:{}" to {}'.format(
            entry_title, entry_username, entry_password, group_path
        )
    )
    samba_db = False
    if kdbx_file.startswith('smb://'):
        samba_db = True
        smb_kdbx_file = smb_retrieve(kdbx_file)
    kp = PyKeePass(
        smb_kdbx_file if samba_db else kdbx_file,
        password=kdbx_password,
        keyfile=kdbx_keyfile
    )
    dest_group = kp.find_groups_by_path(group_path, first=True)
    kp.add_entry(
        destination_group=dest_group,
        title=entry_title,
        username=entry_username,
        password=entry_password,
        url=entry_url,
        notes=entry_notes,
        tags=entry_tags,
        force_creation=force_creation
    )

    if outfile:
        if outfile.startswith('smb://'):
            file_written = kp.save()
            smb_send(file_written.name, outfile)
            logging.info('Sent database file to {}'.format(outfile))
        else:
            file_written = kp.save(kdbx_file)
            logging.info('KeePass DB written to {}'.format(file_written.name))
    else:
        if samba_db:
            file_written = kp.save()
            smb_send(file_written.name, kdbx_file)
            logging.info('Sent database file to {}'.format(kdbx_file))
        else:
            file_written = kp.save(kdbx_file)
            logging.info('KeePass DB written to {}'.format(file_written.name))


def main():
    args = parse_args()
    write_entry(
        kdbx_file=args.database,
        kdbx_password=args.password,
        kdbx_keyfile=args.keyfile,
        outfile=args.outfile,
        group_path=args.destination,
        force_creation=args.force,
        entry_title=args.entry,
        entry_username=args.entry_username,
        entry_password=args.entry_password,
        entry_url=args.entry_url,
        entry_notes=args.entry_notes,
        entry_tags=args.entry_tags,
    )


if __name__ == '__main__':
    main()
