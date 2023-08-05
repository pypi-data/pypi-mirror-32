import os
import pprint

from dotenv import find_dotenv, load_dotenv

import s3qlite
import error_handlers

load_dotenv(find_dotenv())

client = s3qlite.client(
        region_name='nyc3',
        endpoint_url='https://nyc3.digitaloceanspaces.com',
        aws_access_key_id=os.environ['DIGITALOCEAN_ACCESS_KEY_ID'],
        aws_secret_access_key=os.environ['DIGITALOCEAN_SECRET_ACCESS_KEY'])
client.connect('mittab-backups')

if __name__ == '__main__':
    pprint.pprint(client.execute('select ballot_code from tab_judge',
                                 on_error=error_handlers.raise_error))
