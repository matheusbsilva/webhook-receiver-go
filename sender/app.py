import os
import uuid
import logging
from datetime import datetime

import requests

logging.basicConfig(format='%(levelname)s: %(message)s', level=logging.INFO)


def generate_data():
    return {
        'id': str(uuid.uuid4()),
        'data': {
            'name': f"e-{uuid.uuid4()}",
            'created_at': datetime.now().isoformat()
        }

    }


if __name__ == "__main__":
    webhook_url = os.environ.get('WEBHOOK_URL')

    logging.info('Starting webhook sender job...')
    logging.info(f'Sending data to {webhook_url}')

    while True:
        current_dt = datetime.now().isoformat()
        response = requests.post(
            webhook_url,
            json=generate_data(),
        )

        logging.info(f'{current_dt} - {response.status_code} - {response.json()}')
