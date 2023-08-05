import decimal
import json
from math import sin, cos, sqrt, atan2, radians
from time import sleep

from botocore.exceptions import ClientError
from dateutil.parser import parse

from settings import dynamodb_resource as dynamodb

COLUMN_NAME_MAPPING = {
    "Bedrms": "bedrooms",
    "Bedrms Above Grade": "bedroomsAG",
    "Beds Total": "bedroomsTotal",
    "Baths Full": "bathroomsFull",
    "Baths Half": "bathroomsHalf",
    "Tot Flr Area A.G. (SF)": "livingSpaceSquareFootageAG",
    "Tot Flr Area AG Metres": "livingSpaceSquareMetersAG",
    "Listing Firm 1 Name": "listingFirmName",
    "MLS\u00ae Number": "mlsNumber",
}


def isfloat(x):
    try:
        a = float(x)
    except ValueError:
        return False
    else:
        return True


def isint(x):
    try:
        a = float(x)
        b = int(a)
    except ValueError:
        return False
    else:
        return a == b


def filter_dict_keys(keys, raw_dict: dict):
    new_dict = {}

    for key in keys:
        if key in raw_dict.keys():
            new_dict[key] = raw_dict[key]

    return new_dict


def date_string_to_iso_string(date_string: str):
    iso_string = date_string
    try:
        iso_string = parse(date_string).isoformat()
    except Exception as e:
        pass
    return iso_string


def string_to_float(raw_str: str):
    if raw_str[0] == '$':
        raw_str = raw_str.replace('$', '')

    if isfloat(raw_str.replace(',', '')):
        raw_str = raw_str.replace(',', '')
    if isint(raw_str):
        raw_str = float(raw_str)
        raw_str = int(raw_str)
    elif isfloat(raw_str):
        raw_str = float(raw_str)

    return raw_str


def dynamo_to_dict(dynamo_data):
    if not dynamo_data:
        return dynamo_data
    return json.loads(DecimalEncoder().encode(dynamo_data))


def dict_to_dynamo(raw_dict):
    dynamo_dict = {}
    for key, val in raw_dict.items():
        if isinstance(val, float):
            dynamo_dict[key] = decimal.Decimal(str(val))
        elif isinstance(val, dict):
            dynamo_dict[key] = dict_to_dynamo(raw_dict[key])
        else:
            dynamo_dict[key] = val
    return dynamo_dict


class DecimalEncoder(json.JSONEncoder):
    def default(self, o):
        if isinstance(o, decimal.Decimal):
            if abs(o % 1) > 0:  # without abs() negative float values were being converted to int
                return float(o)
            else:
                return int(o)
        elif isinstance(o, set):
            return list(o)
        return super(DecimalEncoder, self).default(o)


def get_dynamo_items(table_name, table_filter=None, *args, **kwargs):
    RETRY_EXCEPTIONS = ('ProvisionedThroughputExceededException',
                        'ThrottlingException')
    table = dynamodb.Table(table_name)

    response = table.scan(**table_filter)
    items = response['Items']

    # todo watch for memory overflow in batching all the items as once
    # https://gist.github.com/shentonfreude/8d26ca1fc93fdb801b2c
    # https://github.com/boto/boto3/issues/597#issuecomment-323982159
    retries = 0
    max_retry = 3
    while 'LastEvaluatedKey' in response and retries < max_retry:
        try:

            response = table.scan(ExclusiveStartKey=response['LastEvaluatedKey'], **table_filter)
            items.extend(response['Items'])
            retries = 0
        except ClientError as err:
            if err.response['Error']['Code'] not in RETRY_EXCEPTIONS:
                raise
            print('WHOA, too fast, slow it down retries={}, items={}'.format(retries, len(items)), )
            sleep(2 ** retries)
            retries += 1  # TODO max limit

    return dynamo_to_dict(items)


def lat_long_distance(pointa, pointb):
    # approximate radius of earth in km
    R = 6373.0000000000

    lat1 = radians(pointa['latitude'])
    lon1 = radians(pointa['longitude'])
    lat2 = radians(pointb['latitude'])
    lon2 = radians(pointb['longitude'])

    dlon = lon2 - lon1
    dlat = lat2 - lat1

    float_2 = 2.000000000  # to prevent loss of precision
    a = sin(dlat / float_2) ** float_2 + cos(lat1) * cos(lat2) * sin(dlon / float_2) ** float_2
    c = float_2 * atan2(sqrt(a), sqrt(1.000000000 - a))

    distance = R * c

    return distance
