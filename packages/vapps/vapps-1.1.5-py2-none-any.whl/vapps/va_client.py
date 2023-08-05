import os.path

from pprint import PrettyPrinter
import argparse, json
from va_api import APIManager

from va_client_utils import module_mappings, module_args, get_mapping_arguments, extract_data_with_keys

base_path = os.path.dirname(os.path.abspath(__file__))

def call_function(user, password, endpoint, va_master_url = "https://127.0.0.1:443", data = '{}', method = 'get'):
    if type(data) == str: 
        data = json.loads(data)
#    print 'In call function with : ', user, password, endpoint, data, method
    api = APIManager(va_url=va_master_url, va_user=user, va_pass=password, verify=False)
    return api.api_call(endpoint, data = data, method = method)

def call_function_with_mapping(user, password, mapping, command, va_master_url = 'https://127.0.0.1:443', **data):
    cmd_map = get_mapping_arguments(mapping, command)
    kwargs = cmd_map['kwargs']
    kwargs['user'] = user
    kwargs['password'] = password
    kwargs['data'] = data
    kwargs['va_master_url'] = va_master_url

    result = call_function(**kwargs)
    if type(result) == dict: 
        if not result['success'] : 
            raise Exception('API call failed : ', result['message'])

        result = result['data']
        if type(result) == dict: 
            result = result[result.keys()[0]]
            result = extract_data_with_keys(result, cmd_map.get('keys'))

    return result

def prepare_request_args(sub_parsers):
    api_request = sub_parsers.add_parser('request', description = 'Performs an api call on the va_master. ')
    api_request.set_defaults(sub = 'api_request')

    api_request.add_argument('--endpoint', help = 'The endpoint for the request. ')
    api_request.add_argument('--method', help = 'The method to use. Default is GET. ', default = 'get')
    api_request.add_argument('--data', help= 'The data to send to the request. Must be properly formed JSON. Default: \'{}\'', default = '{}')

def prepare_conf_args(sub_parsers):
    confs = sub_parsers.add_parser('conf', description = 'Creates a VapourApps configuration file which you can use instead of using the username and password arguments all the time. ')
    confs.set_defaults(sub = 'confs')

    confs.add_argument('--user', description = 'The username which will be saved in the configuration')
    confs.add_argument('--password', description = 'The password for the configuration. ')
    confs.add_argument('--va-master-url', description = 'The URL of the VA master')

def create_va_configuration(path, username, password, va_url):
    conf_data = {"user" : username, "password" : password, "va_master_url" : va_url}
    with open(path, 'w') as f: 
        f.write(json.dumps(conf_data))
    return conf_data

def load_va_configuration(path):
    if not os.path.exists(path):
        return None

    with open(path) as f: 
        conf = f.read()

    conf = json.loads(conf)
    return conf

def get_credentials(args):
    conf = load_va_configuration(args.conf_path)
    if conf: 
        return conf
    else: 
        print 'Could not find a va_master configuration - either the --load-conf path is incorrect, or a configuration has not been generated yet. Please answer the following prompts : '
        user = raw_input("Enter the va_master username : ")
        password = raw_input("Enter the va_master password : ")
        print 'Enter the VA master url. Include the http:// and the port, example : https://127.0.0.1:443'
        va_url = raw_input("VA master url : ")

        new_conf = create_va_configuration(args.conf_path, user, password, va_url)
        return new_conf
        

def main():
    parser = argparse.ArgumentParser(description = 'VA client argument parse. ')
#    parser.add_argument('--user', help = 'The username with which to make the request. ', default = '')
#    parser.add_argument('--password', help = 'The password with which to authenticate the request. ', default = '')
    parser.add_argument('--conf-path', help = 'If you have a configuration file in a custom location, you can specify said location with this argument. You may also have multiple configurations and choose one with this. ', default = base_path + '/vapp_config.json')

    sub_parsers = parser.add_subparsers()
    prepare_request_args(sub_parsers)

    for module in module_args: 
        m = sub_parsers.add_parser(module, description = module_args[module].get('description', ''))
        m.set_defaults(sub = module)

        m.add_argument('command', nargs = '*', choices = module_args[module]['choices'], help = module_args[module].get('help', ''))

        for a in module_args[module].get('extra_args', []):
            m.add_argument(a)


    args, unknown = parser.parse_known_args()
    creds = get_credentials(args)

    kwargs = vars(args)
    kwargs = {x : kwargs[x] for x in kwargs if x[0] != '_' and x not in ['sub']}
    kwargs.update(creds)

    if args.sub == 'api_request' : 
        result = call_function(**kwargs)
        print result

    else: 
        mapping = module_mappings[args.sub]
        kwargs['mapping'] = mapping
        result = call_function_with_mapping(**kwargs)

    if type(result) in [list, dict]: 
        pprinter = PrettyPrinter(indent = 4)
        pprinter.pprint(result)
    else: 
        print result

#main()
