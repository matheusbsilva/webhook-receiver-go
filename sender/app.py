import os
import uuid
import logging
import time
import random
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
        time.sleep(random.choice(range(1, 5)))
        current_dt = datetime.now().isoformat()
        data = generate_data()
        response = requests.post(
            webhook_url,
            json=data,
        )

        logging.info(f'{current_dt} - {response.status_code} - {data["id"]}')
