#!/usr/bin/env python
# -*- coding: utf-8 -*-

from QcloudApi.qcloudapi import QcloudApi
import json
import requests

def my_wan_ip():
   url = 'https://api.ipify.org?format=json'
   r = requests.get(url, timeout=10) 
   ip = r.json()['ip']
   return ip

class TX_CNS:

    def __init__(self, secretId, secretKey, domain):
        module = 'cns'
        config = {
            'secretId': secretId,
            'secretKey': secretKey
        }
        self.service = QcloudApi(module, config)

        self.domain = domain

    def get_record_id(self, subDomain):
        action = 'RecordList'
        action_params = {
            'domain': self.domain,
            'subDomain': subDomain,
        }

        try:
            self.service.generateUrl(action, action_params)
            res = json.loads(self.service.call(action, action_params))
        except Exception as e:
            import traceback
            print('traceback.format_exc():\n%s' % traceback.format_exc())

        return res['data']['records'][0]['id']

    def modify_record(self, subDomain, recordId, value):
        action = 'RecordModify'
        action_params = {
            'domain': self.domain,
            'subDomain': subDomain,
            'recordId': recordId,
            'recordType': 'A',
            'recordLine': '默认',
            'value': str(value)
        }
        
        try:
            self.service.generateUrl(action, action_params)
            res = json.loads(self.service.call(action, action_params))
        except Exception as e:
            import traceback
            print('traceback.format_exc():\n%s' % traceback.format_exc())

        return res['code'], res['codeDesc']

def main():
    import sys
    if len(sys.argv) != 4:
        print 'Vars inputed err!'
        sys.exit(1)

    import os
    secretId = os.environ.get('TX_Secret_Id')
    secretKey = os.environ.get('TX_Secret_Key')

    domain = sys.argv[1]
    subDomain = sys.argv[2]
    value = sys.argv[3]

    if value == "auto":
        value = my_wan_ip()

    cns = TX_CNS(secretId, secretKey, domain)
    record_id = cns.get_record_id(subDomain)
    code, ret = cns.modify_record(subDomain, record_id, value)
    print "Update record: %s.%s A %s, %s!" % (subDomain, domain, value, ret)

if __name__ == '__main__':
    #main()
    print(my_wan_ip())
